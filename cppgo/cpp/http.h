#pragma once

#include "buffer.h"// 使用之前实现的Buffer
#include <string>
#include <functional>
#include <map>
#include <vector>
#include <sstream>
#include <stdexcept>
#include <cctype>
#include <cstdlib>
#include <ctime>
#include <fcntl.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

struct HTTPError : std::runtime_error {
    explicit HTTPError(const std::string &msg) : std::runtime_error(msg) {}
};

enum HTTPMethod { GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH };

struct Request {
    HTTPMethod                         method;
    std::string                        path;
    std::string                        protocol;
    std::map<std::string, std::string> headers;
    Buffer                             body;
    int                                clientSocket;

    std::string header(const std::string &key) const {
        auto it = headers.find(key);
        if (it != headers.end()) { return it->second; }
        return "";
    }

    std::map<std::string, std::string> query() const {
        std::map<std::string, std::string> params;
        size_t                             pos = path.find('?');
        if (pos != std::string::npos) {
            std::string queryString = path.substr(pos + 1);
            size_t      start = 0;

            while (start < queryString.size()) {
                size_t eqPos = queryString.find('=', start);
                if (eqPos == std::string::npos) break;

                size_t ampPos = queryString.find('&', eqPos);
                if (ampPos == std::string::npos) ampPos = queryString.size();

                std::string key = queryString.substr(start, eqPos - start);
                std::string value = queryString.substr(eqPos + 1, ampPos - eqPos - 1);

                params[urlDecode(key)] = urlDecode(value);
                start = ampPos + 1;
            }
        }
        return params;
    }

private:
    static std::string urlDecode(const std::string &str) {
        std::string result;
        for (size_t i = 0; i < str.size(); i++) {
            if (str[i] == '+') {
                result += ' ';
            } else if (str[i] == '%' && i + 2 < str.size()) {
                int                value;
                std::istringstream iss(str.substr(i + 1, 2));
                if (iss >> std::hex >> value) {
                    result += static_cast<char>(value);
                    i += 2;
                } else {
                    result += '%';
                }
            } else {
                result += str[i];
            }
        }
        return result;
    }
};

struct Response {
    int                                statusCode = 200;
    std::map<std::string, std::string> headers;
    Buffer                             body;

    void setContentType(const std::string &contentType) { headers["Content-Type"] = contentType; }

    void setStatus(int code) { statusCode = code; }

    void write(const Buffer &content) { body.WriteBuffer(content); }

    void writeString(const std::string &str) { body.WriteString(str); }
};

using HandlerFunc = std::function<void(Request &, Response &)>;

struct Server {
    int                                serverSocket = -1;
    std::map<std::string, HandlerFunc> handlers;

    void handle(HTTPMethod method, const std::string &path, HandlerFunc handler) {
        std::string key = methodToString(method) + ":" + path;
        handlers[key] = handler;
    }

    void listenAndServe(const std::string &addr) {
        serverSocket = socket(AF_INET, SOCK_STREAM, 0);
        if (serverSocket == -1) { throw HTTPError("Failed to create socket"); }

        size_t      pos = addr.find(':');
        std::string host = "0.0.0.0";
        int         port = 8080;
        if (pos != std::string::npos) {
            host = addr.substr(0, pos);
            port = std::stoi(addr.substr(pos + 1));
        }

        int opt = 1;
        setsockopt(serverSocket, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));

        sockaddr_in serverAddr{};
        serverAddr.sin_family = AF_INET;
        serverAddr.sin_addr.s_addr = inet_addr(host.c_str());
        serverAddr.sin_port = htons(port);

        if (bind(serverSocket, (sockaddr*)&serverAddr, sizeof(serverAddr)) {
            throw HTTPError("Failed to bind socket");
        }
        
        if (listen(serverSocket, 10) == -1) {
            throw HTTPError("Failed to listen on socket");
        }
        
        std::cout << "HTTP server running on " << addr << "\n";
        
        while (true) {
            sockaddr_in clientAddr{};
            socklen_t   clientLen = sizeof(clientAddr);
            int         clientSocket = accept(serverSocket, (sockaddr *) &clientAddr, &clientLen);
            if (clientSocket == -1) {
                std::cerr << "Failed to accept connection\n";
                continue;
            }

            try {
                handleConnection(clientSocket);
            } catch (const std::exception &e) { std::cerr << "Error handling connection: " << e.what() << "\n"; }

            close(clientSocket);
        }
    }

    void shutdown() {
        if (serverSocket != -1) {
            close(serverSocket);
            serverSocket = -1;
        }
    }

private:
    void handleConnection(int clientSocket) {
        Request  req = parseRequest(clientSocket);
        Response res;

        std::string key = methodToString(req.method) + ":" + req.path;
        auto        it = handlers.find(key);

        if (it != handlers.end()) {
            try {
                it->second(req, res);
            } catch (const std::exception &e) {
                res.setStatus(500);
                res.writeString("Internal Server Error: " + std::string(e.what()));
            }
        } else {
            res.setStatus(404);
            res.writeString("Not Found");
        }

        sendResponse(clientSocket, res);
    }

