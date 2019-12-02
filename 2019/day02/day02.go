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

	pint := make([]int, len(strIntcode))

	for x := 0; x < len(strIntcode); x++ {
		pint[x], _ = strconv.Atoi(strIntcode[x])
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			intcode := make([]int, len(pint))
			copy(intcode, pint)
			
			intcode[1] = i
			intcode[2] = j
			
			for x := 0; x < len(intcode); x += 4 {
				optcode := intcode[x]

				if optcode == 99 {
					break
				}

				firstParam := intcode[x+1]
				secondParam := intcode[x+2]
				finalParam := intcode[x+3]

				if optcode == 1 {
					intcode[finalParam] = intcode[firstParam] + intcode[secondParam]
				} else if optcode == 2 {
					intcode[finalParam] = intcode[firstParam] * intcode[secondParam]
				}
			}

			// Part 1 Answer
			fmt.Println(intcode)

			// Part 2 Answer
			if intcode[0] == 19690720 {
				fmt.Println(100 * i + j)
				break
			}
		}
	}
}
