package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n\n")

	return lines
}
func parseInput(input []string) (pairs [][]any, err error) {
	pairs = make([][]any, 0)

	for _, l := range input {
		var left, right any
		pair := make([]any, 2, 2)
		line := strings.Split(l, "\n")

		if err := json.Unmarshal([]byte(line[0]), &left); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(line[1]), &right); err != nil {
			return nil, err
		}

		pair[0] = left
		pair[1] = right

		pairs = append(pairs, pair)

	}
	return pairs, nil
}
func isLeftSmaller(left any, right any) int {
	l, okL := left.(float64)
	r, okR := right.(float64)

	if okL && okR {
		return int(l) - int(r)
	}

	var lList []any
	var rList []any

	switch left.(type) {
	case []any, []float64:
		lList = left.([]any)
	case float64:
		lList = []any{left}
	}

	switch right.(type) {
	case []any, []float64:
		rList = right.([]any)
	case float64:
		rList = []any{right}
	}

	for i := range lList {
		if len(rList) <= i {
			return 1
		}
		//fmt.Printf("Compare %v vs %v\n", lList[i], rList[i])
		if r := isLeftSmaller(lList[i], rList[i]); r != 0 {
			return r
		}
	}
	if len(lList) == len(rList) {
		return 0
	}
	return -1
}

func day13a(input []string) int {
	pairs, _ := parseInput(input)

	count := 0

	for index, pair := range pairs {
		if isLeftSmaller(pair[0], pair[1]) <= 0 {
			count += index + 1
		}
	}

	return count
}

func day13b(input []string) int {
	input = append(input, []string{"[[2]]\n[[6]]"}...)

	pairs, _ := parseInput(input)

	allPackets := make([]any, 0, len(pairs)*2)

	for _, pair := range pairs {
		allPackets = append(allPackets, pair...)
	}

	sort.Slice(allPackets, func(i, j int) bool {
		l := allPackets[i]
		r := allPackets[j]
		return isLeftSmaller(l, r) < 0
	})

	var count int = 1
	for index, packet := range allPackets {
		str := fmt.Sprintf("%v", packet)
		if string(str) == "[[2]]" || string(str) == "[[6]]" {
			count *= index + 1
		}
	}

	return count
}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day13a(readInput(filename)))
	fmt.Printf("Part 2: %d\n", day13b(readInput(filename)))
}
