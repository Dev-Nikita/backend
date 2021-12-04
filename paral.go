package main

import (
	"fmt"
	"time"
)

func printnumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printletters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printnumbers()
	go printletters()
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Printing from main")
}
