package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type SubMarine struct {
	Z int64 // forward, backward
	Y int64 // up, down
}

type SubMarinev2 struct {
	*SubMarine
	aim int64
}
type movement struct {
	command string
	value   int64
}

func (s *SubMarine) Up(i int64) error {
	if s.Y+i > 0 {
		return fmt.Errorf("cannot go above 0 depth.")
	}
	s.Y += i
	return nil
}
func (s *SubMarine) Down(i int64) error {
	s.Y -= i
	return nil
}
func (s *SubMarine) Forward(i int64) error {
	s.Z += i
	return nil
}
func (s *SubMarine) Backward(i int64) error {
	s.Z -= i
	return nil
}

func (s *SubMarinev2) Up(i int64) error {
	s.aim -= i
	return nil
}
func (s *SubMarinev2) Down(i int64) error {
	s.aim += i
	return nil
}
func (s *SubMarinev2) Forward(i int64) error {
	s.Y -= i * s.aim
	s.Z += i
	return nil
}

func (s *SubMarine) Depth() int64 {
	return int64(math.Abs(float64(s.Y)))
}

func NewSubMarine() *SubMarine { return &SubMarine{Z: int64(0), Y: int64(0)} }
func NewSubMarinev2() *SubMarinev2 {
	return &SubMarinev2{SubMarine: NewSubMarine(), aim: int64(0)}
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
func parseInput(input []string) (m []movement) {
	for _, v := range input {
		move := strings.Split(v, " ")
		cmd := move[0]
		value, _ := strconv.Atoi(move[1])

		m = append(m, movement{command: cmd, value: int64(value)})
	}
	return m
}

func part1(input []movement) {
	var err error
	sub := NewSubMarine()

	for _, m := range input {
		switch m.command {
		case "forward":
			err = sub.Forward(m.value)
		case "backward":
			err = sub.Backward(m.value)
		case "up":
			err = sub.Up(m.value)
		case "down":
			err = sub.Down(m.value)
		}
		if err != nil {
			panic(err)
		}
	}
	// Debug
	// fmt.Println(sub)
	fmt.Printf("The resulting depth is %d and the horizontal position is %d\n", sub.Depth(), sub.Z)
	fmt.Printf("The overall result is %d\n", sub.Depth()*sub.Z)

}
func part2(input []movement) {
	var err error
	sub := NewSubMarinev2()

	for _, m := range input {
		switch m.command {
		case "forward":
			err = sub.Forward(m.value)
		case "backward":
			err = sub.Backward(m.value)
		case "up":
			err = sub.Up(m.value)
		case "down":
			err = sub.Down(m.value)
		}
		if err != nil {
			panic(err)
		}
	}
	// Debug
	// fmt.Println(sub)
	fmt.Printf("The resulting depth is %d and the horizontal position is %d\n", sub.Depth(), sub.Z)
	fmt.Printf("The overall result is %d\n", sub.Depth()*sub.Z)

}
func main() {
	//input := parseInput(readInput("testinput"))
	input := parseInput(readInput("input"))
	part1(input)
	part2(input)
}
