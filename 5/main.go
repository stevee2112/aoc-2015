package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"regexp"
	"strings"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	niceCounter1 := 0
	niceCounter2 := 0

	for scanner.Scan() {
		string := scanner.Text()

		if isNicePart1(string) {
			niceCounter1++
		}

		if isNicePart2(string) {
			niceCounter2++
		}

	}

	fmt.Printf("Part 1: %d\n", niceCounter1)
	fmt.Printf("Part 2: %d\n", niceCounter2)
}

func isNicePart1(s string) bool {

	// It contains at least three vowels (aeiou only)
	vowels, _ := regexp.Compile("[aeiou].*[aeiou].*[aeiou]")
	if !vowels.MatchString(s) {
		return false
	}

	// It contains at least one letter that appears twice in a row
	hasTwoInARow := false
	for index, value := range s {
		if (index + 1) != len(s) {
			hasTwoInARow = (string(value) == string(s[index + 1]))
		}

		if hasTwoInARow {
			break
		}
	}

	if !hasTwoInARow {
		return false
	}

	// It does not contain the strings ab, cd, pq, or xy
	badStrings := []string{"ab", "cd", "pq", "xy"}

	for _, bad := range badStrings {
		if strings.Contains(s, bad) {
			return false
		}
	}

	return true
}

func isNicePart2(s string) bool {
	// TODO
	return true
}