    Request parseRequest(int clientSocket) {
        Request req;
        req.clientSocket = clientSocket;

        std::string line = readLine(clientSocket);
        if (line.empty()) throw HTTPError("Empty request line");

        std::istringstream iss(line);
        std::string        methodStr;
        iss >> methodStr >> req.path >> req.protocol;

        req.method = stringToMethod(methodStr);

        while (true) {
            line = readLine(clientSocket);
            if (line.empty() || line == "\r" || line == "\r\n") break;// 头部结束

            size_t pos = line.find(':');
            if (pos != std::string::npos) {
                std::string key = line.substr(0, pos);
                std::string value = line.substr(pos + 1);

                trim(key);
                trim(value);

                req.headers[key] = value;
            }
        }

        int  contentLength = 0;
        auto it = req.headers.find("Content-Length");
        if (it != req.headers.end()) { contentLength = std::stoi(it->second); }

        if (contentLength > 0) { req.body = readBody(clientSocket, contentLength); }

        return req;
    }

    void sendResponse(int clientSocket, Response &res) {
        if (res.headers.find("Content-Type") == res.headers.end()) { res.setContentType("text/plain; charset=utf-8"); }

        res.headers["Content-Length"] = std::to_string(res.body.Len());

        std::ostringstream headers;
        headers << "HTTP/1.1 " << res.statusCode << " " << getStatusText(res.statusCode) << "\r\n";

        for (const auto &[key, value]: res.headers) { headers << key << ": " << value << "\r\n"; }

        headers << "\r\n";

        std::string headerStr = headers.str();
        send(clientSocket, headerStr.data(), headerStr.size(), 0);

        if (res.body.Len() > 0) { send(clientSocket, res.body.Data(), res.body.Len(), 0); }
    }

    static std::string readLine(int socket) {
        std::string line;
        char        c;

        while (recv(socket, &c, 1, 0) > 0) {
            line += c;
            if (c == '\n') break;
        }

        return line;
    }

    static Buffer readBody(int socket, size_t length) {
        Buffer body;
        body.Grow(length);

        ssize_t received = 0;
        while (received < static_cast<ssize_t>(length)) {
            char    buffer[1024];
            ssize_t n = recv(socket, buffer, std::min<size_t>(1024, length - received), 0);
            if (n <= 0) break;

            body.Write(buffer, n);
            received += n;
        }

        return body;
    }

    static void trim(std::string &str) {
        size_t start = 0;
        size_t end = str.size();

        while (start < end && std::isspace(str[start])) start++;
        while (end > start && std::isspace(str[end - 1])) end--;

        str = str.substr(start, end - start);
    }

    static HTTPMethod stringToMethod(const std::string &method) {
        if (method == "GET") return GET;
        if (method == "HEAD") return HEAD;
        if (method == "POST") return POST;
        if (method == "PUT") return PUT;
        if (method == "DELETE") return DELETE;
        if (method == "CONNECT") return CONNECT;
        if (method == "OPTIONS") return OPTIONS;
        if (method == "TRACE") return TRACE;
        if (method == "PATCH") return PATCH;
        throw HTTPError("Unknown HTTP method: " + method);
    }

    static std::string methodToString(HTTPMethod method) {
        switch (method) {
            case GET:
                return "GET";
            case HEAD:
                return "HEAD";
            case POST:
                return "POST";
            case PUT:
                return "PUT";
            case DELETE:
                return "DELETE";
            case CONNECT:
                return "CONNECT";
            case OPTIONS:
                return "OPTIONS";
            case TRACE:
                return "TRACE";
            case PATCH:
                return "PATCH";
            default:
                return "UNKNOWN";
        }
    }

    static std::string getStatusText(int code) {
        switch (code) {
            case 200:
                return "OK";
            case 201:
                return "Created";
            case 204:
                return "No Content";
            case 301:
                return "Moved Permanently";
            case 302:
                return "Found";
            case 400:
                return "Bad Request";
            case 401:
                return "Unauthorized";
            case 403:
                return "Forbidden";
            case 404:
                return "Not Found";
            case 405:
                return "Method Not Allowed";
            case 500:
                return "Internal Server Error";
            case 501:
                return "Not Implemented";
            case 502:
                return "Bad Gateway";
            case 503:
                return "Service Unavailable";
            default:
                return "Unknown Status";
        }
    }
};

void serveFile(Response &res, const std::string &filePath) {
    int fd = open(filePath.c_str(), O_RDONLY);
    if (fd == -1) {
        res.setStatus(404);
        res.writeString("File not found");
        return;
    }

    Buffer  content;
    char    buffer[4096];
    ssize_t n;

    while ((n = read(fd, buffer, sizeof(buffer))) > 0) { content.Write(buffer, n); }

    close(fd);

    res.setContentType("application/octet-stream");
    res.write(content);
}

void redirect(Response &res, const std::string &location, int status = 302) {
    res.setStatus(status);
    res.headers["Location"] = location;
}

HandlerFunc fileServerHandler(const std::string &rootDir) {
    return [rootDir](Request &req, Response &res) {
        std::string filePath = rootDir + req.path;
        serveFile(res, filePath);
    };
}

