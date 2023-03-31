#include <iostream>

extern "C" {
    void hello() {
        std::cout << "Hello, world!" << std::endl;
    }
}
