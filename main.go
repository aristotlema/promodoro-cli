package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var length time.Duration = 25

func main() {
	clear()
	start := time.Now().Truncate(time.Second)
	finish := start.Add(time.Minute * length)

	
	for timeActive(start, finish, time.Now()) {
		fmt.Println("Start time:", start)
		fmt.Println("End time:", finish)
		fmt.Println("Remaining:", time.Until(finish).Truncate(time.Second))

		time.Sleep(time.Second)
		clear()
	}
}

func timeActive(start, finish, current time.Time) bool {
	return current.After(start) && current.Before(finish)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}