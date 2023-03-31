#include <pcre++.h>
#include <pcrecpp.h>

#include <iostream>
#include <string>
extern "C"{
void Pcrecpp_GlobalReplace(char* patten, char* repl, char* src) {
  pcrecpp::RE re(patten);
  pcrecpp::StringPiece pstr(repl);
  std::string res = std::string(src);
  re.GlobalReplace(pstr, &res);
  std::strcpy(src, res.c_str());
}

void Pcrepp_Replace(char* patten, char* repl, char* src) {
  pcrepp::Pcre re(patten);
  std::string res = re.replace(std::string(src), std::string(repl));
  std::strcpy(src, res.c_str());
}
}