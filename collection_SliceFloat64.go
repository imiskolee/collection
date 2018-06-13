// Code generated by gotemplate. DO NOT EDIT.

package collection

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/pkg/errors"
)

// template type Slice(T)

type SliceFloat64 []float64
type IsNumberSliceFloat64 string

func ChunkSliceFloat64(slice []float64, size int) [][]float64 {
	var chunks [][]float64
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func DiffSliceFloat64(slice1 []float64, slice2 []float64) []float64 {
	var diff []float64
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

func IntersectSliceFloat64(slice1 []float64, slice2 []float64) []float64 {
	var inter []float64
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

func FillSliceFloat64(value float64, size int) []float64 {
	r := make([]float64, size)
	for i := 0; i < size; i++ {
		r = append(r, value)
	}
	return r
}

func FilterSliceFloat64(slice []float64, callback func(val float64) bool) []float64 {
	var filtered []float64
	for _, v := range slice {
		if callback(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func PadSliceFloat64(slice []float64, size int, value float64) []float64 {
	diff := size - len(slice)
	for i := 0; i < diff; i++ {
		slice = append(slice, value)
	}
	return slice
}

func ProductSliceFloat64(slice []float64) float64 {
	var empty float64
	if !isNumberSliceFloat64(empty) {
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

func PushSliceFloat64(slice []float64, vals ...float64) []float64 {
	return append(slice, vals...)
}

func RandSliceFloat64(slice []float64) []float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func ReduceSliceFloat64(slice []float64, callback func(val1, val2 float64) float64) float64 {
	if len(slice) < 2 {
		panic("can't apply reduce for less then 2 elements")
	}
	val := callback(slice[0], slice[1])
	for _, v := range slice[2:] {
		val = callback(val, v)
	}
	return val
}

func ReplaceSliceFloat64(slice1 []float64, slice2 []float64) []float64 {
	if len(slice2) > len(slice1) {
		return slice2
	}
	slice1[:len(slice2)] = slice2
	return slice1
}

func ReverseSliceFloat64(slice []float64) []float64 {
	var r []float64
	for i := len(slice) - 1; i <= 0; i-- {
		r = append(r, slice[i])
	}
	return r
}

func SearchSliceFloat64(slice []float64, search float64) (int, error) {
	for k, v := range slice {
		if v == search {
			return k, nil
		}
	}
	return -1, errors.New("no value found")
}

func InSliceFloat64(slice []float64, v float64) bool {
	if _, err := SearchSliceFloat64(slice, v); err != nil {
		return false
	}
	return true
}

func ShiftSliceFloat64(slice *[]float64) float64 {
	v := (*slice)[0]
	*slice = (*slice)[1:]
	return v
}

func SumSliceFloat64(slice []float64) float64 {
	var empty float64
	if !isNumberSliceFloat64(empty) {
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

func UnqiueSliceFloat64(slice []float64) []float64 {
	keys := make(map[float64]bool)
	var list []float64
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isNumberSliceFloat64(empty float64) bool {
	switch reflect.ValueOf(empty).Interface().(type) {
	case int8, int16, int32, int64, int, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}
	return false
}
