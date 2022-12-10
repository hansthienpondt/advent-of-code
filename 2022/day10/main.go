package main

import (
	"fmt"
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

func day10a(input []string) int {
	cycleCheck := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true}

	cycles := make([]int, 0)
	for _, l := range input {
		if l == "noop" {
			cycles = append(cycles, 0)
		} else {
			var command string
			var x int
			fmt.Sscanf(l, "%s %d", &command, &x)
			cycles = append(cycles, 0, x)
		}
	}

	//fmt.Println(cycles, len(cycles))
	var x int = 1
	var count int = 0
	for i := 0; i < len(cycles); i++ {
		if cycleCheck[i+1] {
			count += x * (i + 1)
			//fmt.Println(i+1, x, count)
		}
		x += cycles[i]
	}
	return count
}

func day10b(input []string, visibility bool) string {
	cycles := make([]int, 0)
	for _, l := range input {
		if l == "noop" {
			cycles = append(cycles, 0)
		} else {
			var command string
			var x int
			fmt.Sscanf(l, "%s %d", &command, &x)
			cycles = append(cycles, 0, x)
		}
	}

	//fmt.Println(cycles, len(cycles))
	var x int = 1
	var result string = ""
	for i := 0; i < len(cycles); i++ {
		p := i % 40

		if p == 0 {
			result += fmt.Sprintf("\n")
		}
		if p >= x-1 && p <= x+1 {
			result += fmt.Sprintf("#")
		} else {
			if visibility {
				result += fmt.Sprintf(" ")
			} else {
				result += fmt.Sprintf(".")
			}
		}
		x += cycles[i]
	}
	return result
}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day10a(readInput(filename)))
	fmt.Printf("Part 2: %s\n", day10b(readInput(filename), true))
}
