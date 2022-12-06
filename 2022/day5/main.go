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

//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// [S]                 [T] [Q]
// [L]             [B] [M] [P]     [T]
// [F]     [S]     [Z] [N] [S]     [R]
// [Z] [R] [N]     [R] [D] [F]     [V]
// [D] [Z] [H] [J] [W] [G] [W]     [G]
// [B] [M] [C] [F] [H] [Z] [N] [R] [L]
// [R] [B] [L] [C] [G] [J] [L] [Z] [C]
// [H] [T] [Z] [S] [P] [V] [G] [M] [M]
//  1   2   3   4   5   6   7   8   9

type move struct {
	amount int
	from   int
	to     int
}

func (m *move) String() string {
	return fmt.Sprintf("%d %d %d", m.amount, m.from, m.to)
}

func newMove(s_amount, s_from, s_to string) *move {
	var amount, from, to int
	amount, _ = strconv.Atoi(s_amount)
	from, _ = strconv.Atoi(s_from)
	to, _ = strconv.Atoi(s_to)

	m := move{amount: amount, from: from, to: to}
	return &m
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
func processInput(input []string) []*move {
	var l_move []*move
	for _, l := range input {
		l = strings.ReplaceAll(l, "move ", "")
		l = strings.ReplaceAll(l, " from ", " ")
		l = strings.ReplaceAll(l, " to ", " ")
		line := strings.Split(l, " ")
		l_move = append(l_move, newMove(line[0], line[1], line[2]))
	}
	return l_move
}
func day5a(input []string, v_input map[int][]string) []string {
	i := processInput(input)
	for _, j := range i {
		var remaining, popped []string
		//fmt.Println(j.amount, j.from, j.to)
		remaining = v_input[j.from][:len(v_input[j.from])-j.amount]
		popped = v_input[j.from][len(v_input[j.from])-j.amount : len(v_input[j.from])]

		ReverseSlice(popped)

		v_input[j.to] = append(v_input[j.to], popped...)
		v_input[j.from] = remaining
	}
	result := []string{}

	keys := make([]int, 0, len(v_input))

	for k := range v_input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		result = append(result, v_input[k][len(v_input[k])-1])
	}
	return result
}
func day5b(input []string, v_input map[int][]string) []string {
	i := processInput(input)
	for _, j := range i {
		var remaining, popped []string
		//fmt.Println(j.amount, j.from, j.to)
		remaining = v_input[j.from][:len(v_input[j.from])-j.amount]
		popped = v_input[j.from][len(v_input[j.from])-j.amount : len(v_input[j.from])]

		v_input[j.to] = append(v_input[j.to], popped...)
		v_input[j.from] = remaining
	}
	result := []string{}

	keys := make([]int, 0, len(v_input))

	for k := range v_input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		result = append(result, v_input[k][len(v_input[k])-1])
	}

	return result
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func main() {
	/*
		var testInput5a = map[int][]string{
			1: []string{"Z", "N"},
			2: []string{"M", "C", "D"},
			3: []string{"P"}}
	*/
	var input5a = map[int][]string{
		1: []string{"H", "R", "B", "D", "Z", "F", "L", "S"},
		2: []string{"T", "B", "M", "Z", "R"},
		3: []string{"Z", "L", "C", "H", "N", "S"},
		4: []string{"S", "C", "F", "J"},
		5: []string{"P", "G", "H", "W", "R", "Z", "B"},
		6: []string{"V", "J", "Z", "G", "D", "N", "M", "T"},
		7: []string{"G", "L", "N", "W", "F", "S", "P", "Q"},
		8: []string{"M", "Z", "R"},
		9: []string{"M", "C", "L", "G", "V", "R", "T"}}

	/*
		var testInput5b = map[int][]string{
			1: []string{"Z", "N"},
			2: []string{"M", "C", "D"},
			3: []string{"P"}}
	*/

	var input5b = map[int][]string{
		1: []string{"H", "R", "B", "D", "Z", "F", "L", "S"},
		2: []string{"T", "B", "M", "Z", "R"},
		3: []string{"Z", "L", "C", "H", "N", "S"},
		4: []string{"S", "C", "F", "J"},
		5: []string{"P", "G", "H", "W", "R", "Z", "B"},
		6: []string{"V", "J", "Z", "G", "D", "N", "M", "T"},
		7: []string{"G", "L", "N", "W", "F", "S", "P", "Q"},
		8: []string{"M", "Z", "R"},
		9: []string{"M", "C", "L", "G", "V", "R", "T"}}

	filename := os.Args[1]

	fmt.Printf("Part 1: %s\n", strings.Join(day5a(readInput(filename), input5a), ""))
	fmt.Printf("Part 2: %s\n", strings.Join(day5b(readInput(filename), input5b), ""))
}
