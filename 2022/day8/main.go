package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type TreeMap [][]int

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	return lines
}

func day8a(input []string) int {
	treeMap := make([][]int, len(input))

	for i := 0; i < len(input); i++ {
		strs := strings.Split(input[i], "")
		treeMap[i] = make([]int, len(strs))
		for j := 0; j < len(strs); j++ {
			item, _ := strconv.Atoi(strs[j])
			treeMap[i][j] = item
		}
	}
	count := 0
	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap)-1; j++ {
			if !Visible(i, j, treeMap) {
				count += 1
			}
		}
	}
	return (len(treeMap) * len(treeMap)) - count
}

func day8b(input []string) int {
	treeMap := make([][]int, len(input))

	for i := 0; i < len(input); i++ {
		strs := strings.Split(input[i], "")
		treeMap[i] = make([]int, len(strs))
		for j := 0; j < len(strs); j++ {
			item, _ := strconv.Atoi(strs[j])
			treeMap[i][j] = item
		}
	}
	var count = 0
	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap)-1; j++ {
			r := VisibleCount(j, i, treeMap)
			if r > count {
				count = r
			}
		}
	}
	return count
}

func Visible(x, y int, treeMap [][]int) bool {
	// process each column ; x = stable
	var top, bottom, left, right bool = true, true, true, true
	for i := 0; i < y; i++ {
		if treeMap[i][x] >= treeMap[y][x] {
			//fmt.Println(treeMap[i][x], treeMap[y][x])
			//fmt.Println("toggled top")
			top = false
		}
	}
	for i := y + 1; i < len(treeMap); i++ {
		if treeMap[i][x] >= treeMap[y][x] {
			//fmt.Println(treeMap[i][x], treeMap[y][x])
			//fmt.Println("toggled bottom")
			bottom = false
		}
	}
	// process each column ; y = stable
	for i := 0; i < x; i++ {
		if treeMap[y][i] >= treeMap[y][x] {
			//fmt.Println(treeMap[y][i], treeMap[y][x])
			//fmt.Println("toggled left")
			left = false
		}
	}
	for i := x + 1; i < len(treeMap); i++ {
		if treeMap[y][i] >= treeMap[y][x] {
			//fmt.Println(treeMap[y][i], treeMap[y][x])
			//fmt.Println("toggled right")
			right = false
		}
	}
	return (top || bottom) || (left || right)
}

func VisibleCount(x, y int, treeMap [][]int) int {
	// process each column ; x = stable
	var top, bottom, left, right int = 0, 0, 0, 0
	for i := y - 1; i >= 0; i-- {
		if treeMap[i][x] < treeMap[y][x] {
			//fmt.Println(treeMap[i][x], treeMap[y][x])
			//fmt.Println("toggled top")
			top += 1
		} else {
			top += 1
			break
		}
	}
	for i := y + 1; i < len(treeMap); i++ {
		if treeMap[i][x] < treeMap[y][x] {
			//fmt.Println(treeMap[i][x], treeMap[y][x])
			//fmt.Println("toggled bottom")
			bottom += 1
		} else {
			bottom += 1
			break
		}
	}
	// process each column ; y = stable
	for i := x - 1; i >= 0; i-- {
		if treeMap[y][i] < treeMap[y][x] {
			//fmt.Println(treeMap[y][i], treeMap[y][x])
			//fmt.Println("toggled left")
			left += 1
		} else {
			left += 1
			break
		}
	}
	for i := x + 1; i < len(treeMap); i++ {
		if treeMap[y][i] < treeMap[y][x] {
			//fmt.Println(treeMap[y][i], treeMap[y][x])
			//fmt.Println("toggled right")
			right += 1
		} else {
			right += 1
			break
		}
	}
	return top * bottom * left * right
}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day8a(readInput(filename)))
	fmt.Printf("Part 2: %d\n", day8b(readInput(filename)))

}
