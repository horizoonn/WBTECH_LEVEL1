// Данный фрагмент кода может привести к утечке памяти.
// Подстрока разделяет память с исходной строкой, 
// поэтому маленькая justString не даёт освободить большой массив до тех пор, 
// пока сама живёт, что ведёт к скрытому росту памяти.

//justString = strings.Clone(v[:100]) — создаётся новый буфер только на 100 байт, память под v больше не удерживается и может быть освобождена.


package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	return strings.Repeat("a", n)
}

func someFunc() {
	v := createHugeString(1 << 20)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
	fmt.Println(len(justString))
}
