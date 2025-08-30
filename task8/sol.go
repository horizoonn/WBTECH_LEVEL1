package main

import "fmt"

func SetBit(num int64, pos uint, value int) int64 {
	if pos >= 64 {
		return num
	}
	mask := int64(1) << pos
	switch value{
	case 1:
		return num | mask
	case 0:
		return num &^ mask
	default:
		return num
	}
}

func main() {
	var num int64 = 5

	fmt.Printf("Исходное число: %d, в двоичном виде: %b\n", num, num)
	result := SetBit(num, 0, 0)
	fmt.Printf("После установки бита c индексом 0 в 0: %d, в двоичном виде: %b", result, result)
}