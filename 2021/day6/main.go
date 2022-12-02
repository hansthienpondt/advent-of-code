package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type School struct {
	Age map[int]int
}

func NewSchool() *School {
	m := make(map[int]int, 9)
	for i := 0; i < 9; i++ {
		m[i] = 0
	}
	return &School{Age: m}
}
func (s *School) Init(fish []int) {
	for _, v := range fish {
		s.Age[v] += 1
	}
}
func (s *School) DecrementAge() {
	var ages []int
	for key := range s.Age {
		ages = append(ages, key)
	}
	sort.Ints(ages)
	m := make(map[int]int, 9)

	for i := len(ages) - 1; i >= 0; i-- {
		if i == 0 {
			//fmt.Printf("Adding %d fishes\n", s.Age[0])
			m[8] = s.Age[0]
			m[6] += s.Age[0]
		} else {
			m[i-1] = s.Age[i]
		}

	}
	s.Age = m
}
func (s *School) Sum() int {
	sum := 0
	for _, v := range s.Age {
		sum += v
	}
	return sum
}
func readInput(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)

	return data
}
func parseInput(input string) (fish []int) {
	for _, number := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(number)
		fish = append(fish, n)
	}
	return fish
}

func part1(fishes []int, days int) int {
	school := NewSchool()
	school.Init(fishes)

	for i := 1; i <= days; i++ {
		school.DecrementAge()
	}
	return school.Sum()
}
func main() {
	filename := os.Args[1]
	fish := parseInput(readInput(filename))
	p1days := 80
	amount_part1 := part1(fish, p1days)
	fmt.Printf("Part1: after %d days, there are %d fish\n", p1days, amount_part1)

	p2days := 256
	amount_part2 := part1(fish, p2days)
	fmt.Printf("Part2: after %d days, there are %d fish\n", p2days, amount_part2)

	p3days := 1000
	amount_part3 := part1(fish, p3days)
	fmt.Printf("Part2: after %d days, there are %d fish\n", p3days, amount_part3)

}
