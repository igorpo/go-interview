package strings

import "strings"

// Given a string s, implement and algorithm that will replace all spaces with %20
// Note: uses a builder, not done in place
func UrlifyBuilder(s string) string {
	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			sb.WriteString("%20")
		} else {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}
