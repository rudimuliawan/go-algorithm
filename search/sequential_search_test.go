package search

import (
	"go-algorithm/internal/assert"
	"testing"
)

func TestSequentialSearchST(t *testing.T) {
	st := new(SequentialSearchST[string, int])
	st.Put("one", 1)
	st.Put("two", 2)
	st.Put("three", 3)

	tests := []struct {
		key  string
		want int
		ok   bool
	}{
		{
			key:  "one",
			want: 1,
			ok:   true,
		},
		{
			key:  "two",
			want: 2,
			ok:   true,
		},
		{
			key:  "three",
			want: 3,
			ok:   true,
		},
		{
			key:  "four",
			want: 0,
			ok:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			val, ok := st.Get(tt.key)
			assert.Equal(t, ok, tt.ok)
			assert.Equal(t, val, tt.want)
		})
	}

	_ = st.Delete("one")
	_ = st.Delete("three")

	st.Put("four", 4)

	tests = []struct {
		key  string
		want int
		ok   bool
	}{
		{
			key:  "one",
			want: 0,
			ok:   false,
		},
		{
			key:  "two",
			want: 2,
			ok:   true,
		},
		{
			key:  "three",
			want: 0,
			ok:   false,
		},
		{
			key:  "four",
			want: 4,
			ok:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			val, ok := st.Get(tt.key)
			assert.Equal(t, ok, tt.ok)
			assert.Equal(t, val, tt.want)
		})
	}
}
