package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Started goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

type Job struct {
	id  int
	num int
}

type Result struct {
	job Job
	sum int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func sumDigits(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	time.Sleep(2 * time.Second)
	return num
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, sumDigits(job.num)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randonNum := rand.Intn(999)
		job := Job{i, randonNum}
		jobs <- job
	}
	close(jobs)
}

func printResult(done chan bool) {
	for result := range results {
		fmt.Printf("Job ID: %d, input num: %d, sum: %d\n", result.job.id, result.job.num, result.sum)
	}
	done <- true
}

func main() {
	// WaitGroup
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(i)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished executing.")

	// worker pool implementation
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go printResult(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken:", diff.Seconds(), "(s)")
}
