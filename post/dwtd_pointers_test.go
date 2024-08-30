package badgo

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"sync"
	"testing"
	"unsafe"
)

func TestRandomUnsafePointers(t *testing.T) {
	x := make([]unsafe.Pointer, 1e9)

	for i := range x {
		// Possible misuse of unsafe.Pointer? Definite misuse of unsafe.Pointer!
		x[i] = unsafe.Pointer(uintptr(i * 8))
	}

	runtime.GC()

	runtime.KeepAlive(x)
}

func TestRandomUnsafePointers2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := make([]unsafe.Pointer, 1e9)

		for i := range x {
			// Possible misuse of unsafe.Pointer? Definite misuse of unsafe.Pointer!
			x[i] = unsafe.Pointer(uintptr(rand.Int64()))
		}

		runtime.GC()

		for range 10 {
			for i := range x {
				// Possible misuse of unsafe.Pointer? Definite misuse of unsafe.Pointer!
				x[i] = unsafe.Add(x[i], 3)
			}

			runtime.GC()
		}
		runtime.KeepAlive(x)
	}()
	runtime.GC()
	wg.Wait()

	runtime.GC()
}

func TestUnsafePointerBadNumber(t *testing.T) {
	y := make([]int, 1e4)
	yptr := uintptr(unsafe.Pointer(unsafe.SliceData(y)))
	y = nil
	runtime.GC()
	runtime.GC()
	x := unsafe.Pointer(yptr)
	runtime.GC()
	runtime.GC()

	runtime.KeepAlive(x)
}

func TestUnsafePointersGC(t *testing.T) {
	var initialStats runtime.MemStats
	runtime.ReadMemStats(&initialStats)

	x := make([]unsafe.Pointer, 1e5)

	type thing struct {
		val *int
	}

	for i := range x {
		val := 100000 + i
		// Possible misuse of unsafe.Pointer? Definite misuse of unsafe.Pointer!
		x[i] = unsafe.Pointer(&thing{val: &val})
	}

	runtime.GC()
	runtime.GC()
	runtime.GC()

	var finalStats runtime.MemStats
	runtime.ReadMemStats(&finalStats)

	fmt.Printf("Allocated: %d\n", finalStats.HeapObjects-initialStats.HeapObjects)

	runtime.KeepAlive(x)
}
