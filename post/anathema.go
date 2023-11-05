package badgo

import (
	"fmt"
	"unsafe"
)

func run() {
	var limited int
	var out int
	doThing(out, func(out interface{}) {
		if limited == 0 {
			fmt.Printf("limited is zero. %d\n", limited)
		}
		limited++
	})
}

//go:noinline
func doThing(out interface{}, f func(out interface{})) {
	p := (*eface)(unsafe.Pointer(&out)).data
	*(*int)(p) = 42
	f(out)
}

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}
