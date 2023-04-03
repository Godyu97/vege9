/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: mypcre.i

package vegePcre

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
extern void _wrap_Swig_free_vegePcre_f83a3cf6dbc14c45(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_vegePcre_f83a3cf6dbc14c45(swig_intgo arg1);
extern swig_type_1 _wrap_Pcrepp_Replace_vegePcre_f83a3cf6dbc14c45(swig_type_2 arg1, swig_type_3 arg2, swig_type_4 arg3, swig_type_5 arg4);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_vegePcre_f83a3cf6dbc14c45(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_vegePcre_f83a3cf6dbc14c45(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func Pcrepp_Replace(arg1 string, arg2 string, arg3 string, arg4 string) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r_p := C._wrap_Pcrepp_Replace_vegePcre_f83a3cf6dbc14c45(*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_4)(unsafe.Pointer(&_swig_i_2)), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_3)))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg4
	}
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}


