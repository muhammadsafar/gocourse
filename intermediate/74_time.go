package main

import (
	"fmt"
	"time"
)

func main74() {

	now := time.Now()
	fmt.Println(now)

	specificTime := time.Date(2024, time.November, 12, 20, 0, 0, 0, time.Local)
	fmt.Println("Specific time >>", specificTime)

	//parse time harus Mon Jan 2 15:04:05 MST 2006 -> 2006-01-02
	parsedTime, _ := time.Parse("2006-01-02", "2002-10-10") //format_layout, valueToParse
	parsedTime1, _ := time.Parse("06-01-02", "20-10-10")
	parsedTime2, _ := time.Parse("06-1-2", "20-5-9")
	parsedTime3, _ := time.Parse("06-1-2 15-04", "20-5-9 18-05")
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	//Formatting time

	t := time.Now()
	fmt.Println(`Formatting time >>`, t.Format("06-01-02 15-04-05"))

	oneDayLatr := t.Add(time.Hour * 24) //add 24 jam
	fmt.Println("One day later >>", oneDayLatr.Format("2006-01-02 15:04:05"))
	fmt.Println("weekday>>", oneDayLatr.Weekday())

	fmt.Println("Rounded time hour:", t.Round(time.Hour)) //round terdekat, <30 bulat ke bawah, >30 bulat ke atas

	loc, _ := time.LoadLocation("Asia/Jakarta")
	t = time.Date(2024, time.December, 8, 16, 23, 54, 00, time.UTC)

	//convert to specfov time zone
	tlocal := t.In(loc)

	//preform rounding
	roundedTime := t.Round(time.Hour)
	roudedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original time (UTC)", t)
	fmt.Println("Original time (Local)", tlocal)
	fmt.Println("Rounded time (UTC)", roundedTime)
	fmt.Println("Rounded time (Local)", roudedTimeLocal)

	fmt.Println("Truncated time : ", t.Truncate(time.Hour))

	loc2, _ := time.LoadLocation("America/New_York")
	t1NY := time.Now().In(loc2)
	fmt.Println("New York time >>", t1NY)

	t1 := time.Date(2025, time.August, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, time.August, 4, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("time 1>>", t1)
	fmt.Println("time 2>>", t2)
	fmt.Println("difference t1 and t2 >>", duration)

	//compare times return true
	fmt.Println("t2 is after t1", t2.After(t1))
}
