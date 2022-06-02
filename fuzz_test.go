package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestReverseString(t *testing.T) {
// 	testcases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}
// 	for _, tc := range testcases {
// 		rev, _ := ReverseString(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse: %q, want %q", rev, tc.want)
// 		}
// 	}
// }

// func FuzzReverseString(f *testing.F) {
// 	testcases := []string{"Hello, world", " ", "!12345"}
// 	for _, tc := range testcases {
// 		f.Add(tc) // Use f.Add to provide a seed corpus
// 	}
// 	f.Fuzz(func(t *testing.T, orig string) {
// 		rev, err1 := ReverseString(orig)
// 		if err1 != nil {
// 			return
// 		}
// 		doubleRev, err2 := ReverseString(rev)
// 		if err2 != nil {
// 			return
// 		}
// 		if orig != doubleRev {
// 			t.Errorf("Before: %q, after: %q", orig, doubleRev)
// 		}
// 		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
// 			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
// 		}
// 	})
// }

func FuzzDivide(f *testing.F) {
	up := []int{1, 2, 3, 4, 0, 1000, 10101}
	bottom := []int{0, 8, 10, 2048, 1000}
	for _, x := range up {
		for _, y := range bottom {
			f.Add(x, y) // Use f.Add to provide a seed corpus
		}
	}
	f.Fuzz(func(t *testing.T, a, b int) {
		result, r, err := Divide(a, b)
		if err != nil {
			assert.Error(t, err)
		}
		assert.NotNil(t, result)
		assert.NotNil(t, r)
	})
}
