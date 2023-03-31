#ifndef MYPCRE_H
#define MYPCRE_H
#include <pcre++.h>

#include <cstring>
#include <iostream>

extern "C" {
void Pcrepp_Replace(char* patten, char* repl, char*& src);
}

#endif