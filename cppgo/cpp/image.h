#pragma once
#include <cstdint>
#include <vector>
#include <memory>
#include <algorithm>
#include <cmath>
#include <stdexcept>
#include <string>
#include <map>
#include <functional>

using uint8 = uint8_t;
using uint16 = uint16_t;
using uint32 = uint32_t;
using int32 = int32_t;

struct Point {
    int X;
    int Y;

    Point() : X(0), Y(0) {}

    Point(int x, int y) : X(x), Y(y) {}

    Point Add(Point p) const { return Point(X + p.X, Y + p.Y); }

    Point Sub(Point p) const { return Point(X - p.X, Y - p.Y); }

    Point Mul(int k) const { return Point(X * k, Y * k); }

    Point Div(int k) const { return Point(X / k, Y / k); }

    bool Eq(Point p) const { return X == p.X && Y == p.Y; }

    bool In(Rectangle r) const;

    int Distance(Point p) const {
        int dx = X - p.X;
        int dy = Y - p.Y;
        return static_cast<int>(std::sqrt(dx * dx + dy * dy));
    }
};

struct Rectangle {
    Point Min;
    Point Max;

    Rectangle() {}

    Rectangle(Point min, Point max) : Min(min), Max(max) {}

    Rectangle(int x0, int y0, int x1, int y1) : Min(Point(x0, y0)), Max(Point(x1, y1)) {}

    Point Dx() const { return Point(Dx(), Dy()); }

    int Dx() const { return Max.X - Min.X; }

    int Dy() const { return Max.Y - Min.Y; }

    bool Empty() const { return Min.X >= Max.X || Min.Y >= Max.Y; }

    bool Eq(Rectangle r) const { return Min.Eq(r.Min) && Max.Eq(r.Max); }

    Rectangle Canon() const {
        if (Min.X > Max.X) std::swap(Min.X, Max.X);
        if (Min.Y > Max.Y) std::swap(Min.Y, Max.Y);
        return *this;
    }

    Rectangle Inset(int d) const { return Rectangle(Point(Min.X + d, Min.Y + d), Point(Max.X - d, Max.Y - d)); }

    Rectangle Intersect(Rectangle r) const {
        if (Min.X < r.Min.X) Min.X = r.Min.X;
        if (Min.Y < r.Min.Y) Min.Y = r.Min.Y;
        if (Max.X > r.Max.X) Max.X = r.Max.X;
        if (Max.Y > r.Max.Y) Max.Y = r.Max.Y;
        if (Empty()) return Rectangle();
        return *this;
    }

    Rectangle Union(Rectangle r) const {
        if (Empty()) return r;
        if (r.Empty()) return *this;
        if (Min.X > r.Min.X) Min.X = r.Min.X;
        if (Min.Y > r.Min.Y) Min.Y = r.Min.Y;
        if (Max.X < r.Max.X) Max.X = r.Max.X;
        if (Max.Y < r.Max.Y) Max.Y = r.Max.Y;
        return *this;
    }

    bool Overlaps(Rectangle r) const { return !Empty() && !r.Empty() && Min.X < r.Max.X && Max.X > r.Min.X && Min.Y < r.Max.Y && Max.Y > r.Min.Y; }

    bool Contains(Point p) const { return p.X >= Min.X && p.X < Max.X && p.Y >= Min.Y && p.Y < Max.Y; }
};

bool Point::In(Rectangle r) const { return r.Contains(*this); }

struct Color {
    uint8 R, G, B, A;

    Color() : R(0), G(0), B(0), A(255) {}

    Color(uint8 r, uint8 g, uint8 b) : R(r), G(g), B(b), A(255) {}

    Color(uint8 r, uint8 g, uint8 b, uint8 a) : R(r), G(g), B(b), A(a) {}

    static Color RGBA(uint8 r, uint8 g, uint8 b, uint8 a) { return Color(r, g, b, a); }

    static Color RGB(uint8 r, uint8 g, uint8 b) { return Color(r, g, b); }

