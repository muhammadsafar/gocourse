package main

import (
	"flag"
	"fmt"
	"os"
)

func main91() {

	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstflag := subCommand1.Bool("processing", false, "Command proccessing status")
	secondflag := subCommand1.Int("byte", 1024, "byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("this program require additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subcommand1:")
		fmt.Println("Processing status:", *firstflag)
		fmt.Println("bytes:", *secondflag)

	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subcommand2:")
		fmt.Println("language:", *flagsc2)

	default:
		fmt.Println("unknown subcommand")
		os.Exit(1)
	}

}