struct Client {
    static Response get(const std::string &url) {
        size_t pos = url.find("://");
        if (pos == std::string::npos) { throw HTTPError("Invalid URL: " + url); }

        std::string protocol = url.substr(0, pos);
        std::string hostPath = url.substr(pos + 3);

        pos = hostPath.find('/');
        std::string host = pos == std::string::npos ? hostPath : hostPath.substr(0, pos);
        std::string path = pos == std::string::npos ? "/" : hostPath.substr(pos);

        std::string hostname = host;
        int         port = protocol == "https" ? 443 : 80;

        pos = host.find(':');
        if (pos != std::string::npos) {
            hostname = host.substr(0, pos);
            port = std::stoi(host.substr(pos + 1));
        }

        int sock = socket(AF_INET, SOCK_STREAM, 0);
        if (sock == -1) { throw HTTPError("Failed to create socket"); }

        sockaddr_in serverAddr{};
        serverAddr.sin_family = AF_INET;
        serverAddr.sin_port = htons(port);

        if (inet_pton(AF_INET, hostname.c_str(), &serverAddr.sin_addr) <= 0) {
            close(sock);
            throw HTTPError("Invalid hostname: " + hostname);
        }

        if (connect(sock, (sockaddr *) &serverAddr, sizeof(serverAddr)) == -1) {
            close(sock);
            throw HTTPError("Connection failed to " + hostname + ":" + std::to_string(port));
        }

        std::string request = "GET " + path +
                              " HTTP/1.1\r\n"
                              "Host: " +
                              host +
                              "\r\n"
                              "Connection: close\r\n\r\n";

        send(sock, request.data(), request.size(), 0);

        Response res;

        std::string line = readLine(sock);
        if (line.empty()) {
            close(sock);
            throw HTTPError("Empty response");
        }

        std::istringstream iss(line);
        std::string        protocol, statusText;
        int                statusCode;
        iss >> protocol >> statusCode >> statusText;
        res.statusCode = statusCode;

        while (true) {
            line = readLine(sock);
            if (line.empty() || line == "\r" || line == "\r\n") break;

            size_t pos = line.find(':');
            if (pos != std::string::npos) {
                std::string key = line.substr(0, pos);
                std::string value = line.substr(pos + 1);
                trim(key);
                trim(value);
                res.headers[key] = value;
            }
        }

        int  contentLength = 0;
        auto it = res.headers.find("Content-Length");
        if (it != res.headers.end()) { contentLength = std::stoi(it->second); }

        if (contentLength > 0) {
            res.body = readBody(sock, contentLength);
        } else {
            char    buffer[4096];
            ssize_t n;
            while ((n = recv(sock, buffer, sizeof(buffer), 0)) > 0) { res.body.Write(buffer, n); }
        }

        close(sock);
        return res;
    }

private:
    static std::string readLine(int socket) {
        std::string line;
        char        c;

        while (recv(socket, &c, 1, 0) > 0) {
            line += c;
            if (c == '\n') break;
        }

        return line;
    }

    static Buffer readBody(int socket, size_t length) {
        Buffer body;
        body.Grow(length);

        ssize_t received = 0;
        while (received < static_cast<ssize_t>(length)) {
            char    buffer[1024];
            ssize_t n = recv(socket, buffer, std::min<size_t>(1024, length - received), 0);
            if (n <= 0) break;

            body.Write(buffer, n);
            received += n;
        }

        return body;
    }

    static void trim(std::string &str) {
        size_t start = 0;
        size_t end = str.size();

        while (start < end && std::isspace(str[start])) start++;
        while (end > start && std::isspace(str[end - 1])) end--;

        str = str.substr(start, end - start);
    }
};

#include "http.h"
#include <iostream>

void homeHandler(Request &req, Response &res) {
    res.writeString("<h1>Welcome to HTTP Server</h1>");
    res.writeString("<p>Current time: " + getCurrentTime() + "</p>");
    res.writeString("<a href='/about'>About</a>");
    res.setContentType("text/html; charset=utf-8");
}

void aboutHandler(Request &req, Response &res) {
    res.writeString("<h1>About Page</h1>");
    res.writeString("<p>This is a simple HTTP server implemented in C++</p>");
    res.setContentType("text/html; charset=utf-8");
}

std::string getCurrentTime() {
    time_t now = time(nullptr);
    return std::string(ctime(&now));
}

int main() {
    Server server;

    server.handle(GET, "/", homeHandler);
    server.handle(GET, "/about", aboutHandler);
    server.handle(GET, "/files", fileServerHandler("./public"));

    try {
        server.listenAndServe(":8080");
    } catch (const HTTPError &e) {
        std::cerr << "Server error: " << e.what() << "\n";
        return 1;
    }

    return 0;
}

void testClient() {
    try {
        Response res = Client::get("http://localhost:8080");

        std::cout << "Status: " << res.statusCode << "\n";
        std::cout << "Body: " << res.body.String() << "\n";
    } catch (const HTTPError &e) { std::cerr << "Client error: " << e.what() << "\n"; }
}
