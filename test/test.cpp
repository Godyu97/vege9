#include <iostream>
#include <string>
extern "C" {
    void hello() {
        std::string a="Hello, world!"
        std::cout << a << std::endl;
    }
}
