package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	done := make(chan struct{})

	go func() {
		<-time.After(duration)
		close(done)
	}()

	<-done
}

func main() {
	start := time.Now()
	fmt.Println("Начало:", start)
	sleep(3 * time.Second)
	fmt.Println("Конец:", time.Now())
	fmt.Println("Прошло времени:", time.Since(start))
}