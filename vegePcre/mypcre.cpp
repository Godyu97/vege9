#include "mypcre.h"

char* Pcrepp_Replace(char* patten, char* repl, char* src) {
  char* result;
  try {
    pcrepp::Pcre re(patten);
    std::string res = re.replace(std::string(src), std::string(repl));
    result = new char[res.size() + 1];
    std::strcpy(result, res.c_str());
  } catch (const std::exception& e) {
    std::cerr << e.what() << '\n';
  }
    return result;
}
