package types

import "testing"

func Test_ArrayType(t *testing.T) {

	var nums interface{} = []int{1, 2, 3, 4}

	if _, ok := nums.([]interface{}); ok {
		t.Logf("convert success")
	} else {
		t.Errorf("convert failed")
	}
}
