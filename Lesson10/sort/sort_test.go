package sort

import (
	"fmt"
	"testing"
)

func ArrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var wanted = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestSortBubble(t *testing.T) {
	given := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
	received := SortBubble(given)
	if !ArrayEqual(received, wanted) {
		t.Errorf("Received: \t%d, Wanted:\t%d", received, wanted)
	}
}

func ExampleSortBubble() {
	given := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
	fmt.Println(SortBubble(given))
	// Output: [0 1 2 3 4 5 6 7 8 9]
}

func TestSortInsertion(t *testing.T) {
	given := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
	received := SortBubble(given)
	if !ArrayEqual(received, wanted) {
		t.Errorf("Received: \t%d, Wanted:\t%d", received, wanted)
	}
}

func TestSortQuick(t *testing.T) {
	given := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
	received := SortBubble(given)
	if !ArrayEqual(received, wanted) {
		t.Errorf("Received: \t%d, Wanted:\t%d", received, wanted)
	}
}

func TestSortGoSlice(t *testing.T) {
	given := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
	received := SortBubble(given)
	if !ArrayEqual(received, wanted) {
		t.Errorf("Received: \t%d, Wanted:\t%d", received, wanted)
	}
}
