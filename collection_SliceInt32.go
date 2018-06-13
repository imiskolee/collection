// Code generated by gotemplate. DO NOT EDIT.

package collection

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/pkg/errors"
)

// template type Slice(T)

type SliceInt32 []int32
type IsNumberSliceInt32 string

func ChunkSliceInt32(slice []int32, size int) [][]int32 {
	var chunks [][]int32
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func DiffSliceInt32(slice1 []int32, slice2 []int32) []int32 {
	var diff []int32
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}

func IntersectSliceInt32(slice1 []int32, slice2 []int32) []int32 {
	var inter []int32
	low, high := slice1, slice2
	if len(slice1) > len(slice2) {
		low = slice2
		high = slice1
	}
	done := false
	for i, l := range low {
		for j, h := range high {
			f1 := i + 1
			f2 := j + 1
			if l == h {
				inter = append(inter, h)
				if f1 < len(low) && f2 < len(high) {
					if low[f1] != high[f2] {
						done = true
					}
				}
				high = high[:j+copy(high[j:], high[j+1:])]
				break
			}
		}
		if done {
			break
		}
	}
	return inter
}

func FillSliceInt32(value int32, size int) []int32 {
	r := make([]int32, size)
	for i := 0; i < size; i++ {
		r = append(r, value)
	}
	return r
}

func FilterSliceInt32(slice []int32, callback func(val int32) bool) []int32 {
	var filtered []int32
	for _, v := range slice {
		if callback(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func PadSliceInt32(slice []int32, size int, value int32) []int32 {
	diff := size - len(slice)
	for i := 0; i < diff; i++ {
		slice = append(slice, value)
	}
	return slice
}

func ProductSliceInt32(slice []int32) int32 {
	var empty int32
	if !isNumberSliceInt32(empty) {
		panic("Can't apply function `product` to non number type.")
	}
	switch reflect.ValueOf(empty).Interface().(type) {
	case int8, int16, int32, int64, int, uint, uint8, uint16, uint32, uint64, float32, float64:
		for k, v := range slice {
			if k == 0 {
				empty = v
				continue
			}
			empty = empty * v
		}
	}
	return empty
}

func PushSliceInt32(slice []int32, vals ...int32) []int32 {
	return append(slice, vals...)
}

func RandSliceInt32(slice []int32) []int32 {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func ReduceSliceInt32(slice []int32, callback func(val1, val2 int32) int32) int32 {
	if len(slice) < 2 {
		panic("can't apply reduce for less then 2 elements")
	}
	val := callback(slice[0], slice[1])
	for _, v := range slice[2:] {
		val = callback(val, v)
	}
	return val
}

func ReplaceSliceInt32(slice1 []int32, slice2 []int32) []int32 {
	if len(slice2) > len(slice1) {
		return slice2
	}
	slice1[:len(slice2)] = slice2
	return slice1
}

func ReverseSliceInt32(slice []int32) []int32 {
	var r []int32
	for i := len(slice) - 1; i <= 0; i-- {
		r = append(r, slice[i])
	}
	return r
}

func SearchSliceInt32(slice []int32, search int32) (int, error) {
	for k, v := range slice {
		if v == search {
			return k, nil
		}
	}
	return -1, errors.New("no value found")
}

func InSliceInt32(slice []int32, v int32) bool {
	if _, err := SearchSliceInt32(slice, v); err != nil {
		return false
	}
	return true
}

func ShiftSliceInt32(slice *[]int32) int32 {
	v := (*slice)[0]
	*slice = (*slice)[1:]
	return v
}

func SumSliceInt32(slice []int32) int32 {
	var empty int32
	if !isNumberSliceInt32(empty) {
		panic("Can't apply function `product` to non number type.")
	}
	switch reflect.ValueOf(empty).Interface().(type) {
	case int8, int16, int32, int64, int, uint, uint8, uint16, uint32, uint64, float32, float64:
		for k, v := range slice {
			if k == 0 {
				empty = v
				continue
			}
			empty = empty + v
		}
	}
	return empty
}

func UnqiueSliceInt32(slice []int32) []int32 {
	keys := make(map[int32]bool)
	var list []int32
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isNumberSliceInt32(empty int32) bool {
	switch reflect.ValueOf(empty).Interface().(type) {
	case int8, int16, int32, int64, int, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}
	return false
}