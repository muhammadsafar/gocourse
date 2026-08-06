package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// 3.
func doWork2(ctx context.Context) {
	//infinite
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled: ", ctx.Err())
			return
		default:
			fmt.Println("working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	rootCtx := context.Background()
	// ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	// defer cancel()

	//manual
	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		time.Sleep(2 * time.Second) //mirip withtimeout diatas
		cancel()
	}()

	ctx = context.WithValue(ctx, "name", "Safar Baharuddin")
	ctx = context.WithValue(ctx, "habibi", "my habibi")

	go doWork2(ctx)
	time.Sleep(3 * time.Second)

	name := ctx.Value("name")
	if name != nil {
		fmt.Println("Name >> ", name)
	} else {
		fmt.Println("No value..")
	}

	logWithContext(ctx, "Test message")
}

func logWithContext(ctx context.Context, msg string) {
	name := ctx.Value("name")
	log.Printf("Name : %v - %v", name, msg)
}

//2.

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()
// 	result := checkEvenOdd(ctx, 31)
// 	fmt.Println("result TODO >>", result)

// 	//misal ditimeout
// 	ctx, cancelTmOut := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancelTmOut() //cencel kemudian
// 	result2 := checkEvenOdd(ctx, 31)
// 	fmt.Println("result before TO >>", result2)
// 	time.Sleep(2 * time.Second)
// 	result2 = checkEvenOdd(ctx, 31)
// 	fmt.Println("result after TO >>", result2)

// 	//misal dicancel
// 	ctx, cancel := context.WithCancel(context.Background())
// 	cancel()
// 	result3 := checkEvenOdd(ctx, 31)
// 	fmt.Println("result Cancel >>", result3)

// }

// 1.

//=======Difference context.TODO and context.backgournd
// func main() {

// 	todoCtx := context.TODO()
// 	contextBg := context.Background()

// 	ctx := context.WithValue(todoCtx, "name", "syakir")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx2 := context.WithValue(contextBg, "city", "ppu")
// 	fmt.Println(ctx2)
// 	fmt.Println(ctx2.Value("city"))

// }
