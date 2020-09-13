package main

/*
#cgo CFLAGS: -I/usr/local/include/
#cgo LDFLAGS: -L/usr/local/lib -lglpk
#include <glpk.h>
#include <stdio.h>
#include <stdlib.h>
static void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"
import "unsafe"
import (
	"fmt"
)

const ()

var ()

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func CPrint(s string) {
	cs := C.CString(s)
	// A common idiom in cgo programs is to defer the free immediately after allocating
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
	// Calling variadic C functions is not supported.
	// It is possible to circumvent this by using a C function wrapper.
	C.myprint(cs)
}

func main() {
	fmt.Println("hello")
	Seed(12345)
	fmt.Println("random:", Random())
	CPrint("hello2")

	// declare variables
	// glp_prob *lp;
	var lp *C.glp_prob
	//defer C.glp_delete_prob(lp)
	// create problem
	// lp = glp_create_prob();
	lp = C.glp_create_prob()
	lp_name := C.CString("short")
	defer C.free(unsafe.Pointer(lp_name))

	C.glp_set_prob_name(lp, lp_name)
}
