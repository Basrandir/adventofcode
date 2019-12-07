package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"fmt"
	"strings"
)

type Coordinates struct {
	x, y, steps int
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")

	wires := strings.Split(string(file), "\n")

	//wires = []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72",
	//	"U62,R66,U55,R34,D71,R55,D58,R83"}
	
	points := [2][]Coordinates{}
	intersectionPoints := []Coordinates{}

	for i := 0; i < 2; i++ {
		x, y, steps := 0, 0, 0
		for _, path := range strings.Split(wires[i], ",") {
			direction := string(path[0])
			distance, _ := strconv.Atoi(path[1:])

			dx, dy := 0, 0

			switch direction {
			case "R":
				dx = 1
			case "L":
				dx = -1
			case "U":
				dy = 1
			case "D":
				dy = -1
			}

			for j := 0; j < distance; j++ {
				x += dx
				y += dy
				steps++
				
				points[i] = append(points[i], Coordinates{x, y, steps})
			}
		}
	}

	for _, firstWire := range points[0] {
		for _, secondWire := range points[1] {
			if firstWire.x == secondWire.x && firstWire.y == secondWire.y {
				intersectionPoints = append(intersectionPoints, Coordinates{firstWire.x, secondWire.y, firstWire.steps + secondWire.steps})
			}
		}
	}

	// Part 1
	min := math.Abs(float64(intersectionPoints[0].x)) + math.Abs(float64(intersectionPoints[0].y))
	for i := 1; i < len(intersectionPoints); i++ {
		current := math.Abs(float64(intersectionPoints[i].x)) + math.Abs(float64(intersectionPoints[i].y))
		if current < min {
			min = current
		}
	}

	fmt.Println(min)

	// Part 2
	min2 := intersectionPoints[0].steps
	for i := 1; i < len(intersectionPoints); i++ {
		current := intersectionPoints[i].steps
		if current < min2 {
			min2 = current
		}
	}

	fmt.Println(min2)
}
