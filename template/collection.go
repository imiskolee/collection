package template

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/pkg/errors"
)

// template type Slice(T)
type T int
type Slice []T
type IsNumber string
func Chunk(slice []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func Diff(slice1 []T, slice2 []T) []T {
	var diff []T
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

func Intersect(slice1 []T, slice2 []T) []T {
	var inter []T
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

func Fill(value T, size int) []T {
	r := make([]T,size)
	for i :=0; i<size;i++ {
		r = append(r,value)
	}
	return r
}

func Filter(slice []T, callback func(val T) bool) []T {
	var filtered []T
	for _,v := range slice {
		if callback(v) {
			filtered = append(filtered,v)
		}
	}
	return filtered
}

func Pad(slice []T, size int, value T) []T {
	diff := size - len(slice)
	for i :=0;i<diff;i++ {
		slice = append(slice,value)
	}
	return slice
}

func Product(slice []T) T {
	var empty T
	if !isNumber(empty) {
		panic("Can't apply function `product` to non number type.")
	}
	switch reflect.ValueOf(empty).Interface().(type) {
	 case int8,int16,int32,int64,int,uint,uint8,uint16,uint32,uint64,float32,float64:
	 	for k,v := range slice {
			if k == 0 {
				empty = v
				continue
			}
			empty = empty * v
		}
	}
	return empty
}

func Push(slice []T, vals ...T) []T {
	return append(slice,vals...)
}

func Rand(slice []T) []T {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func Reduce(slice []T, callback func(val1, val2 T) T) T {
	if len(slice) < 2 {
		panic("can't apply reduce for less then 2 elements")
	}
	val := callback(slice[0],slice[1])
	for _,v := range slice[2:] {
		val = callback(val,v)
	}
	return val
}

func Replace(slice1 []T, slice2 []T) []T {
	if len(slice2) > len(slice1) {
		return slice2
	}
	slice1[:len(slice2)] = slice2
	return slice1
}

func Reverse(slice []T) []T {
	var r []T
	for i := len(slice) -1;i <=0;i-- {
		r = append(r,slice[i])
	}
	return r
}

func Search(slice []T, search T) (int, error) {
	for k,v := range slice {
		if v == search {
			return k,nil
		}
	}
	return -1,errors.New("no value found")
}

func In(slice []T, v T) bool {
	if _,err := Search(slice,v); err != nil {
		return false
	}
	return true
}

func Shift(slice *[]T) T {
	v := (*slice)[0]
	*slice = (*slice)[1:]
	return v
}

func Sum(slice []T) T {
	var empty T
	if !isNumber(empty) {
		panic("Can't apply function `product` to non number type.")
	}
	switch reflect.ValueOf(empty).Interface().(type) {
	case int8,int16,int32,int64,int,uint,uint8,uint16,uint32,uint64,float32,float64:
		for k,v := range slice {
			if k == 0 {
				empty = v
				continue
			}
			empty = empty + v
		}
	}
	return empty
}

func Unqiue(slice []T) []T {
	keys := make(map[T]bool)
	var list []T
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isNumber(empty T) bool {
	switch reflect.ValueOf(empty).Interface().(type) {
	case int8,int16,int32,int64,int,uint,uint8,uint16,uint32,uint64,float32,float64:
		return true
	}
	return false
}
