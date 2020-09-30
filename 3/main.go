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

	directedGraphAll := util.NewDirectedGraph("")
	directedGraphSanta := util.NewDirectedGraph("")
	directedGraphRoboSanta := util.NewDirectedGraph("")

	whichSanta := 0
	santas := []*util.DirectedGraph{directedGraphSanta, directedGraphRoboSanta}

	for scanner.Scan() {
		direction := scanner.Text()

		switch (direction) {
		case "^":
			directedGraphAll.Move(util.North)
			santas[whichSanta].Move(util.North)
		case "v":
			directedGraphAll.Move(util.South)
			santas[whichSanta].Move(util.South)
		case "<":
			directedGraphAll.Move(util.West)
			santas[whichSanta].Move(util.West)
		case ">":
			directedGraphAll.Move(util.East)
			santas[whichSanta].Move(util.East)
		}

		whichSanta = 1 - whichSanta
	}

	// Combine santas
	combinedVisits := map[string]int{}
	for coor, visits := range santas[0].Visits {
		combinedVisits[coor] += visits
	}
	for coor, visits := range santas[1].Visits {
		combinedVisits[coor] += visits
	}

	fmt.Printf("Part 1: %d\n", len(directedGraphAll.Visits))
	fmt.Printf("Part 2: %d\n", len(combinedVisits))
}
