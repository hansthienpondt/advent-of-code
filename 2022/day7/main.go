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

func day7a(input []string) int {
	return 0
}
func day7b(input []string) int {
	return 0
}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day7a(readInput(filename)))
	fmt.Printf("Part 2: %d\n", day7b(readInput(filename)))

}
