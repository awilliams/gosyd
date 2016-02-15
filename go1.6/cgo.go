package main

/*
int fn(void* arg) { return 42; }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type T struct{ a, b int }
type X struct{ t *T }

func main() {
	t := T{a: 1, b: 2}
	fmt.Println(C.fn(unsafe.Pointer(&t)))

	//x := X{t: &t} // Holds pointer to x // HL
	//fmt.Println(C.fn(unsafe.Pointer(&x)))
}
