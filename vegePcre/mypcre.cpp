
#include <pcre++.h>
#include <cstring>
#include <iostream>
extern "C"{

void Pcrepp_Replace(char* patten, char* repl, char* src) {
  //   std::string str = "Hello (world)!";
  //   std::string replace_str = "<$1>";
  //   std::string patten = "\\((.*?)\\)";
  pcrepp::Pcre re(patten);
  std::string res = re.replace(std::string(src), std::string(repl));
  std::strcpy(src, res.c_str());
}
}