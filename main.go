package go_solve_kit

import (
	"bufio"
	"os"
	"strings"
)


func NewArray(size int) TypeArray {
	return make(TypeArray, size)
}

func NewRange(start, end, step int) IntArray {
	var output IntArray
	for i := start; i <= end; i += step {
		output = append(output, Int(i))
	}
	return output
}

func FromIntArray(items []int) IntArray {
	var output IntArray
	for _, v := range items {
		output = append(output, Int(v))
	}
	return output
}

func FromStringArray(items []string) StringArray {
	var output StringArray
	for _, v := range items {
		output = append(output, String(v))
	}
	return output
}

var reader = bufio.NewReader(os.Stdin)
func LineInput() String {
	line , _ := reader.ReadString('\n')
	return String(strings.TrimSpace(line))
}