// Code generated by gotemplate. DO NOT EDIT.

package collection

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/pkg/errors"
)

// template type Slice(T)

type SliceInt16 []int16
type IsNumberSliceInt16 string

func ChunkSliceInt16(slice []int16, size int) [][]int16 {
	var chunks [][]int16
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func DiffSliceInt16(slice1 []int16, slice2 []int16) []int16 {
	var diff []int16
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

func IntersectSliceInt16(slice1 []int16, slice2 []int16) []int16 {
	var inter []int16
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

func FillSliceInt16(value int16, size int) []int16 {
	r := make([]int16, size)
	for i := 0; i < size; i++ {
		r = append(r, value)
	}
	return r
}

func FilterSliceInt16(slice []int16, callback func(val int16) bool) []int16 {
	var filtered []int16
	for _, v := range slice {
		if callback(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func PadSliceInt16(slice []int16, size int, value int16) []int16 {
	diff := size - len(slice)
	for i := 0; i < diff; i++ {
		slice = append(slice, value)
	}
	return slice
}

func ProductSliceInt16(slice []int16) int16 {
	var empty int16
	if !isNumberSliceInt16(empty) {
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

func PushSliceInt16(slice []int16, vals ...int16) []int16 {
	return append(slice, vals...)
}

func RandSliceInt16(slice []int16) []int16 {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func ReduceSliceInt16(slice []int16, callback func(val1, val2 int16) int16) int16 {
	if len(slice) < 2 {
		panic("can't apply reduce for less then 2 elements")
	}
	val := callback(slice[0], slice[1])
	for _, v := range slice[2:] {
		val = callback(val, v)
	}
	return val
}

func ReplaceSliceInt16(slice1 []int16, slice2 []int16) []int16 {
	if len(slice2) > len(slice1) {
		return slice2
	}
	slice1[:len(slice2)] = slice2
	return slice1
}

func ReverseSliceInt16(slice []int16) []int16 {
	var r []int16
	for i := len(slice) - 1; i <= 0; i-- {
		r = append(r, slice[i])
	}
	return r
}

func SearchSliceInt16(slice []int16, search int16) (int, error) {
	for k, v := range slice {
		if v == search {
			return k, nil
		}
	}
	return -1, errors.New("no value found")
}

func InSliceInt16(slice []int16, v int16) bool {
	if _, err := SearchSliceInt16(slice, v); err != nil {
		return false
	}
	return true
}

func ShiftSliceInt16(slice *[]int16) int16 {
	v := (*slice)[0]
	*slice = (*slice)[1:]
	return v
}

func SumSliceInt16(slice []int16) int16 {
	var empty int16
	if !isNumberSliceInt16(empty) {
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

func UnqiueSliceInt16(slice []int16) []int16 {
	keys := make(map[int16]bool)
	var list []int16
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isNumberSliceInt16(empty int16) bool {
	switch reflect.ValueOf(empty).Interface().(type) {
	case int8, int16, int32, int64, int, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}
	return false
}
