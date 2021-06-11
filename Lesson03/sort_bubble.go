package main

import "fmt"

func sort_bubble(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func main() {
	data := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
	fmt.Println(sort_bubble(data))
}
