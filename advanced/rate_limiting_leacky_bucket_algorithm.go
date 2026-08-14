package main

import (
	"fmt"
	"sync"
	"time"
)

type LeackyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	mux      sync.Mutex
}

func NewLeackyBucket(capacity int, leakRate time.Duration) *LeackyBucket {
	return &LeackyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeackyBucket) Allow() bool {
	lb.mux.Lock()
	defer lb.mux.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens = lb.tokens + tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
	// lb.lastLeak = lb.lastLeak.Add(elapsedTime) //sama seperti di atas

	fmt.Printf("Tokens addeed %d, token subtracted %d, total tokens : %d \n", tokensToAdd, 1, lb.tokens)
	fmt.Printf("Last leak time : %v\n", lb.lastLeak)
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false
}

func mainleackybucket() {
	LeackyBucketInstance := NewLeackyBucket(5, 500*time.Millisecond) //max token 5,  refill tiap 500 ms
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if LeackyBucketInstance.Allow() {
				fmt.Println("Current time : ", time.Now())
				fmt.Println("Request Allowed")
			} else {
				fmt.Println("Current time : ", time.Now())
				fmt.Println("XXX---Request Denied")
			}
			time.Sleep(200 * time.Millisecond) //req each 200ms
		}()
	}
	wg.Wait()

}
