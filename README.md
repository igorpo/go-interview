# Interview Prep in Golang

### Notes for Go

1. Length of string can be found with `len(s)`
2. Accessing char at index `i` of string `s`: `string(s[i])`
  * `s[i]` on its own will be a `byte` type.
  * `s[:i]` does not include index `i` of `s`
  * `s[i:]` include index `i` of `s`
  * `s[:i]+s[i:] = s`
3. `r := []rune(s)` will give the `char`-like codepoint representation of a string `s`
4. Fast way to sort a string `s`. A better way could be to implement Len, Less, Swap interface
   for a custom class of []rune. Perhaps there is an even better way...
   ```golang
   sArr := strings.Split(s, "")
   sort.Strings(sArr)
   sorted := strings.Join(sArr, "")
   ```
5. When defining tree nodes, use pointers to children
6. Append lists with `list = append(list, other_list...)`
7. Generally, Go passes by value, except with maps and lists, pass a pointer to the list instead
 * `*list = append(*list, value)`
8. If you want to key a struct in Go within a map, use a pointer to the struct as the key.