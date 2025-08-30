package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	var n int
	fmt.Print("Введите количество секунд работы программы: ")
	if _, err := fmt.Scanf("%d", &n); err != nil || n <= 0 {
		fmt.Println("Ошибка: введите положительное число")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n) * time.Second)
	defer cancel()

	go func() {
		defer close(channel)
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				select {
				case channel <- i:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	for value := range channel {
		fmt.Println(value)
	}
}