package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing Job %d\n", id, job)
	}
}

func main() {

	jobs := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	go worker(1, jobs, &wg)
	go worker(2, jobs, &wg)
	go worker(3, jobs, &wg)

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
}
