package main

import (
	"day4/game"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)

	return data
}
func parseInput(input string) (drawNumbers, boardNumbers []int) {
	var drawStrings, boardStrings []string
	drawStrings = strings.Split(strings.Split(input, "\n")[0], ",")
	boardStrings = strings.Fields(input)[1:]

	for _, number := range drawStrings {
		n, _ := strconv.Atoi(number)
		drawNumbers = append(drawNumbers, n)
	}
	for _, number := range boardStrings {
		n, _ := strconv.Atoi(number)
		boardNumbers = append(boardNumbers, n)
	}
	return
}
func main() {
	filename := os.Args[1]
	drawNumbers, boardNumbers := parseInput(readInput(filename))
	fmt.Println("1st Puzzle")
	fmt.Println(strings.Repeat("#", 64))
	bingo1 := new(game.Bingo)
	bingo1.Initialize(boardNumbers)
	bingo1.RunGame(drawNumbers)

	fmt.Println("2nd Puzzle")
	fmt.Println(strings.Repeat("#", 64))

	bingo2 := new(game.Bingo)
	bingo2.Initialize(boardNumbers)
	bingo2.RunGame2(drawNumbers)

}
