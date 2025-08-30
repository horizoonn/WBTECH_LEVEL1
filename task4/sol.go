package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер № %d получил сигнал завершения\n", id)
			return
		case job, ok :=<- channel:
			if !ok {
				return
			}
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Воркер № %d выполнил задачу %d\n", id, job)
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var numWorkers, numJobs int
	flag.IntVar(&numWorkers, "workers", 10, "Количество воркеров")
	flag.IntVar(&numJobs, "jobs", 1000, "Количество задач")
	flag.Parse()

	wg := &sync.WaitGroup{}
	jobs := make(chan int, numWorkers)

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, jobs, wg)
	}

	go func() {
		defer close(jobs)
		for i := 1; i <= numJobs; i++ {
			select {
			case jobs <- i:
			case <-ctx.Done():
				fmt.Println("Прервали отправление задач")
				return
			} 
		}
	}()

	wg.Wait()

	select {
	case <-ctx.Done():
		fmt.Println("Остановлено по сигналу")
	default:
		fmt.Println("Все задачи выполнены")
	}
}