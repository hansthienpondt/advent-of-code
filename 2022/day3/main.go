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

func day3a(input []string) []rune {
	var result []rune

	for _, line := range input {
		half := len(line) / 2
		partOne := line[0:half]
		partTwo := line[half:len(line)]

		result = append(result, commonLetter(partOne, partTwo))
	}
	return result
}
func commonLetter(partOne string, partTwo string) rune {
	var d rune = '🦡'

	for _, a := range partOne {
		for _, b := range partTwo {
			if a == b {
				return a
			}
		}
	}

	return d
}
func runesToScores(scores []rune) int {
	total := 0
	for _, s := range scores {
		score := int(s) - 96
		if score < 0 {
			score = int(s) - 64 + 26
		}
		total += score
	}
	return total
}

func day3b(input []string) []rune {
	scores := []rune{}
OUTER:
	for i := 0; i < len(input)/3; i++ {
		elf1 := input[i*3]
		elf2 := input[i*3+1]
		elf3 := input[i*3+2]
		sharedMap := map[rune]bool{}
		elf1Map := map[rune]bool{}
		for _, e1 := range elf1 {
			elf1Map[e1] = true
		}
		for _, e2 := range elf2 {
			sharedMap[e2] = elf1Map[e2]
		}
		for _, e3 := range elf3 {
			if sharedMap[e3] {
				scores = append(scores, e3)
				continue OUTER
			}
		}
	}
	return scores
}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", runesToScores(day3a(readInput(filename))))
	fmt.Printf("Part 2: %d\n", runesToScores(day3b(readInput(filename))))
}
