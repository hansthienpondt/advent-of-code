package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	return lines
}

func day9a(input []string) int {
	directions := map[rune]image.Point{
		'U': {0, 1},
		'D': {0, -1},
		'L': {-1, 0},
		'R': {1, 0}}

	// Head position: rope[0] and Tail position: rope[1]
	rope := make([]image.Point, 2)

	logCount := map[image.Point]int{}

	// Start is on 2D X/Y-axis (0,0)
	for _, l := range input {
		var command rune
		var steps int

		fmt.Sscanf(l, "%c %d", &command, &steps)

		//fmt.Println(string(command), steps)

		for i := 0; i < steps; i++ {
			rope[0] = rope[0].Add(directions[command])

			if distance := rope[0].Sub(rope[1]); absolute(distance.X) > 1 || absolute(distance.Y) > 1 {
				rope[1] = rope[1].Add(image.Point{sign(distance.X), sign(distance.Y)})
			}
			//fmt.Println(string(command), steps, i, rope[0], rope[1])
			logCount[rope[1]] += 1
		}
	}

	return len(logCount)
}

func day9b(input []string) int {
	directions := map[rune]image.Point{
		'U': {0, 1},
		'D': {0, -1},
		'L': {-1, 0},
		'R': {1, 0}}

	// Head position: rope[0] and Tail position: rope[1]
	rope := make([]image.Point, 10)

	logCount := map[image.Point]int{}

	// Start is on 2D X/Y-axis (0,0)
	for _, l := range input {
		var command rune
		var steps int

		fmt.Sscanf(l, "%c %d", &command, &steps)

		//fmt.Println(string(command), steps)

		for i := 0; i < steps; i++ {
			rope[0] = rope[0].Add(directions[command])

			for j := 1; j < len(rope); j++ {
				if distance := rope[j-1].Sub(rope[j]); absolute(distance.X) > 1 || absolute(distance.Y) > 1 {
					rope[j] = rope[j].Add(image.Point{sign(distance.X), sign(distance.Y)})
				}
			}
			//fmt.Println(string(command), steps, i, rope[0], rope[1], rope[2], rope[3], rope[4], rope[5], rope[6], rope[7], rope[8], rope[9])
			logCount[rope[len(rope)-1]] += 1
		}
	}
	return len(logCount)
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day9a(readInput(filename)))
	if filename == "input" {
		fmt.Printf("Part 2: %d\n", day9b(readInput(filename)))
	} else {
		fmt.Printf("Part 2: %d\n", day9b(readInput(filename+"-part2")))
	}
}
