package main

import (
	"fmt"
	"time"
)

//TICKER RESET

func mainticker() {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for tick := range ticker.C {
			fmt.Println("Tick at :", tick)
		}
	}()

	//beri space waktu 3 s untuk biarkan new ticker jalan tiap 1 s
	time.Sleep(3 * time.Second)

	fmt.Println("Reset ticker menjadi 3 detik")
	ticker.Reset(3 * time.Second)

	//beri space waktu 10s untuk biarkan new reset ticker jalan tiap 3s
	time.Sleep(10 * time.Second)

	//output
	// Tick at : 2026-08-07 11:23:10.4771248 +0800 +08 m=+1.000508601
	// Tick at : 2026-08-07 11:23:11.4771248 +0800 +08 m=+2.000508601
	// Tick at : 2026-08-07 11:23:12.4771248 +0800 +08 m=+3.000508601
	// Reset ticker menjadi 3 detik
	// Tick at : 2026-08-07 11:23:15.4777682 +0800 +08 m=+6.001152001
	// Tick at : 2026-08-07 11:23:18.4777682 +0800 +08 m=+9.001152001
	// Tick at : 2026-08-07 11:23:21.4777682 +0800 +08 m=+12.001152001
}

//=====TICKER STOP
// func main() {
// 	ticker := time.NewTicker(time.Second) //interval 1s
// 	stop := time.After(5 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case tick := <-ticker.C:
// 			fmt.Println("Tick at ", tick)
// 		case <-stop:
// 			fmt.Println("Stopping ticker.")
// 			return
// 		}
// 	}
// }

//SCHEDULING SAMPLE ITERATION

// func periodicTask() {
// 	fmt.Println("Performing periodic task at :", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()

// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				periodicTask()
// 			}
// 		}
// 	}()

// 	<-time.After(5 * time.Second)

// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at : ", tick)
// 	// }

// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// 	for tick := range ticker.C {
// 		i++
// 		fmt.Println("Tick at : ", tick)
// 	}

// }
