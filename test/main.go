package main

// #cgo LDFLAGS: -L. -lhello
// void hello();
import "C"

func main() {
	C.hello()
}
