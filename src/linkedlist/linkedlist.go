package linkedlist

// #cgo CFLAGS: -g -Wall
// #include "linkedlist.h"
// #include "stdlib.h"
import "C"
import (
	"unsafe"
)

func Put(key, value int) {
	C.put(C.int(key), C.int(value))
}

func GetBy(key int) int {
	return int(C.get(C.int(key)))
}

func GetAllValues() []int {
	length := int(C.length())
	intPointer := C.getAllValues()
	defer C.free(unsafe.Pointer(intPointer))

	slice := (*[1 << 28]C.int)(unsafe.Pointer(intPointer))[:length:length]

	var values []int
	for index := 0; index < length; index++ {
		values = append(values, int(slice[index]))
	}
	return values
}

func Close() {
	C.close()
}
