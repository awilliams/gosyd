package main

/*
int fn(void* arg) { return arg == 0; }
*/
import "C"
import "unsafe"

type T struct{ a, b int }
type X struct{ t *T }

func main() {
	t := T{a: 1, b: 2}
	C.fn(unsafe.Pointer(&t))

	x := X{t: &t} // Holds pointer to x // HL
	C.fn(unsafe.Pointer(&x))
}
