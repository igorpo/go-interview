package strings

// Given a string, determine if it a permutation of a palindrome

// We can check the even and odd counts using a table
func IsPalindromePermutationTable(s string) bool {
	charCounts := make(map[rune]int)
	var oddCount int

	r := []rune(s)
	for i := 0; i < len(r); i++ {
		charCounts[r[i]]++
	}

	for _, v := range charCounts {
		if v%2 != 0 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}

	return true
}

// We actually do not need to know the counts, only whether the counts are
// odd or even. Assuming the string is in english, or some char set knowledge of string, we can
// keep a bit vector and toggle the bits as we go, then check that at most one bit is set to 1.
// Check no bits are set: compare to 0
// Check one bit set: 00010000 - 1 = 00001111
//                    00010000 & 00001111 = 0
func IsPalindromePermutationBitVector(s string) bool {
	return false
}
