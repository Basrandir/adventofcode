package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		module, _ := strconv.Atoi(scanner.Text())
		total += massToFuel1(module)
	}

	file.Close()

	fmt.Println(total)
}

// For Part 2
func massToFuel1(x int) int {
	fuel := x/3 - 2
	return fuel
}

// For Part 2
func massToFuel2(x int) int {
	fuel := x/3 - 2

	if fuel <= 0 {
		return 0
	} else {
		return fuel + massToFuel2(fuel)
	}
}