    uint32 RGBA() const { return (static_cast<uint32>(R) << 24) | (static_cast<uint32>(G) << 16) | (static_cast<uint32>(B) << 8) | static_cast<uint32>(A); }

    Color Blend(Color over) const {
        if (over.A == 0) return *this;
        if (over.A == 255) return over;

        float alpha = static_cast<float>(over.A) / 255.0f;
        float invAlpha = 1.0f - alpha;

        uint8 r = static_cast<uint8>(over.R * alpha + R * invAlpha);
        uint8 g = static_cast<uint8>(over.G * alpha + G * invAlpha);
        uint8 b = static_cast<uint8>(over.B * alpha + B * invAlpha);
        uint8 a = static_cast<uint8>(A + over.A * (255 - A) / 255);

        return Color(r, g, b, a);
    }

    uint8 Y() const { return static_cast<uint8>(0.299f * R + 0.587f * G + 0.114f * B); }

    bool Eq(Color c) const { return R == c.R && G == c.G && B == c.B && A == c.A; }
};

struct ColorModel {
    virtual Color Convert(Color c) const = 0;
};

struct RGBA64Model : ColorModel {
    Color Convert(Color c) const override { return Color(static_cast<uint8>(c.R), static_cast<uint8>(c.G), static_cast<uint8>(c.B), static_cast<uint8>(c.A)); }
};

struct NRGBAModel : ColorModel {
    Color Convert(Color c) const override {
        float alpha = c.A / 255.0f;
        return Color(static_cast<uint8>(c.R * alpha), static_cast<uint8>(c.G * alpha), static_cast<uint8>(c.B * alpha), c.A);
    }
};

struct GrayModel : ColorModel {
    Color Convert(Color c) const override {
        uint8 y = c.Y();
        return Color(y, y, y, c.A);
    }
};

struct Gray16Model : ColorModel {
    Color Convert(Color c) const override {
        uint8 y = c.Y();
        return Color(y, y, y, c.A);
    }
};

struct CMYKModel : ColorModel {
    Color Convert(Color c) const override {
        float r = c.R / 255.0f;
        float g = c.G / 255.0f;
        float b = c.B / 255.0f;

        float k = 1.0f - std::max({r, g, b});
        float cyan = (1.0f - r - k) / (1.0f - k);
        float magenta = (1.0f - g - k) / (1.0f - k);
        float yellow = (1.0f - b - k) / (1.0f - k);

        return Color(static_cast<uint8>(cyan * 255), static_cast<uint8>(magenta * 255), static_cast<uint8>(yellow * 255), static_cast<uint8>(k * 255));
    }
};

RGBA64Model rgba64Model;
NRGBAModel  nrgbaModel;
GrayModel   grayModel;
Gray16Model gray16Model;
CMYKModel   cmykModel;

struct Image {
    virtual ColorModel ColorModel() const = 0;
    virtual Rectangle  Bounds() const = 0;
    virtual Color      At(int x, int y) const = 0;
    virtual void       Set(int x, int y, Color c) = 0;
};

struct PalettedImage : Image {
    virtual std::vector<Color> Palette() const = 0;
    virtual uint8              IndexAt(int x, int y) const = 0;
    virtual void               SetIndex(int x, int y, uint8 index) = 0;
};

struct RGBA : Image {
    Rectangle          rect;
    std::vector<uint8> pix;// R,G,B,A 顺序
    int                stride;

    RGBA() : stride(0) {}

    RGBA(Rectangle r) : rect(r) {
        int w = r.Dx();
        int h = r.Dy();
        stride = w * 4;
        pix.resize(stride * h);
    }

    ColorModel ColorModel() const override { return rgba64Model; }

    Rectangle Bounds() const override { return rect; }

    Color At(int x, int y) const override {
        if (!Point(x, y).In(rect)) return Color();

        int i = (y - rect.Min.Y) * stride + (x - rect.Min.X) * 4;
        return Color(pix[i], pix[i + 1], pix[i + 2], pix[i + 3]);
    }

