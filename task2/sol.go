package main

import (
	"fmt"
	"sync"
)

func Square(numbers []int) []int {
	res := make([]int, len(numbers))
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for i, num := range numbers {
		i, num := i, num
		go func(n int) {
			defer wg.Done()
			res[i] = num * num
		}(num)
	}

	wg.Wait()
	return res
}

func main() {
	nums := []int {2, 4, 6, 8, 10}
	result := Square(nums)

	for _, val := range result {
		fmt.Print(val, " ")
	}
}