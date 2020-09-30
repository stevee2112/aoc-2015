package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"stevee2112/aoc-2015/util"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	directedGraph := util.NewDirectedGraph("")

	for scanner.Scan() {
		direction := scanner.Text()

		switch (direction) {
		case "^":
			directedGraph.Move(util.North)
		case "v":
			directedGraph.Move(util.South)
		case "<":
			directedGraph.Move(util.West)
		case ">":
			directedGraph.Move(util.East)
		}
	}

	fmt.Printf("Part 1: %d\n", len(directedGraph.Visits))
	//	fmt.Printf("Part 2: %d\n", atBasement)
}
