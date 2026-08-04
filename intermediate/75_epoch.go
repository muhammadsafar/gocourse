package main

import (
	"fmt"
	"time"
)

func main75() {

	//00:00:00 UTC on Jan 1, 1970

	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
	now := time.Now().In(loc)
	unixTime := now.Unix()

	t := time.Unix(unixTime, 0)

	fmt.Println("Current unix time:", unixTime)
	fmt.Println(t)
}
