package main

import (
	"fmt"
	"time"
)

//SAMPLE 1

// // jobId, receiver, sender
// func worker(id int, tasks <-chan int, result chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		//simulate work
// 		time.Sleep(time.Second)
// 		result <- task * 2
// 	}
// }

// func main() {

// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	//create workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	//send task to channel
// 	for i := range numJobs {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	//collect the results
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println(result)
// 	}
// }

// SAMPLE 2

type ticketRequest struct {
	personID     int
	numberTicket int
	cost         int
}

// simulate
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket[s] of PersonID %d with total cost %d\n", req.numberTicket, req.personID, req.cost)
		//simulate Processing
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func mainwp() {

	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	//start ticket processor/worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	//send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numberTicket: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for pesonID %d proccessed successfully\n", <-ticketResults)
	}
}
