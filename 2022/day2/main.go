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

func day2a(input []string, points map[string]int) int {
	var score int = 0
	for _, line := range input {
		fields := strings.Split(line, " ")

		if fields[0] == "A" && fields[1] == "X" {
			score += points["Rock"]
			score += points["Draw"]
		} else if fields[0] == "A" && fields[1] == "Y" {
			score += points["Paper"]
			score += points["Win"]
		} else if fields[0] == "A" && fields[1] == "Z" {
			score += points["Scissors"]
			score += points["Loss"]
		} else if fields[0] == "B" && fields[1] == "X" {
			score += points["Rock"]
			score += points["Loss"]
		} else if fields[0] == "B" && fields[1] == "Y" {
			score += points["Paper"]
			score += points["Draw"]
		} else if fields[0] == "B" && fields[1] == "Z" {
			score += points["Scissors"]
			score += points["Win"]
		} else if fields[0] == "C" && fields[1] == "X" {
			score += points["Rock"]
			score += points["Win"]
		} else if fields[0] == "C" && fields[1] == "Y" {
			score += points["Paper"]
			score += points["Loss"]
		} else if fields[0] == "C" && fields[1] == "Z" {
			score += points["Scissors"]
			score += points["Draw"]
		}
	}
	return score
}

func day2b(input []string, points map[string]int) int {
	var score int = 0
	for _, line := range input {
		fields := strings.Split(line, " ")

		if fields[0] == "A" && fields[1] == "X" {
			score += points["Scissors"]
			score += points["Loss"]
		} else if fields[0] == "A" && fields[1] == "Y" {
			score += points["Rock"]
			score += points["Draw"]
		} else if fields[0] == "A" && fields[1] == "Z" {
			score += points["Paper"]
			score += points["Win"]
		} else if fields[0] == "B" && fields[1] == "X" {
			score += points["Rock"]
			score += points["Loss"]
		} else if fields[0] == "B" && fields[1] == "Y" {
			score += points["Paper"]
			score += points["Draw"]
		} else if fields[0] == "B" && fields[1] == "Z" {
			score += points["Scissors"]
			score += points["Win"]
		} else if fields[0] == "C" && fields[1] == "X" {
			score += points["Paper"]
			score += points["Loss"]
		} else if fields[0] == "C" && fields[1] == "Y" {
			score += points["Scissors"]
			score += points["Draw"]
		} else if fields[0] == "C" && fields[1] == "Z" {
			score += points["Rock"]
			score += points["Win"]
		}
	}
	return score
}

func main() {
	filename := os.Args[1]
	// A: ROCK, B: PAPER, C: SCISSORS, X: ROCK, Y: PAPER, Z: SCISSORS
	points := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3, "Win": 6, "Loss": 0, "Draw": 3}

	fmt.Printf("Score according to strategy guide would be %d\n", day2a(readInput(filename), points))
	fmt.Printf("Score according to ultra top secret strategy guide, would be %d\n", day2b(readInput(filename), points))
}
