package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)
	
func main() {
	file, _ := ioutil.ReadFile("input.txt")
	strIntcode := strings.Split(string(file), ",")

	intcode := make([]int, len(strIntcode))

	for x := 0; x < len(strIntcode); x++ {
		intcode[x], _ = strconv.Atoi(strIntcode[x])
	}

	for x := 0; x < len(intcode); x += 4 {
		optcode := intcode[x]

		if optcode == 99 {
			break
		}

		first := intcode[x+1]
		second := intcode[x+2]
		final := intcode[x+3]

		if optcode == 1 {
			intcode[final] = intcode[first] + intcode[second]
		} else if optcode == 2 {
			intcode[final] = intcode[first] * intcode[second]
		}
	}

	fmt.Println(intcode)
}
