package main

// #cgo CPPFLAGS: -std=c++11 -I/usr/include/c++/9
// #cgo LDFLAGS: -L. -ltest
// void hello();
import "C"

func main() {
	C.hello()
}
