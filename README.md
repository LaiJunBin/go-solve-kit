# Go-Solve-Kit
Designed specifically for programming problem solving.

It's like VB LINQ, JavaScript array Function ... more.

## Install
```
$ go get github.com/laijunbin/go-solve-kit
```

## Import
```go
import (
	sk "github.com/laijunbin/go-solve-kit"
)
```

## Example
---

### Example 1: two number sum.
```go
func main() {
	sum := sk.LineInput().Split(" ")[:2].ToIntArray().Sum()
	fmt.Println(sum)
}
```

---
### Example 2:
```go
func main() {
	nums := []int {1, 2, 3, 4, 5}
	output := sk.FromIntArray(nums).Map(func(v sk.Int, _ int) interface{} {
		return v * v
	}).Filter(func(v sk.Type, _ int) bool {
		return v.ToInt() % 2 == 0
	}).ToStringArray().Join("\n")
	fmt.Println(output)
}
```
Output:
```
4
16
```

---
### Example 3:
```go
func main() {
	array2D := sk.NewArray(5).Map(func(_ sk.Type, i int) interface{} {
		return sk.NewRange(1, 5, 2).Map(func(v sk.Int, j int) interface{} {
			return v.ValueOf() * i
		})
	})
	fmt.Println(array2D)
	fmt.Println(array2D.Get(2, 2).ToInt())
	fmt.Println(array2D.Get(4, 1).ToInt())
	fmt.Println("-------------------")
	fmt.Println(array2D)
	fmt.Println(array2D.Length())
	fmt.Println(array2D.Flatten())
	fmt.Println(array2D.Flatten().Length())
}
```
Output:
```
[{[{0} {0} {0}]} {[{1} {3} {5}]} {[{2} {6} {10}]} {[{3} {9} {15}]} {[{4} {12} {20}]}]
10
12
-------------------
[{[{0} {0} {0}]} {[{1} {3} {5}]} {[{2} {6} {10}]} {[{3} {9} {15}]} {[{4} {12} {20}]}]
5
[{0} {0} {0} {1} {3} {5} {2} {6} {10} {3} {9} {15} {4} {12} {20}]
15
```


# Docs

## Types
Type      | Description  
----------|-----------------------
Int            | A wrapper of int.
String         | A wrapper of string.
Type           | A wrapper of any type.
IntArray       | A Int array.
StringArray    | A String array.
TypeArray      | A Type array.

---

## Common Method
Method                          | Return Type | Description  
--------------------------------|-------------|-----------------------
NewArray(size int)              | TypeArray   | New empty array.
NewRange(start, end, step int)  | IntArray    | Generate a range IntArray.
FromIntArray(items []int)       | IntArray    | From []int create IntArray.
FromStringArray(items []string) | StringArray | From []string create StringArray.
LineInput()                     | String      | Read a line of standard input.
If(boolVar bool, trueVal, falseVal interface{})                     | interface {}      | If boolVar is true, return trueVal, otherwise return falseVal.

## Basic Type

### Int
Method                 | Return Type | Description  
-----------------------|-------------|-----------------------
ValueOf()              | int         | Get int value.
ToString()             | String      | To String type.

### String
Method                         | Return Type | Description  
-------------------------------|-------------|-----------------------
ValueOf()                      | string      | Get string value.
ToInt()                        | Int         | To Int type.
Split(sep string)              | StringArray | Separated by sep.
Length()                       | Int         | Get string length.
Trim(cutset string)            | String      | A wrapper of strings.Trim()
TrimSpace()					   | String      | A wrapper of strings.TrimSpace()
TrimLeft(prefix string)        | String      | A wrapper of strings.TrimPrefix()
TrimRight(suffix string)       | String      | A wrapper of strings.TrimSuffix()
IndexOf(v string)             | Int        | A wrapper of strings.Index()
Contains(v string)             | bool        | A wrapper of strings.Contains()
ToLower()        | String      | A wrapper of strings.ToLower()
ToUpper()       | String      | A wrapper of strings.ToUpper()

### Type
Method                   | Return Type | Description  
-------------------------|-------------|-----------------------
ToString()               | String      | To String type.
ToInt()                  | Int         | To Int type.
ToArray()                | TypeArray   | To TypeArray type. (if fails, return an empty TypeArray)
ToStringArray()                | StringArray   | To StringArray type. (if fails, return an empty StringArray)
ToIntArray()                | IntArray   | To IntArray type. (if fails, return an empty IntArray)

---

## Array Type

