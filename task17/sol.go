package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	var m int
	for l <= r {
		m = int(l + (r - l) / 2)
		if arr[m] > target {
			r = m - 1
		} else if arr[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 43, 65, 231, 5423, 534223}

	 if idx := BinarySearch(nums, 65); idx != -1 {
		fmt.Println("значение найдено по индексу", idx)
	 } else {
		fmt.Println("Значение не найдено")
	 }
}