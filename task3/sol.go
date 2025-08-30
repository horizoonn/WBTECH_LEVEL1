package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func worker(id int, channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range channel {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Воркер № %d выполнил задачу %d\n", id, job)
	}
}

func main() {
	var numWorkers, numJobs int
	flag.IntVar(&numWorkers, "workers", 10, "Количество воркеров")
	flag.IntVar(&numJobs, "jobs", 30, "Количество задач")
	flag.Parse()

	wg := &sync.WaitGroup{}
	jobs := make(chan int, numJobs)

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, wg)
	}

	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i 
		}
		close(jobs)
	}()

	wg.Wait()
	fmt.Println("Все задачи выполнены")
}