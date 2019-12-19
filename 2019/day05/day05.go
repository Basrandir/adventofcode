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

	num := 0
	switch opcode {
	case 1:
		num = 3
	case 2:
		num = 3
	case 3:
		num = 1
	case 4:
		num = 1
	case 5:
		num = 2
	case 6:
		num = 2
	case 7:
		num = 3
	case 8:
		num = 3
	}

	for i := 0; i < num; i++ {
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
		pint[x], _ = strconv.Atoi(strings.TrimSpace(input[x]))
	}

loop:
	for i := 0; i < len(pint); {
		opcode, modes := parseFirstInstruction(pint[i])

		params := []int{}

		for j, mode := range modes {
			if mode == 1 {
				params = append(params, i+j+1)
			} else {
				params = append(params, pint[i+j+1])
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
			pint[params[0]] = 5
			i += 2
		case 4:
			fmt.Println(pint[params[0]])
			i += 2
		case 5:
			if pint[params[0]] != 0 {
				i = pint[params[1]]
			} else {
				i += 3
			}
		case 6:
			if pint[params[0]] == 0 {
				i = pint[params[1]]
			} else {
				i += 3
			}
		case 7:
			if pint[params[0]] < pint[params[1]] {
				pint[params[2]] = 1
			} else {
				pint[params[2]] = 0
			}
			i += 4
		case 8:
			if pint[params[0]] == pint[params[1]] {
				pint[params[2]] = 1
			} else {
				pint[params[2]] = 0
			}
			i += 4
		case 99:
			break loop
		}
	}
}
