package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	args := os.Args
	
	if len(args) == 1 {
		printUsage()
		os.Exit(0)
		return
	}

	if args[1] == "-h" || args[1] == "-help" {
		printUsage()
		os.Exit(0)
		return
	}
	remainingTime, err := time.ParseDuration(args[1])
	if err != nil {
		fmt.Println("Invalid Duration format. Please run again with a valid Duration.")
		os.Exit(1)
		return
	}


	for remainingTime > 0 {
		clearLine()
		fmt.Print(remainingTime.String())

		time.Sleep(time.Second)
		remainingTime = remainingTime - time.Second
		
	}
	clearLine()
	beeep.Alert("Hey","Timer is over", "")
	fmt.Println("Done!")
	os.Exit(0)
}

func clearLine(){
	fmt.Print("\033[2K")
	fmt.Print("\r")
}

func printUsage(){
	fmt.Print("Timer\n",
	"Usage: timer [**h][**m][**s]\n")
}
