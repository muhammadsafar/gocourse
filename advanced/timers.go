package main

import (
	"fmt"
	"time"
)

// func main() {
// 	timer1 := time.NewTimer(1 * time.Second)
// 	timer2 := time.NewTimer(2 * time.Second)

// 	select {
// 	case <-timer1.C:
// 		fmt.Println("timer1 expired")
// 	case <-timer2.C:
// 		fmt.Println("timer2 expired")
// 	}
// }

// 3. SCHEDULING DELAYED OPERATIONS

// func main() {
// 	timer := time.NewTimer(2 * time.Second)
// 	go func() {
// 		// timer.Reset(4 * time.Second)
// 		<-timer.C
// 		fmt.Println("Delayed operation executed...")
// 	}()
// 	fmt.Println("waiting...")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("End of proram")
// }

// 2. =====TIMEOUT
func longRunningOperation() {
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func maintimer() {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	go func() {
		longRunningOperation()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("operation timeout")
	case <-done:
		fmt.Println("operation complete")
	}
}

//1.

// func main() {
// 	// time.Sleep(time.Second)
// 	fmt.Println("Starting app...")
// 	timer := time.NewTimer(5 * time.Second)
// 	fmt.Println("waiting for timer.c...")
// 	stopped := timer.Stop() //stopped berarti lama timer diabaikan dan langsung stop
// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}
// 	fmt.Println("timer reset...")
// 	timer.Reset(3 * time.Second) //atur ulang newtimer jadi 3s
// 	<-timer.C                    //start blocking here selama waktu timer
// 	fmt.Println("Timer expired...")
// }
