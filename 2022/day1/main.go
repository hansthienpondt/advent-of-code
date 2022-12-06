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

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	return lines
}

func day1a(input []string) (raw map[int][]int, elf map[int]int) {
	raw = make(map[int][]int)
	elf = make(map[int]int)

	elfcount := 0

	for _, v := range input {
		if len(v) == 0 {
			//fmt.Println("new Elf!")
			elfcount += 1
			continue
		}
		n, _ := strconv.Atoi(v)
		raw[elfcount+1] = append(raw[elfcount+1], n)
		elf[elfcount+1] += n
	}
	return
}

func main() {
	filename := os.Args[1]
	_, elf := day1a(readInput(filename))

	highkey := 0
	highvalue := 0

	for k, v := range elf {
		if v > highvalue {
			highvalue = v
			highkey = k
		}
	}
	fmt.Printf("Elf %d carries %d Calories\n", highkey, highvalue)

	var elfList, topthree []int
	for _, v := range elf {
		elfList = append(elfList, v)
	}
	sort.Ints(elfList)

	topthree = elfList[len(elfList)-3 : len(elfList)]

	sum := 0
	for _, v := range topthree {
		sum += v
	}
	fmt.Printf("The top three carries %d Calories\n", sum)
}
