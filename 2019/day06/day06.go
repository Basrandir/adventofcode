package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getPaths(count int, m map[string]string, val string) int {
	if val == "" {
		count -= 1
		return count
	} else {
		count += 1
		return getPaths(count, m, m[val])
	}
}

func findCommonNode(n1, n2 string, m map[string]string) int {
	firstPath := []string{}

	for m[n1] != "" {
		firstPath = append(firstPath, m[n1])
		n1 = m[n1]
	}

	l1 := 0
	for m[n2] != "" {
		l2 := 0
		for _, p := range firstPath {
			if p == m[n2] {
				return l1+l2
			}
			l2++
		}
		n2 = m[n2]
		l1++
	}
	return 0
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(file), "\n")
	input = input[:len(input)-1]

	/*
	input = []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN"}
*/

	m := map[string]string{}

	for _, el := range input {
		orbit := strings.Split(el, ")")
		m[orbit[1]] = orbit[0]
	}

	// Part 1
	count := 0
	for key, _ := range m {
		count = getPaths(count, m, key)
	}
	fmt.Println(count)

	// Part 2
	count = findCommonNode("YOU", "SAN", m)
	fmt.Println(count)
}
