package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MonkeyList []*Monkey

func (ml MonkeyList) Len() int           { return len(ml) }
func (ml MonkeyList) Less(i, j int) bool { return ml[i].itemCounter < ml[j].itemCounter }
func (ml MonkeyList) Swap(i, j int)      { ml[i], ml[j] = ml[j], ml[i] }

type Monkey struct {
	id          int
	sItem       []int
	itemCounter int
	operation   func(int) int
	boredFunc   func(int) int
	testFunc    func(int) int
	divisibleBy int
}

func (m *Monkey) Inspect(l MonkeyList) {
	//fmt.Printf("Monkey %d:\n", m.id)
	// inspect each item
	for _, item := range m.sItem {
		var worryLevel int = 0
		//fmt.Printf("   Monkey inspects an item with a worry level of %d.\n", item)
		// increase worrylevel
		worryLevel += m.operation(item)
		//fmt.Printf("     Worry level is increased to %d.\n", worryLevel)
		// Monkey gets bored
		worryLevel = m.boredFunc(worryLevel)
		//fmt.Printf("     Monkey gets bored with item. Worrylevel %d.\n", worryLevel)
		// conduct the test
		dest := m.testFunc(worryLevel)
		//fmt.Printf("     Item with worry level %d is thrown to monkey %d.\n", worryLevel, dest)
		l[dest].Append(worryLevel)
		// increment the item counter
		m.itemCounter += 1
	}
	m.sItem = []int{}
}

func (m *Monkey) Append(i int) {
	m.sItem = append(m.sItem, i)
}

func (m *Monkey) ID() int {
	return m.id
}

func (m *Monkey) String() string {
	return fmt.Sprintf("%d", m.ID())
}

func (m *Monkey) Contents() string {
	return fmt.Sprintf("Monkey %d: %v", m.id, m.sItem)
}
func (m *Monkey) ItemCounter() string {
	return fmt.Sprintf("Monkey %d inspected items %d times.", m.ID(), m.itemCounter)
}

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n\n")

	return lines
}

func day11a(input []string) int {
	ml := make(MonkeyList, 0)
	bfunc := parseOperandFunc("/", 3)

	for _, l := range input {
		ml = append(ml, parseMonkey(l, bfunc))
	}
	var rounds int = 20

	for r := 0; r < rounds; r++ {
		for _, m := range ml {
			m.Inspect(ml)
		}
	}
	sort.Sort(ml)
	var count int = 1
	for _, m := range ml[len(ml)-2 : len(ml)] {
		count *= m.itemCounter
		//fmt.Println(m.ItemCounter())
	}
	return count
}

func day11b(input []string) int {
	ml := make(MonkeyList, 0)
	var largest int = 1

	bfunc := parseOperandFunc("/", 1)

	for _, l := range input {
		ml = append(ml, parseMonkey(l, bfunc))
	}
	for _, m := range ml {
		largest *= m.divisibleBy
	}
	for _, m := range ml {
		m.boredFunc = parseOperandFunc("%", largest)
	}
	var rounds int = 10000

	for r := 0; r < rounds; r++ {
		for _, m := range ml {
			m.Inspect(ml)
		}
	}

	sort.Sort(ml)
	var count int = 1
	for _, m := range ml[len(ml)-2 : len(ml)] {
		count *= m.itemCounter
		//fmt.Println(m.ItemCounter())
	}
	//fmt.Println(count)
	return count
}

func parseMonkey(s string, boredfunc func(int) int) *Monkey {
	// replace comma seperated to make sscanf work
	// replace * old by operand digit (^ 2) to make it uniform
	s = strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(s)
	// fmt.Println(s)

	var monkey, divisible, operand_v, test_true, test_false int
	var s_items, operand string

	template := `Monkey %d:
		Starting items: %v
		Operation: new = old %s %d
		Test: divisible by %d
		  If true: throw to monkey %d
		  If false: throw to monkey %d`

	fmt.Sscanf(s, template, &monkey, &s_items, &operand, &operand_v, &divisible, &test_true, &test_false)
	//fmt.Println(monkey, toIntList(string(s_items)), string(operand), operand_v, divisible, test_true, test_false)

	m := Monkey{
		id:          monkey,
		sItem:       toIntList(string(s_items)),
		itemCounter: 0,
		operation:   parseOperandFunc(string(operand), operand_v),
		boredFunc:   boredfunc,
		testFunc:    parseTestFunc(divisible, test_true, test_false),
		divisibleBy: divisible}

	return &m
}

func toIntList(s string) (l []int) {
	for _, i := range strings.Split(s, ",") {
		j, _ := strconv.Atoi(i)
		l = append(l, j)
	}
	return
}
func parseOperandFunc(operand string, value int) func(int) int {
	switch operand {
	case "+":
		return func(i int) int {
			return i + value
		}
	case "*":
		return func(i int) int {
			return i * value
		}
	case "^":
		return func(i int) int {
			return int(math.Pow(float64(i), float64(value)))
		}
	case "/":
		return func(i int) int {
			return i / value
		}
	case "%":
		return func(i int) int {
			return i % value
		}
	}
	return func(i int) int { return i }
}
func parseTestFunc(j, test_true, test_false int) func(int) int {
	return func(i int) int {
		if i%j == 0 {
			return test_true
		} else {
			return test_false
		}
	}
}
func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day11a(readInput(filename)))
	fmt.Printf("Part 2: %d\n", day11b(readInput(filename)))

}
