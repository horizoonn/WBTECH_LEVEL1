package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words) - 1; i < j; i, j = i + 1, j - 1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Println(ReverseWords(input))
	}
}