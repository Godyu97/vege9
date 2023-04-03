#include "mypcre.h"

std::string Pcrepp_Replace(std::string patten, std::string repl,
                           std::string src) {
  pcrepp::Pcre re(patten);
  return re.replace(src, repl);
}
