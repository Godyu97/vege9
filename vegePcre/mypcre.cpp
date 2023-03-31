#include <pcre++.h>
#include <cstring>
#include <iostream>

void Pcrepp_Replace(char* patten, char* repl, char* src) {
  pcrepp::Pcre re(patten);
  std::string res = re.replace(std::string(src), std::string(repl));
  std::strcpy(src, res.c_str());
}
