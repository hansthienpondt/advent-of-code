package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	a1, a2 int
	b1, b2 int
}

func (c *coord) String() string {
	return fmt.Sprintf("%d-%d,%d-%d", c.a1, c.a2, c.b1, c.b2)
}

func (c *coord) Overlaps() bool {
	if c.a1 >= c.b1 && c.a1 <= c.b2 {
		return true
	}

	if c.a2 >= c.b1 && c.a2 <= c.b2 {
		return true
	}

	if c.b1 >= c.a1 && c.b1 <= c.a2 {
		return true
	}

	if c.b2 >= c.a1 && c.b2 <= c.a2 {
		return true
	}

	return false
}

func (c *coord) Contained() bool {
	if c.a1 >= c.b1 && c.a2 <= c.b2 {
		return true
	}
	if c.b1 >= c.a1 && c.b2 <= c.a2 {
		return true
	}
	return false
}
func newCoord(s_a1, s_a2, s_b1, s_b2 string) *coord {
	var a1, a2 int
	var b1, b2 int
	a1, _ = strconv.Atoi(s_a1)
	a2, _ = strconv.Atoi(s_a2)
	b1, _ = strconv.Atoi(s_b1)
	b2, _ = strconv.Atoi(s_b2)

	c := coord{a1: a1, a2: a2, b1: b1, b2: b2}
	return &c
}

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	return lines
}

func day4a(input []string) int {
	var l_coord []*coord

	// process each newline in input
	for _, l := range input {
		// for each line, process each team (2)
		line := strings.FieldsFunc(l, CustomSplit)
		l_coord = append(l_coord, newCoord(line[0], line[1], line[2], line[3]))
	}
	var count int = 0
	for _, c := range l_coord {
		if c.Contained() {
			count += 1
		}
	}
	return count
}
func day4b(input []string) int {
	var l_coord []*coord

	// process each newline in input
	for _, l := range input {
		// for each line, process each team (2)
		line := strings.FieldsFunc(l, CustomSplit)
		l_coord = append(l_coord, newCoord(line[0], line[1], line[2], line[3]))
	}
	var count int = 0
	for _, c := range l_coord {
		if c.Overlaps() {
			count += 1
		}
	}
	return count
}

func CustomSplit(r rune) bool {
	return r == ',' || r == '-'
}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day4a(readInput(filename)))
	fmt.Printf("Part 2: %d\n", day4b(readInput(filename)))

}
