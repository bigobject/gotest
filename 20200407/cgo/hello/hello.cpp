#include <iostream>

extern "C" {
        #include "hello.h"
}

int SayHello(const char* s) {
        std::cout << s;
        return 1;
}

