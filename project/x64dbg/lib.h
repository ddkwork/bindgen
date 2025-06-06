#define package struct

package math_ {
     auto Add(auto a, auto b) {
        return a + b;
    };
     auto Sub(int a, int b) {
        return a - b;
    }
}math;

int main() {
    int result = math.Add(3, 2);
}
