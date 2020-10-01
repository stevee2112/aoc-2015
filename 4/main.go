package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"crypto/md5"
	"io"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Scan()

	secretKey := scanner.Text()
	foundAt1 := 0
	foundAt2 := 0

	found1 := false
	found2 := false
	for i := 0; (!(found1 && found2)) ; i++ {
		h := md5.New()
		io.WriteString(h, secretKey)
		io.WriteString(h, fmt.Sprintf("%d", i))
		hash := fmt.Sprintf("%x", h.Sum(nil))

		if !found1 {
			found1 = (hash[0:5] == "00000")
			if found1 {
				foundAt1 = i
				fmt.Printf("Part 1: %d\n", foundAt1)
			}
		}

		if !found2 {
			found2 = (hash[0:6] == "000000")
			if found2 {
				foundAt2 = i
				fmt.Printf("Part 2: %d\n", foundAt2)
			}
		}
	}
}
