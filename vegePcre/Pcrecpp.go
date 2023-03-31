package vegePcre

/*
 #cgo CXXFLAGS: -std=c++11 -I/usr/include/c++/9
 #cgo LDFLAGS:  -lmypcre
 void Pcrepp_Replace(char* patten, char* repl, char* src);
*/
import "C"

import (
	"unsafe"
)

func Replace(pattern string, repl string, src string) string {
	pattern1 := C.CString(pattern)
	defer C.free(unsafe.Pointer(pattern1))
	src1 := C.CString(src)
	defer C.free(unsafe.Pointer(src1))
	repl1 := C.CString(repl)
	defer C.free(unsafe.Pointer(repl1))
	C.Pcrepp_Replace(pattern1, repl1, src1)
	return string(src1)
}