    void Set(int x, int y, Color c) override {
        if (!Point(x, y).In(rect)) return;

        int i = (y - rect.Min.Y) * stride + (x - rect.Min.X) * 4;
        pix[i] = c.R;
        pix[i + 1] = c.G;
        pix[i + 2] = c.B;
        pix[i + 3] = c.A;
    }

    RGBA SubImage(Rectangle r) {
        r = r.Intersect(rect);
        if (r.Empty()) return RGBA();

        RGBA sub(r);
        for (int y = r.Min.Y; y < r.Max.Y; y++) {
            for (int x = r.Min.X; x < r.Max.X; x++) { sub.Set(x, y, At(x, y)); }
        }
        return sub;
    }

    RGBA ToGray() {
        RGBA gray(rect);
        for (int y = rect.Min.Y; y < rect.Max.Y; y++) {
            for (int x = rect.Min.X; x < rect.Max.X; x++) {
                Color c = At(x, y);
                uint8 y = c.Y();
                gray.Set(x, y, Color(y, y, y, c.A));
            }
        }
        return gray;
    }
};

struct Gray : Image {
    Rectangle          rect;
    std::vector<uint8> pix;// 灰度值
    int                stride;

    Gray() : stride(0) {}

    Gray(Rectangle r) : rect(r) {
        int w = r.Dx();
        int h = r.Dy();
        stride = w;
        pix.resize(stride * h);
    }

    ColorModel ColorModel() const override { return grayModel; }

    Rectangle Bounds() const override { return rect; }

    Color At(int x, int y) const override {
        if (!Point(x, y).In(rect)) return Color();

        int   i = (y - rect.Min.Y) * stride + (x - rect.Min.X);
        uint8 v = pix[i];
        return Color(v, v, v);
    }

    void Set(int x, int y, Color c) override {
        if (!Point(x, y).In(rect)) return;

        int i = (y - rect.Min.Y) * stride + (x - rect.Min.X);
        pix[i] = c.Y();
    }
};

struct Paletted : PalettedImage {
    Rectangle          rect;
    std::vector<uint8> pix;// 调色板索引
    std::vector<Color> palette;
    int                stride;

    Paletted() : stride(0) {}

    Paletted(Rectangle r, std::vector<Color> pal) : rect(r), palette(pal) {
        int w = r.Dx();
        int h = r.Dy();
        stride = w;
        pix.resize(stride * h);
    }

    ColorModel ColorModel() const override { return rgba64Model; }

    Rectangle Bounds() const override { return rect; }

    Color At(int x, int y) const override {
        uint8 index = IndexAt(x, y);
        if (index >= palette.size()) return Color();
        return palette[index];
    }

    void Set(int x, int y, Color c) override {
        uint8 bestIndex = 0;
        int   bestDist = 256 * 256 * 3;

        for (uint8 i = 0; i < palette.size(); i++) {
            int dr = c.R - palette[i].R;
            int dg = c.G - palette[i].G;
            int db = c.B - palette[i].B;
            int dist = dr * dr + dg * dg + db * db;

            if (dist < bestDist) {
                bestDist = dist;
                bestIndex = i;
            }
        }

        SetIndex(x, y, bestIndex);
    }

    std::vector<Color> Palette() const override { return palette; }

    uint8 IndexAt(int x, int y) const override {
        if (!Point(x, y).In(rect)) return 0;

        int i = (y - rect.Min.Y) * stride + (x - rect.Min.X);
        return pix[i];
    }

    void SetIndex(int x, int y, uint8 index) override {
        if (!Point(x, y).In(rect)) return;
        if (index >= palette.size()) return;

        int i = (y - rect.Min.Y) * stride + (x - rect.Min.X);
        pix[i] = index;
    }
};

struct Draw {
    static void Rectangle(Image &img, Rectangle r, Color c) {
        int x0 = r.Min.X, y0 = r.Min.Y;
        int x1 = r.Max.X, y1 = r.Max.Y;

        Draw::Line(img, Point(x0, y0), Point(x1, y0), c);// 上边
        Draw::Line(img, Point(x0, y1), Point(x1, y1), c);// 下边
        Draw::Line(img, Point(x0, y0), Point(x0, y1), c);// 左边
        Draw::Line(img, Point(x1, y0), Point(x1, y1), c);// 右边
    }

