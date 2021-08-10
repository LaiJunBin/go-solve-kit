package go_solve_kit

import (
	"sort"
	"strconv"
)


type Type struct {
	variable interface{}
}
type TypeArray []Type


func (t Type) ToString() String {
	if t.variable == nil {
		return ""
	}

	if val, ok := t.variable.(String); ok {
		return val
	}

	if i, ok := t.variable.(int); ok {
		return String(strconv.Itoa(i))
	}

	return String(t.variable.(string))
}

func (t Type) ToInt() Int {
	if t.variable == nil {
		return 0
	}

	if val, ok := t.variable.(Int); ok {
		return val
	}

	val, _ := strconv.Atoi(t.ToString().ValueOf())
	return Int(val)
}

func (t Type) ToArray() TypeArray {
	if val, ok := t.variable.(TypeArray); ok {
		return val
	}

	return TypeArray{}
}

func (array TypeArray) Length() Int {
	return Int(len(array))
}

func (array TypeArray) Map(lambda func(v Type, i int) interface{}) TypeArray {
	var output TypeArray
	for i, v := range array {
		output = append(output, Type{lambda(v, i)})
	}
	return output
}

func (array TypeArray) Filter(lambda func(s Type, i int) bool) TypeArray {
	var output TypeArray
	for i, v := range array {
		if lambda(v, i) {
			output = append(output, v)
		}
	}
	return output
}

func (array TypeArray) ToStringArray() StringArray {
	var output StringArray
	for _, v := range array {
		output = append(output, v.ToString())
	}
	return output
}

func (array TypeArray) ToIntArray() IntArray {
	var output IntArray
	for _, v := range array {
		output = append(output, v.ToInt())
	}
	return output
}

func (array TypeArray) Fill(v interface{}) TypeArray {
	return NewArray(len(array)).Map(func(_ Type, _ int) interface{} {
		return v
	})
}

func (array TypeArray) Every(lambda func(v Type, i int) bool) bool {
	for i, val := range array {
		if !lambda(val, i) {
			return false
		}
	}
	return true
}

func (array TypeArray) Some(lambda func(v Type, i int) bool) bool {
	for i, val := range array {
		if lambda(val, i) {
			return true
		}
	}
	return false
}

func (array TypeArray) Sort(lambda func(x, y Type) bool) {
	sort.SliceStable(array, func(i, j int) bool {
		return lambda(array[i], array[j])
	})
}

func (array TypeArray) Get(indexes ...int) Type {
	if len(indexes) == 1 {
		return array[indexes[0]]
	}

	return array[indexes[0]].ToArray().Get(indexes[1:]...)
}

func (array TypeArray) Flatten() TypeArray {
	var output TypeArray
	for _, v := range array {
		output = append(output, v.ToArray()...)
	}
	return output
}