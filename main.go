package main

import (
	branfuck "Brainfuck/brainfuck"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.bf")
	if err != nil {
		log.Fatal(err)

	}
	branfuck.Start(file)
}
