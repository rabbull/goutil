package goutil

import (
	"math/rand"
	"testing"
)

func makeRandomSlice(length uint) []int {
	result := make([]int, length)
	for i := 0; i < int(length); i++ {
		result[i] = rand.Int()
	}
	return result
}

func isIntSliceEq(lhs, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	length := len(lhs)
	for i := 0; i < length; i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func TestMarshalIntSlice(t *testing.T) {
	for i := 0; i < 1e3; i++ {
		length := rand.Intn(50) + 50
		from := makeRandomSlice(uint(length))
		data := MarshalIntSlice(from)
		to := UnmarshalIntSlice(data)
		if !isIntSliceEq(from, to) {
			t.Errorf("result mismatch: from=%#v to=%#v", from, to)
		}
	}
}