    static void FillRectangle(Image &img, Rectangle r, Color c) {
        for (int y = r.Min.Y; y < r.Max.Y; y++) {
            for (int x = r.Min.X; x < r.Max.X; x++) { img.Set(x, y, c); }
        }
    }

    static void Line(Image &img, Point p0, Point p1, Color c) {
        int dx = std::abs(p1.X - p0.X);
        int dy = std::abs(p1.Y - p0.Y);
        int sx = p0.X < p1.X ? 1 : -1;
        int sy = p0.Y < p1.Y ? 1 : -1;
        int err = dx - dy;

        while (true) {
            img.Set(p0.X, p0.Y, c);
            if (p0.X == p1.X && p0.Y == p1.Y) break;

            int e2 = 2 * err;
            if (e2 > -dy) {
                err -= dy;
                p0.X += sx;
            }
            if (e2 < dx) {
                err += dx;
                p0.Y += sy;
            }
        }
    }

    static void Circle(Image &img, Point center, int radius, Color c) {
        int x = radius;
        int y = 0;
        int err = 0;

        while (x >= y) {
            img.Set(center.X + x, center.Y + y, c);
            img.Set(center.X + y, center.Y + x, c);
            img.Set(center.X - y, center.Y + x, c);
            img.Set(center.X - x, center.Y + y, c);
            img.Set(center.X - x, center.Y - y, c);
            img.Set(center.X - y, center.Y - x, c);
            img.Set(center.X + y, center.Y - x, c);
            img.Set(center.X + x, center.Y - y, c);

            if (err <= 0) {
                y += 1;
                err += 2 * y + 1;
            }
            if (err > 0) {
                x -= 1;
                err -= 2 * x + 1;
            }
        }
    }

    static void FillCircle(Image &img, Point center, int radius, Color c) {
        for (int y = -radius; y <= radius; y++) {
            for (int x = -radius; x <= radius; x++) {
                if (x * x + y * y <= radius * radius) { img.Set(center.X + x, center.Y + y, c); }
            }
        }
    }

    static void DrawImage(Image &dst, Point pt, Image &src) {
        Rectangle srcRect = src.Bounds();
        for (int y = 0; y < srcRect.Dy(); y++) {
            for (int x = 0; x < srcRect.Dx(); x++) {
                Color c = src.At(x, y);
                if (c.A > 0) { dst.Set(pt.X + x, pt.Y + y, c); }
            }
        }
    }

    static void DrawImageBlend(Image &dst, Point pt, Image &src) {
        Rectangle srcRect = src.Bounds();
        for (int y = 0; y < srcRect.Dy(); y++) {
            for (int x = 0; x < srcRect.Dx(); x++) {
                Color srcColor = src.At(x, y);
                if (srcColor.A == 0) continue;

                Color dstColor = dst.At(pt.X + x, pt.Y + y);
                dst.Set(pt.X + x, pt.Y + y, dstColor.Blend(srcColor));
            }
        }
    }
};

struct ImageFormat {
    virtual std::string              Name() const = 0;
    virtual std::vector<std::string> Extensions() const = 0;
    virtual bool                     Decode(Image &img, const std::vector<uint8> &data) = 0;
    virtual bool                     Encode(const Image &img, std::vector<uint8> &data) = 0;
};

struct PNGFormat : ImageFormat {
    std::string Name() const override { return "PNG"; }

    std::vector<std::string> Extensions() const override { return {"png"}; }

