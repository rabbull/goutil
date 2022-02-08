package goutil

import (
	"reflect"
	"unsafe"
)

const (
	sizeOfInt = int(unsafe.Sizeof(int(0)))
)

func MarshalIntSlice(slice []int) []byte {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	result := make([]byte, hdr.Len*sizeOfInt)
	copy(result, *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Cap:  len(result),
		Len:  len(result),
		Data: hdr.Data,
	})))
	return result
}

func UnmarshalIntSlice(data []byte) []int {
	if len(data)%sizeOfInt != 0 {
		return nil
	}
	result := make([]int, len(data)/sizeOfInt)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&result))
	copy(*(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Cap:  hdr.Len * sizeOfInt,
		Len:  hdr.Len * sizeOfInt,
		Data: hdr.Data,
	})), data)
	return result
}
