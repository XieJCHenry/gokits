// Package slices contains some helper functions for built-in slice
package slices

import (
	"github.com/XieJCHenry/gokits/collections/set"
	"github.com/XieJCHenry/gokits/collections/tuple"
	"math"
	"math/rand"
)

func ContainsV2[T comparable](arr []T, x T) bool {
	return Contains(arr, x)
}

func Contains[T comparable](arr []T, x T) bool {
	return IndexOf(arr, x) != -1
}

func IndexOf[T comparable](arr []T, x T) int {
	idx := -1
	for i := range arr {
		if arr[i] == x {
			idx = i
			break
		}
	}
	return idx
}

func Remove[T comparable](arr []T, x T) []T {
	idx := IndexOf(arr, x)
	if idx != -1 {
		arr = append(arr[0:idx], arr[idx+1:]...)
	}
	return arr
}

func RemoveAt[T comparable](arr []T, index int) (newArr []T, val T) {
	if index < 0 || index >= len(arr) {
		return arr, val
	}
	x := arr[index]
	arr = Remove(arr, x)
	return arr, x
}

func ToSet[T comparable](arr []T) set.Set[T] {
	return set.NewFrom[T](arr...)
}

func ToTupleList(arr1 []interface{}, arr2 []interface{}) []tuple.Tuple {
	minSize := int(math.Min(float64(len(arr1)), float64(len(arr2))))
	result := make([]tuple.Tuple, 0, minSize)
	for i := 0; i < minSize; i++ {
		t := tuple.New(2)
		t.Set(0, arr1[i])
		t.Set(1, arr2[i])
		result[i] = t
	}
	return result
}

func ToMap[T comparable](keys []T, vals []interface{}) map[T]interface{} {
	result := make(map[T]interface{})
	minSize := int(math.Min(float64(len(keys)), float64(len(vals))))

	for i := 0; i < minSize; i++ {
		key := keys[i]
		val := vals[i]
		// 仅不存在时插入，存在重复时略过
		if _, ok := result[key]; !ok {
			result[key] = val
		}
	}
	return result
}

func Copy[T any](arr []T) []T {
	result := make([]T, 0, len(arr))
	copy(result, arr)
	return result
}

func Swap[T any](arr []T, x int, y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}

func Map[T any](arr []T, mapFunc func(v T) T) []T {
	for i := range arr {
		arr[i] = mapFunc(arr[i])
	}
	return arr
}

func Filter[T any, U any](arr []T, filterFunc func(v T) (U, bool)) []U {
	resultList := make([]U, 0, len(arr))
	for i := range arr {
		t := arr[i]
		if u, ok := filterFunc(t); ok {
			resultList = append(resultList, u)
		}
	}
	return resultList
}

func GroupBy[K comparable, T any](arr []T, keyFunc func(elem T) (K, bool)) map[K][]T {
	result := make(map[K][]T)
	for _, elem := range arr {
		if key, ok := keyFunc(elem); ok {
			if _, ok := result[key]; !ok {
				result[key] = make([]T, 0)
			}
			result[key] = append(result[key], elem)
		}
	}
	return result
}

func Uniq[T comparable](arr []T) []T {
	return ToSet(arr).ToArray()
}

func Shuffle[T comparable](arr []T) []T {
	result := Copy(arr)
	times := len(result) / 2
	var pa, pb int = 0, len(result) - 1
	for i := 0; i < times; i++ {
		for j := 0; j < 3; j++ {
			pa = rand.Intn(len(result))
			pb = rand.Intn(len(result))
			if pa != pb {
				break
			}
		}
		Swap(arr, pb, pb)
	}
	return result
}
