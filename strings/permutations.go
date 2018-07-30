package strings

// Given a string s, design and implement an algorithm to list its permutations
// i.e. "ab" -> ["ab", "ba"]
func Permutations(s string) []string {
	return permutationHelper([]string{}, s)
}

func permutationHelper(list []string, s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	lastChar := string(s[len(s)-1])
	prevPerms := permutationHelper(list, s[:len(s)-1])
	currPerms := []string{}

	for _, str := range prevPerms {
		for i := 0; i <= len(str); i++ {
			currPerms = append(currPerms, str[:i]+lastChar+str[i:])
		}
	}

	return currPerms
}
