package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseFirstInstruction(first int) (int, []int) {
	opcode := first % 100
	paramModes := []int{}
	first /= 100

	for first > 0 {
		paramModes = append(paramModes, first%10)
		first /= 10
	}

	return opcode, paramModes
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(file), ",")

	pint := make([]int, len(input))

	for x := 0; x < len(pint); x++ {
		pint[x], _ = strconv.Atoi(input[x])
	}

loop:
	for i := 0; i < len(pint); {
		opcode, modes := parseFirstInstruction(pint[i])
		
		params := []int{pint[i+1], pint[i+2], pint[i+3]}

		for j, mode := range modes {
			if mode == 1 {
				params[j] = i+j+1
			}
		}

		switch opcode {
		case 1:
			pint[params[2]] = pint[params[0]] + pint[params[1]]
			i += 4
		case 2:
			pint[params[2]] = pint[params[0]] * pint[params[1]]
			i += 4
		case 3:
			pint[params[0]] = 1
			i += 2
		case 4:
			fmt.Println(pint[params[0]])
			i += 2
		case 99:
			break loop
		}
	}
}
