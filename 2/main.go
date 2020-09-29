package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"strings"
	"strconv"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	totalPaperNeeded := 0;
	totalRibbonNeeded := 0;

	for scanner.Scan() {
		dimensions := strings.Split(scanner.Text(), "x")

		totalPaperNeeded += paperNeeded(dimensions[0], dimensions[1], dimensions[2])
		totalRibbonNeeded += ribbonNeeded(dimensions[0], dimensions[1], dimensions[2])
	}

	fmt.Printf("Part 1: %d\n", totalPaperNeeded)
	fmt.Printf("Part 2: %d\n", totalRibbonNeeded)
}


func paperNeeded(ls, ws, hs string) int {

	l, _ := strconv.Atoi(ls)
	w, _ := strconv.Atoi(ws)
	h, _ := strconv.Atoi(hs)

	side1 := (l * w)
	smallest := side1

	side2 := (w * h)

	if side2 < smallest {
		smallest = side2
	}

	side3 := (h * l)

	if side3 < smallest {
		smallest = side3
	}

	return (2 * side1) + (2 * side2) + (2 * side3) + smallest
}

func ribbonNeeded(ls, ws, hs string) int {

	l, _ := strconv.Atoi(ls)
	w, _ := strconv.Atoi(ws)
	h, _ := strconv.Atoi(hs)

	box := []int{l, w, h}

	max := 0
	bow := 1
	var maxAt int
	for at, side := range box {
		bow *= side
		if side > max {
			maxAt = at
			max = side
		}
	}

	box = append(box[:maxAt], box[maxAt + 1:]...)

	needed := 0

	for _, side := range box {
		needed += side + side
	}

	return needed + bow;
}