### IntArray
Method                                             | Return Type | Description  
---------------------------------------------------|-------------|----------------
Length()                                           | Int         | Get array length.
Map(lambda func(v Int, i int) interface{})         | TypeArray   | Same as javascript array.map() or vb LINQ Select().
ForEach(lambda func(v Int, i int))         | void   | Same as javascript array.forEach().
Filter(lambda func(v Int, i int) bool)                 | IntArray    | Same as javascript array.filter() or vb LINQ FindAll().
ToStringArray()                                    | StringArray | To StringArray.
ToTypeArray()                                      | TypeArray   | To TypeArray.
Sum()                                              | Int         | Get the sum of the array.
Fill(v Int)                                        | IntArray    | Same as javascript array.fill().
Every(lambda func(v Int, i int) bool)                  | bool        | Same as javascript array.every().
Some(lambda func(v Int, i int) bool)                   | bool        | Same as javascript array.some().
Contains(v int)                                    | bool        | Same as javascript array.includes() or vb collection type Contains().
FindIndex(lambda func(v Int, i int) bool)              | Int         | Same as javascript array.findIndex().
IndexOf(v int)                                     | Int         | Same as javascript array.indexOf() or vb collection type IndexOf().
Append(v int)                                      | void        | Append the item to last.
Pop()                                              | Int         | Get last item and remove.
Enqueue(v int)                                          | void         | Append the item to first.
Dequeue()                                          | Int         | Get first item and remove.
Remove(v int)                                      | void        | Remove first match item.
First()                                            | Int         | Get first item.
Last()                                             | Int         | Get last item.
Sort()                                             | void        | Sort by increment. 
SortBy(lambda func(x, y Int))                     | void        | Sort by customize.
Copy()                     | IntArray        | Copy array.
Slice(start, end int)                     | IntArray        | Similar to javascript array.slice(), but when end = 0, it is equal to array.length


### StringArray
Method                                             | Return Type | Description  
---------------------------------------------------|-------------|----------------
Length()                                           | Int         | Get array length.
Map(lambda func(s String, i int) interface{})      | TypeArray   | Same as javascript array.map() or vb LINQ Select().
ForEach(lambda func(v String, i int))         | void   | Same as javascript array.forEach().
Filter(lambda func(s String, i int) bool)          | StringArray    | Same as javascript array.filter() or vb LINQ FindAll().
ToIntArray()                                       | IntArray    | To IntArray.
ToTypeArray()                                      | TypeArray   | To TypeArray.
Join(sep string)                                   | String      | Same as javascript array.join().
Fill(v String)                                     | StringArray    | Same as javascript array.fill().
Every(lambda func(v String, i int) bool)           | bool        | Same as javascript array.every().
Some(lambda func(v String, i int) bool)            | bool        | Same as javascript array.some().
Contains(v string)                                 | bool        | Same as javascript array.includes() or vb collection type Contains().
FindIndex(lambda func(v String, i int) bool)       | Int         | Same as javascript array.findIndex().
IndexOf(v string)                                  | Int         | Same as javascript array.indexOf() or vb collection type IndexOf().
Append(v string)                                   | void        | Append the item to last.
Pop()                                              | String      | Get last item and remove.
Enqueue(v string)                                          | void         | Append the item to first.
Dequeue()                                          | String      | Get first item and remove.
Remove(v string)                                   | void        | Remove first match item.
First()                                            | String      | Get first item.
Last()                                             | String      | Get last item.
Sort()                                             | void        | Sort by increment. 
SortBy(lambda func(x, y String))                  | void        | Sort by customize.
Copy()                     | StringArray        | Copy array.
Slice(start, end int)                     | StringArray        | Similar to javascript array.slice(), but when end = 0, it is equal to array.length

### TypeArray
Method                                             | Return Type | Description  
---------------------------------------------------|-------------|----------------
Length()                                           | Int         | Get array length.
ToStringArray()                                    | StringArray | To StringArray.
ToIntArray()                                       | IntArray    | To IntArray.
Map(lambda func(s Type, i int) interface{})        | TypeArray   | Same as javascript array.map() or vb LINQ Select().
ForEach(lambda func(v Type, i int))         | void   | Same as javascript array.forEach().
Filter(lambda func(s Type, i int) bool)            | TypeArray   | Same as javascript array.filter() or vb LINQ FindAll().
Fill(v interface{})                                | TypeArray   | Same as javascript array.fill().
Every(lambda func(v Type, i int) bool)             | bool        | Same as javascript array.every().
Some(lambda func(v Type, i int) bool)              | bool        | Same as javascript array.some().
FindIndex(lambda func(v Type, i int) bool)       | Int         | Same as javascript array.findIndex().
Sort(lambda func(x, y Type))                      | void        | Sort by customize.
Get(indexes ...int)                                | Type        | Get item by indexes.
Flatten()                                          | TypeArray   | Same as javascript array.flat().
Copy()                     | TypeArray        | Copy array.
Slice(start, end int)                     | TypeArray        | Similar to javascript array.slice(), but when end = 0, it is equal to array.length
