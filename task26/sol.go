package main

import (
	"fmt"
	"unicode"
)

func IsUnique(s string) bool {
	set := make(map[rune]struct{})
	for _, r := range s {
		lr := unicode.ToLower(r)
		if _, ok := set[lr]; ok {
			return false
		}
		set[lr] = struct{}{}
	}
	return true
}

func main() {
	str := "abcdA"
	fmt.Println(IsUnique(str))
}