package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReverseString(s string) string {
	t := []rune(s)

	for i, j := 0, len(t) - 1; i < j; i, j = i + 1, j - 1 {
		t[i], t[j] = t[j], t[i]
	}

	return string(t)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println(ReverseString(scanner.Text()))
	}
}