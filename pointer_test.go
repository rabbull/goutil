package goutil

import (
	"testing"
)

func TestPtr(t *testing.T) {
	if ptr := Ptr[int](1); ptr == nil || *ptr != 1 {
		t.Error("int")
	}
	if ptr := Ptr[int8](1); ptr == nil || *ptr != 1 {
		t.Error("int8")
	}
	if ptr := Ptr[int16](1); ptr == nil || *ptr != 1 {
		t.Error("int16")
	}
	if ptr := Ptr[int32](1); ptr == nil || *ptr != 1 {
		t.Error("int32")
	}
	if ptr := Ptr[int64](1); ptr == nil || *ptr != 1 {
		t.Error("int64")
	}

	if ptr := Ptr[uint](1); ptr == nil || *ptr != 1 {
		t.Error("uint")
	}
	if ptr := Ptr[uint8](1); ptr == nil || *ptr != 1 {
		t.Error("uint8")
	}
	if ptr := Ptr[uint16](1); ptr == nil || *ptr != 1 {
		t.Error("uint16")
	}
	if ptr := Ptr[uint32](1); ptr == nil || *ptr != 1 {
		t.Error("uint32")
	}
	if ptr := Ptr[uint64](1); ptr == nil || *ptr != 1 {
		t.Error("uint64")
	}
}
