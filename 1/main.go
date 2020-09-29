package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	at := 0
	position := 0
	atBasement := 0

	for scanner.Scan() {
		direction := scanner.Text()

		switch direction {
		case "(":
			at++
		case ")":
			at--
		}

		position++

		if at < 0 && atBasement == 0 {
			atBasement = position
		}
	}

	fmt.Printf("Part 1: %d\n", at)
	fmt.Printf("Part 2: %d\n", atBasement)
}
