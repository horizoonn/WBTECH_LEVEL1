package main

import "fmt"

func generator(arr []int) <-chan int {
	output := make(chan int)
	result := make(chan int)

	go func() {
		defer close(output)
		for _, val := range arr {
			output <- val
		}
	}()

	go func() {
		defer close(result)
		for val := range output {
			result <- val * 2
		}
	}()

	return result
}

func main() {
	nums := []int{2, 3, 54, 32, 54, 52, 82, 924, 8493, 2324, 5453, 2423, 4322, 3524}

	channel := generator(nums)
	for val := range channel {
		fmt.Println(val)
	}
}