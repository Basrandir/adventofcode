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

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(file), "\n")
	input = input[:len(input)-1]

	/*input = []string{
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
		"K)L"}*/

	m := map[string]string{}

	for _, el := range input {
		orbit := strings.Split(el, ")")
		m[orbit[1]] = orbit[0]
	}

	count := 0
	for key, _ := range m {
		count = getPaths(count, m, key)
	}

	fmt.Println(count)

}
