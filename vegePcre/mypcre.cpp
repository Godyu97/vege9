

#include "mypcre.h"


void Pcrepp_Replace(char* patten, char* repl, char*& src) {
  pcrepp::Pcre re(patten);
  std::string res = re.replace(std::string(src), std::string(repl));
  src = new char[res.size() + 1];
  std::strcpy(src, res.c_str());
}
