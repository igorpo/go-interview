package strings

// Given two strings s and t, determine if t can be made from either replacing a char,
// inserting a char, or removing a char from s.

// We only really need to consider the length of the target string
func IsOneEditAway(s, t string) bool {
	sL, tL := len(s), len(t)
	if sL-tL == 0 {
		return checkReplace(s, t)
	} else if sL-tL == 1 {
		return checkInsertOrRemoval(t, s)
	} else if sL-tL == -1 {
		return checkInsertOrRemoval(s, t)
	}

	return false
}

func checkReplace(s, t string) bool {
	pS, pT, countDiffs := 0, 0, 0
	for pT < len(t) {
		if string(s[pS]) != string(t[pT]) {
			countDiffs++
			if countDiffs > 1 {
				return false
			}
		}
		pS++
		pT++
	}
	return true
}

func checkInsertOrRemoval(s, t string) bool {
	pS, pT, countDiffs := 0, 0, 0
	for pS < len(s) {
		if string(s[pS]) != string(t[pT]) {
			countDiffs++
			if countDiffs > 1 {
				return false
			}
			pS--
		}
		pS++
		pT++
	}

	return true
}
