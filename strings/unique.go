package strings

import (
	"sort"
	"strings"
	"unicode/utf8"
)

// Implement an algorithm to determine if a given string has all unique characters

// Neat trick if we are using ASCII chars, can store in array (len 256) instead of
// using a hashmap or sorting.
// Note: we assume case sensitivity
func UniqueASCII(s string) bool {
	// we know 128 is the number of ASCII chars
	var ascii [128]int

	runeStr := []rune(s)
	for i := 0; i < len(runeStr); i++ {
		v := ascii[runeStr[i]]
		if v == 0 {
			ascii[runeStr[i]]++
			continue
		}
		return false
	}

	return true
}

// Another way is to sort the string and check for chars next to each other being identical
// Note we assume case sensitivity here.
func Unique(s string) bool {

	// quick way to sort a string, better would be to implemenat the Less, Len, Swap
	// funcs of a custom type like []rune
	sArr := strings.Split(s, "")
	sort.Strings(sArr)
	sorted := strings.Join(sArr, "")

	// this would work if non utf8 chars were allowed 
	//for i := 0; i < len(sorted)-1; i++ {
	//	if sorted[i] == sorted[i+1] {
	//		return false
	//	}
	//}
	//return true
	return compareCharsUTF8(sorted)
}

func compareCharsUTF8(s string) bool {
	bytes := []byte(s)
	for utf8.RuneCount(bytes) > 1 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		next, size := utf8.DecodeRune(bytes)
		if r == next {
			return false
		}
	}
	return true
}
