package main

import "fmt"

func QuickSort(arr []int, left int, right int) []int {
	if left < right {
	var p int
	arr, p = Partition(arr, left, right)
	arr = QuickSort(arr, left, p-1)
	arr = QuickSort(arr, p+1, right)
	}
	return arr
}

func QuickSortStart(arr []int) []int {
	return QuickSort(arr, 0, len(arr)-1)
}

func Partition(arr []int, left, right int) ([]int, int) {
	p := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}

func main() {
	arr := []int{3, 43, 11, 943, 13, 4353, 232, 4303, 53532, 23142}

	res := QuickSortStart(arr)
	fmt.Println(res)
}