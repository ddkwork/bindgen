
#include "pch.h"

int CommunicationClientConnectToServer(PCSTR Ip, PCSTR Port,
                                       SOCKET *ConnectSocketArg) {
  WSADATA wsaData;
  SOCKET ConnectSocket = INVALID_SOCKET;
  struct addrinfo *result = NULL, *ptr = NULL, hints;
  int iResult;
  iResult = WSAStartup(MAKEWORD(2, 2), &wsaData);
  if (iResult != 0) {
    ShowMessages("err, WSAStartup failed (%x)\n", iResult);
    return 1;
  }
  ZeroMemory(&hints, sizeof(hints));
  hints.ai_family = AF_UNSPEC;
  hints.ai_socktype = SOCK_STREAM;
  hints.ai_protocol = IPPROTO_TCP;
  iResult = getaddrinfo(Ip, Port, &hints, &result);
  if (iResult != 0) {
    ShowMessages("getaddrinfo failed (%x)\n", iResult);
    WSACleanup();
    return 1;
  }
  for (ptr = result; ptr != NULL; ptr = ptr->ai_next) {
    ConnectSocket = socket(ptr->ai_family, ptr->ai_socktype, ptr->ai_protocol);
    if (ConnectSocket == INVALID_SOCKET) {
      ShowMessages("socket failed with error: %ld\n", WSAGetLastError());
      WSACleanup();
      return 1;
    }
    iResult = connect(ConnectSocket, ptr->ai_addr, (int)ptr->ai_addrlen);
    if (iResult == SOCKET_ERROR) {
      closesocket(ConnectSocket);
      ConnectSocket = INVALID_SOCKET;
      continue;
    }
    break;
  }
  freeaddrinfo(result);
  if (ConnectSocket == INVALID_SOCKET) {
    ShowMessages("unable to connect to the server\n");
    WSACleanup();
    return 1;
  }
  *ConnectSocketArg = ConnectSocket;
  return 0;
}

int CommunicationClientSendMessage(SOCKET ConnectSocket, const char *sendbuf,
                                   int buflen) {
  int iResult;
  iResult = send(ConnectSocket, sendbuf, buflen, 0);
  if (iResult == SOCKET_ERROR) {
    ShowMessages("err, send failed (%x)\n", WSAGetLastError());
    closesocket(ConnectSocket);
    WSACleanup();
    return 1;
  }
  return 0;
}

int CommunicationClientShutdownConnection(SOCKET ConnectSocket) {
  int iResult;
  iResult = shutdown(ConnectSocket, SD_SEND);
  if (iResult == SOCKET_ERROR) {
    closesocket(ConnectSocket);
    WSACleanup();
    return 1;
  }
  return 0;
}

int CommunicationClientReceiveMessage(SOCKET ConnectSocket, CHAR *RecvBuf,
                                      UINT32 MaxBuffLen, PUINT32 BuffLenRecvd) {
  int Result;
  Result = recv(ConnectSocket, RecvBuf, MaxBuffLen, 0);
  if (Result > 0) {
    *BuffLenRecvd = Result;

  } else if (Result == 0) {
  } else {
    ShowMessages("\nrecv failed with error: %d\n", WSAGetLastError());
    ShowMessages("the remote system closes the connection.\n\n");
    return 1;
  }
  return 0;
}

int CommunicationClientCleanup(SOCKET ConnectSocket) {
  closesocket(ConnectSocket);
  WSACleanup();
  return 0;
}
