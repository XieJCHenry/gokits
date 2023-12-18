package set

import (
	"reflect"
	"sort"
	"testing"
)

func Test_Set_NewFrom(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 8}

	s := NewFrom[int](elems...)
	got := s.ToArray()

	sort.Slice(got, func(i, j int) bool {
		return got[i] < got[j]
	})

	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want = %v, got = %v", want, got)
	}

}

func Test_Set_InsertAndRemove(t *testing.T) {
	s := New[int]()

	ret := s.Contains(1)
	if ret {
		t.Fatalf("set is empty, has no element")
	}

	ret = s.PutIfAbsent(4)
	if !ret {
		t.Fatalf("set is empty, insert element must return true")
	}
	ret = s.Contains(4)
	if !ret {
		t.Fatalf("set must contains '%d'", 4)
	}

	ret = s.RemoveIfPresent(4)
	if !ret {
		t.Fatalf("elem '%d' is already exist, remove must retrun true", 4)
	}
}
