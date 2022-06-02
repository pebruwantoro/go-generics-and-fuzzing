package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func ReverseString(s string) (string, error) {
	// b := []byte(s)
	// for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
	// 	b[i], b[j] = b[j], b[i]
	// }
	// return string(b)
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	fmt.Printf("input: %q\n", s)
	r := []rune(s)
	fmt.Printf("runes: %q\n", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func Divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return a / b, a % b, nil
}
