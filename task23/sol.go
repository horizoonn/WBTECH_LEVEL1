package main

import "fmt"

func DeleteElem[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return append([]T{}, slice...)
	}

	result := make([]T, 0, len(slice) - 1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index + 1:]...)
	return result
}

func main() {
	arr := []int{1, 2, 3, 4}
	res := DeleteElem(arr, 1)
	fmt.Println(res)
}