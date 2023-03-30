#include <pcre++.h>
#include <pcrecpp.h>

#include <iostream>
#include <string>
extern "C" {
void pcpp() {
  std::string str = "Hello (world)!";
  std::string replace_str = "<\\1>";
  std::string patten = "\\((.*?)\\)";
  pcrecpp::RE re(patten);

  pcrecpp::StringPiece pstr(replace_str);
  re.GlobalReplace(pstr, &str);
  std::cout << str << std::endl;
}
void p_pp() {
  std::string str = "Hello (world)!";
  std::string replace_str = "<$1>";
  std::string patten = "\\((.*?)\\)";
  pcrepp::Pcre re(patten);
  str = re.replace(str, replace_str);
  std::cout << str << std::endl;
}
}