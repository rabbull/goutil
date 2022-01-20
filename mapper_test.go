package goutil_test

import (
	"testing"

	"github.com/rabbull/goutil"
)

func TestMapper(t *testing.T) {
	i := goutil.IntAdder(1)
	mapper := goutil.NewAdderMapper(&i)

	seq := []string{"a", "b", "c", "d", "a", "c", "e", "f", "a", "b"}
	exp := []int{1, 2, 3, 4, 1, 3, 5, 6, 1, 2}
	for i, k := range seq {
		v := int(*mapper.Feed(k).(*goutil.IntAdder))
		if v != exp[i] {
			t.Errorf("expect=%v, got=%v", exp[i], v)
		}
	}
}
