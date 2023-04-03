#ifndef MYPCRE_H
#define MYPCRE_H
#include <pcre++.h>

#include <cstring>
#include <iostream>

extern "C" {
std::string Pcrepp_Replace(std::string patten, std::string repl,
                           std::string src);
}

#endif