    bool Decode(Image &img, const std::vector<uint8> &data) override {
        if (data.size() < 8) return false;

        if (data[0] != 0x89 || data[1] != 'P' || data[2] != 'N' || data[3] != 'G') { return false; }

        int       width = 32;
        int       height = 32;
        Rectangle rect(0, 0, width, height);

        if (dynamic_cast<RGBA *>(&img)) {
            RGBA &rgba = dynamic_cast<RGBA &>(img);
            rgba = RGBA(rect);

            size_t offset = 8;// 跳过头
            for (int y = 0; y < height; y++) {
                for (int x = 0; x < width; x++) {
                    if (offset + 4 > data.size()) break;
                    rgba.Set(x, y, Color(data[offset], data[offset + 1], data[offset + 2], data[offset + 3]));
                    offset += 4;
                }
            }
            return true;
        }
        return false;
    }

    bool Encode(const Image &img, std::vector<uint8> &data) override {
        data.clear();

        data.push_back(0x89);
        data.push_back('P');
        data.push_back('N');
        data.push_back('G');
        data.push_back(0x0D);
        data.push_back(0x0A);
        data.push_back(0x1A);
        data.push_back(0x0A);

        Rectangle rect = img.Bounds();
        int       width = rect.Dx();
        int       height = rect.Dy();

        for (int y = 0; y < height; y++) {
            for (int x = 0; x < width; x++) {
                Color c = img.At(x, y);
                data.push_back(c.R);
                data.push_back(c.G);
                data.push_back(c.B);
                data.push_back(c.A);
            }
        }
        return true;
    }
};

struct JPEGFormat : ImageFormat {
    std::string Name() const override { return "JPEG"; }

    std::vector<std::string> Extensions() const override { return {"jpg", "jpeg"}; }

    bool Decode(Image &img, const std::vector<uint8> &data) override {
        if (data.size() < 2 || data[0] != 0xFF || data[1] != 0xD8) { return false; }

        int       width = 32;
        int       height = 32;
        Rectangle rect(0, 0, width, height);

        if (dynamic_cast<RGBA *>(&img)) {
            RGBA &rgba = dynamic_cast<RGBA &>(img);
            rgba = RGBA(rect);

            for (int y = 0; y < height; y++) {
                for (int x = 0; x < width; x++) { rgba.Set(x, y, Color(static_cast<uint8>(x * 255 / width), static_cast<uint8>(y * 255 / height), 128)); }
            }
            return true;
        }
        return false;
    }

    bool Encode(const Image &img, std::vector<uint8> &data) override {
        data.clear();

        data.push_back(0xFF);
        data.push_back(0xD8);// SOI

        Rectangle rect = img.Bounds();
        int       width = rect.Dx();
        int       height = rect.Dy();

        for (int y = 0; y < height; y++) {
            for (int x = 0; x < width; x++) {
                Color c = img.At(x, y);
                data.push_back(c.R);
                data.push_back(c.G);
                data.push_back(c.B);
            }
        }

        data.push_back(0xFF);
        data.push_back(0xD9);// EOI
        return true;
    }
};

std::map<std::string, ImageFormat *> formats;

void RegisterFormat(ImageFormat *format) {
    formats[format->Name()] = format;
    for (auto &ext: format->Extensions()) { formats[ext] = format; }
}

void initImageFormats() {
    static bool initialized = false;
    if (initialized) return;

    RegisterFormat(new PNGFormat());
    RegisterFormat(new JPEGFormat());

    initialized = true;
}

RGBA Decode(const std::vector<uint8> &data) {
    initImageFormats();

    if (data.size() < 8) return RGBA();

    for (auto &pair: formats) {
        RGBA img;
        if (pair.second->Decode(img, data)) { return img; }
    }
    return RGBA();
}

bool Encode(const std::string &format, const Image &img, std::vector<uint8> &data) {
    initImageFormats();

    auto it = formats.find(format);
    if (it == formats.end()) return false;

    return it->second->Encode(img, data);
}

void testImageFunctions() {
    Rectangle rect(0, 0, 256, 256);
    RGBA      img(rect);

    for (int y = 0; y < rect.Dy(); y++) {
        for (int x = 0; x < rect.Dx(); x++) { img.Set(x, y, Color(static_cast<uint8>(x), static_cast<uint8>(y), static_cast<uint8>((x + y) / 2), 255)); }
    }

    Draw::FillRectangle(img, Rectangle(50, 50, 100, 100), Color(
