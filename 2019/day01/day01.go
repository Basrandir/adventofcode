package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var total int = 0

	for scanner.Scan() {
		module, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}
		
		total += massToFuel(module)
	}
	
	file.Close()
	
	fmt.Println(total)
}

func massToFuel(x int) int {
	fuel := int(math.Floor(float64(x) / 3.0)) -2

	if fuel <= 0 {
		return 0
	} else {
		return fuel + massToFuel(fuel)
	}
}
