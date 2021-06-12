###Insert sorting algorithm

```go
func SortInsertion(data []int) []int {
	for i := range data {
		for j := i; j > 0 && data[j-1] > data[j]; j-- {
			data[j-1], data[j] = data[j], data[j-1]
		}
	}
	return data
}
```
uses
```go
data := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
fmt.Println(SortInsertion(data))
// [0 1 2 3 4 5 6 7 8 9]
```

###Insert sorting algorithm with direction

```go
func SortInsertionDir(data []int, direction ...int) []int {
	dir := 0
	if len(direction) == 0 {
		dir = 1
	}
	for i := range data {
		for j := i; j > 0 && data[j-dir] > data[j+(dir-1)]; j-- {
			data[j-1], data[j] = data[j], data[j-1]
		}
	}
	return data
}
```
uses
```go
data1 := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
fmt.Println(SortInsertionDir(data1))
// [0 1 2 3 4 5 6 7 8 9]

data2 := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
fmt.Println(SortInsertionDir(data2, 1))
// [9 8 7 6 5 4 3 2 1 0]
```