#include <iostream>
#include <thread>

void worker(int id) {
    std::this_thread::sleep_for(std::chrono::milliseconds(100));
    std::cout << "Worker " << id << " completed\n";
}

int main() {
    auto g = Group::New();

    g->SetLimit(3);

    for (int i = 0; i < 10; i++) {
        g->Go([i] { worker(i); });
    }

    if (!g->TryGo([] { worker(11); })) { std::cout << "Failed to start worker 11 (as expected)\n"; }

    g->Wait();
    std::cout << "All tasks completed\n";

    g->SetLimit(-1);

    g->Go([] { worker(12); });
    g->Wait();

    return 0;
}
