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

func day6a(input []string) int {
	for _, v := range input {
		for i := 0; i < len(v)-3; i++ {
			if HasDuplicate(v[i : i+4]) {
				continue
			} else {
				return i + 4
				//fmt.Printf("%d %d %s %d\n", j, i+4, v[i:i+4], len(v))
				break
			}
		}
	}
	return 0
}
func day6b(input []string) int {
	for _, v := range input {
		for i := 0; i < len(v)-13; i++ {
			if HasDuplicate(v[i : i+14]) {
				continue
			} else {
				return i + 14
				//fmt.Printf("%d %d %s %d\n", j, i+14, v[i:i+14], len(v))
				break
			}
		}
	}
	return 0
}

func HasDuplicate(s string) bool {
	duplicate := make(map[string]int)
	for _, c := range s {
		_, exist := duplicate[string(c)]

		if exist {
			duplicate[string(c)] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate[string(c)] = 1 // else start counting from 1
		}
	}
	for _, v := range duplicate {
		if v > 1 {
			return true
		}
	}
	return false

}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day6a(readInput(filename)))
	fmt.Printf("Part 2: %d\n", day6b(readInput(filename)))

}
