package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	args := os.Args
	timer, err := strconv.ParseInt(args[1], 10, 64)
	
	var unit time.Duration 

	if err != nil {
		fmt.Println("Invalid integer entered. Please run a valid integer.")
		return
	}

	switch input := args[2]; input {
	case "s":
		unit = time.Second
	case "m":
		unit = time.Minute
	case "h":
		unit = time.Hour
	default:
		fmt.Println("Invalid unit. Accepted units: s(econd), m(inute), h(our)")
		return
	}

	remainingTime :=  time.Duration(timer) * unit
	
	for remainingTime > 0 {
		clearLine()
		fmt.Print(remainingTime.String())

		time.Sleep(time.Second)
		remainingTime = remainingTime - time.Second
		
	}
	clearLine()
	beeep.Alert("Hey","Timer is over", "")
	fmt.Println("Done!")
}

func clearLine(){
	fmt.Print("\033[2K")
	fmt.Print("\r")
}
