#ifndef MYPCRE_H
#define MYPCRE_H
#include <pcre++.h>

#include <cstring>
#include <exception>
#include <iostream>

extern "C" {
char* Pcrepp_Replace(char* patten, char* repl, char* src);
}

#endif