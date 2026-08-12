package main

import (
	"fmt"
	"sync"
	"time"
)

// ================ constuctor example============
type Worker struct {
	ID   int
	Task string
}

func (w Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worder ID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("Worker ID %d finished %s\n", w.ID, w.Task)
}

func mainwg() {

	var wg sync.WaitGroup

	//define tasks to be performed by workers
	tasks := []string{"painting", "planting", "prunning"}

	for i, task := range tasks {
		worker := Worker{
			ID:   i + 1,
			Task: task,
		}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	//wait for all workers to finish
	wg.Wait()

	//
	fmt.Println("Constuction finished...")

}

//===============EXAMPLE WITH CHANNEL==============

// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {

// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d is starting\n", id)
// 	time.Sleep(1 * time.Second)
// 	for task := range tasks {
// 		results <- task * 2
// 	}
// 	fmt.Printf("WorkerID %d is finished...\n", id)
// }

// func main() {

// 	var wg sync.WaitGroup

// 	numWorkers := 3
// 	numJobs := 3
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i + 1
// 	}
// 	close(tasks)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for res := range results {
// 		fmt.Println("Result >>", res)
// 	}

// }

// ===============BASIC EXAMPLE WO USING CHANNEL===========
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// wg.Add(1) //wrong practice
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished\n", id)
// }
// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers) //here the correct way to add counter
// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("All workers finished")
// }
