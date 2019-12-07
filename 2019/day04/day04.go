package main

import (
	"fmt"
)

func main() {
	start := 246540
	end := 787419

	numValid := 0
	
	for i := start; i < end; i++ {
		temp := i
		potential := []int{};
		for j := 0; j < 6; j++ {
			potential = append(potential, temp % 10)
			temp = temp / 10
		}

		neverDecrease := true
		adjacentNum := []int{1}
		
		for j := 5; j >= 1; j-- {

			if potential[j] == potential[j-1] {
				adjacentNum[len(adjacentNum) - 1]++
			} else if potential[j] > potential[j-1] {
				neverDecrease = false
				break
			} else {
				adjacentNum = append(adjacentNum, 1)
			}
		}

		dupeAdjacent := false
		for _, num := range adjacentNum {
			if num == 2 {
				dupeAdjacent = true;
			}
		}

		if dupeAdjacent && neverDecrease {
			numValid++
		}
	}

	fmt.Println(numValid)
}
