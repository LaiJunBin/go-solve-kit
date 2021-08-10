package go_solve_kit

import (
	"sort"
	"strconv"
)


type Int int
type IntArray []Int

func (i Int) ValueOf() int {
	return int(i)
}

func (i Int) ToString() String {
	return String(strconv.Itoa(i.ValueOf()))
}

func (array IntArray) Length() Int {
	return Int(len(array))
}

func (array IntArray) Map(lambda func(v, i int) interface{}) TypeArray {
	var output TypeArray
	for i, v := range array {
		output = append(output, Type{lambda(v.ValueOf(), i)})
	}
	return output
}


func (array IntArray) Filter(lambda func(s, i int) bool) IntArray {
	var output IntArray
	for i, v := range array {
		if lambda(v.ValueOf(), i) {
			output = append(output, v)
		}
	}
	return output
}

func (array IntArray) ToTypeArray() TypeArray {
	var output TypeArray
	for _, v := range array {
		output = append(output, Type{v.ValueOf()})
	}
	return output
}

func (array IntArray) ToStringArray() StringArray {
	var output StringArray
	for _, v := range array {
		output = append(output, v.ToString())
	}
	return output
}

func (array IntArray) Sum() Int {
	var sum Int
	for _, v := range array {
		sum += v
	}
	return sum
}

func (array IntArray) Fill(v Int) IntArray {
	return NewArray(len(array)).Map(func(_ Type, _ int) interface{} {
		return v
	}).ToIntArray()
}

func (array IntArray) Every(lambda func(v, i int) bool) bool {
	for i, val := range array {
		if !lambda(val.ValueOf(), i) {
			return false
		}
	}
	return true
}

func (array IntArray) Some(lambda func(v, i int) bool) bool {
	for i, val := range array {
		if lambda(val.ValueOf(), i) {
			return true
		}
	}
	return false
}

func (array IntArray) Contains(v int) bool {
	return array.Some(func(i, _ int) bool {
		return i == v
	})
}

func (array IntArray) FindIndex(lambda func(v, i int) bool) Int {
	for i, val := range array {
		if lambda(val.ValueOf(), i) {
			return Int(i)
		}
	}
	return -1
}


func (array IntArray) IndexOf(v int) Int {
	return array.FindIndex(func(i, _ int) bool {
		return i == v
	})
}

func (array *IntArray) Append(v int) {
	*array = append(*array, Int(v))
}

func (array *IntArray) Pop() Int {
	output := (*array)[0]
	*array = (*array)[1:]
	return output
}

func (array *IntArray) Dequeue() Int {
	output := (*array)[array.Length()-1]
	*array = (*array)[:array.Length()-1]
	return output
}

func (array IntArray) First() Int {
	return array[0]
}

func (array IntArray) Last() Int {
	return array[array.Length()-1]
}

func (array *IntArray) Remove(v int) {
	i := array.IndexOf(v)
	*array = append((*array)[:i], (*array)[i+1:]...)
}

func (array IntArray) Sort() {
	sort.SliceStable(array, func(i, j int) bool {
		return array[i] < array[j]
	})
}

func (array IntArray) SortBy(lambda func(x, y Int) bool) {
	sort.SliceStable(array, func(i, j int) bool {
		return lambda(array[i], array[j])
	})
}