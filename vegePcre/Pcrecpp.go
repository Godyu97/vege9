package vegePcre

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -lpcre++ -lpcrecpp
#cgo CFLAGS: -I/opt/local/include
#include <mypcre.h>
*/
import "C"

func Replace(patten string, src string, repl string) string {
	C.pcpp()
	return ""
}
