package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty data",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Already sorted ascending",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Already sorted descending",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Test for random input",
			input:    []int{5, 3, 8, 7, 6, 10, 2, 1, 4, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			SelectionSort(arr)

			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("InsertionSort() for input %v got %v, want %v", tt.input, arr, tt.expected)
			}
		})
	}
}
