#include "mypcre.h"

char* Pcrepp_Replace(char* patten, char* repl, char* src) {
  pcrepp::Pcre re(patten);
  std::string res = re.replace(std::string(src), std::string(repl));
  char* result = new char[res.size() + 1];
  std::strcpy(result, res.c_str());
  return result;
}
