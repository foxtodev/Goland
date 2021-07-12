package sort

import (
	"sort"
)

func SortBubble(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func SortInsertion(data []int) []int {
	for i := range data {
		for j := i; j > 0 && data[j-1] > data[j]; j-- {
			data[j-1], data[j] = data[j], data[j-1]
		}
	}
	return data
}

func SortQuick(data []int) []int {
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1
	pivot := uint(len(data) / 2)
	data[pivot], data[right] = data[right], data[pivot]

	for i := range data {
		if data[i] < data[right] {
			data[i], data[left] = data[left], data[i]
			left++
		}
	}

	data[left], data[right] = data[right], data[left]

	SortQuick(data[:left])
	SortQuick(data[left+1:])

	return data
}

func SortGoSlice(data []int) []int {
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data
}
