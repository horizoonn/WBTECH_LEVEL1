package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func exitByCondition() {
	fmt.Println("1) Выход по условию")
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Горутина 1, шаг", i)
			time.Sleep(100 * time.Millisecond)
		} 
		fmt.Println("Горутина 1 завершилась по условию")
	}()
	time.Sleep(600 * time.Millisecond)
}

func exitByChannel() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Горутина 2 начала работу")
		for {
			select {
			case <-done:
				fmt.Println("Горутина 2 завершилась через канал уведомлений")
				return
			default:
				fmt.Println("Горутина 2 продолжает работу")
				time.Sleep(150 * time.Millisecond)
			}
		}
	}()

	time.Sleep(600 * time.Millisecond)
	close(done)
}

func exitByContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина 3 завершилась через контекст")
				return
			default:
				fmt.Println("Горутина 3 продолжает работу")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(250 * time.Millisecond)
	cancel()
	time.Sleep(50 * time.Millisecond)
}

func exitByRuntime() {
	go func() {
		defer fmt.Println("Горутина 4 завершила работу при помощи runtime.Goexit()")
		
		fmt.Println("Горутина 4 начала работу")
		time.Sleep(250 * time.Millisecond)
		fmt.Println("Горутина 4: вызываю runtime.Goexit()")
		runtime.Goexit()
	}()
	time.Sleep(300 * time.Millisecond)
}

func exitByTimer() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		fmt.Println("Горутина 5 начала работу")
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Горутина 5 завершилась по таймеру")
		}		
	}()
	<-done
}

func exitByClosedChannel() {
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(done)
		fmt.Println("Горутина 6 начала работу")
		for value := range ch {
			fmt.Println("Горутина 6 получила значение", value)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	go func() {
		defer close(ch)
		for i := 0; i <= 50; i+=10 {
			ch <- i
		}
	}()

	<-done
	fmt.Println("Горутина 6 завершилась")
}

func main() {
	exitByCondition()
	exitByChannel()
	exitByContext()
	exitByRuntime()
	exitByTimer()
	exitByClosedChannel()
}