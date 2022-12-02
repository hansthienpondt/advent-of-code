package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Line struct {
	One *Coordinate
	Two *Coordinate
}

func (l *Line) UpdateLine(s string) {
	strCoordinates := strings.Split(s, " -> ")
	one := NewCoordinate(0, 0)
	two := NewCoordinate(0, 0)

	one.SetCoordinates(strCoordinates[0])
	two.SetCoordinates(strCoordinates[1])

	l.One = one
	l.Two = two
}

func (l *Line) SortbyX() *Line {
	if l.One.X <= l.Two.X {
		return &Line{One: l.One, Two: l.Two}
	} else {
		return &Line{One: l.Two, Two: l.One}
	}
}
func (l *Line) SortbyY() *Line {
	if l.One.Y <= l.Two.Y {
		return &Line{One: l.One, Two: l.Two}
	} else {
		return &Line{One: l.Two, Two: l.One}
	}
}
func (l *Line) String() string {
	return fmt.Sprintf("%s -> %s", l.One, l.Two)
}

// Line is a function: y = ax+b
func (l *Line) A() int {
	return (l.Two.Y - l.One.Y) / (l.Two.X - l.One.X)
}
func (l *Line) B() int {
	return l.One.Y - l.A()*l.One.X
}
func (l *Line) X(y int) int {
	return (y - l.B()) / l.A()
}
func (l *Line) Y(x int) int {
	return (l.A() * x) + l.B()
}

func (c *Coordinate) SetCoordinates(s string) {
	coordinate := strings.Split(s, ",")
	x, _ := strconv.Atoi(coordinate[0])
	y, _ := strconv.Atoi(coordinate[1])
	c.X = x
	c.Y = y

}
func (c *Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{X: x, Y: y}
}
func NewLine() *Line {
	return &Line{One: NewCoordinate(0, 0), Two: NewCoordinate(0, 0)}
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
func parseInput(input []string) (lines []*Line) {
	for _, rline := range input {
		line := NewLine()
		line.UpdateLine(rline)
		lines = append(lines, line)
	}

	return
}

func part1(lines []*Line, boundary int) int {
	f := make(map[int]map[int]int, boundary+1)

	for i := 0; i <= boundary; i++ {
		f[i] = make(map[int]int, boundary+1)
		for j := 0; j <= boundary; j++ {
			f[i][j] = 0
		}
	}

	for _, line := range lines {
		if line.One.X != line.Two.X && line.One.Y != line.Two.Y {
			// Diagonal line
			continue
		}
		if line.One.X == line.Two.X {
			// Vertical line
			for i := line.SortbyY().One.Y; i <= line.SortbyY().Two.Y; i++ {
				f[i][line.One.X] += 1
			}
		}
		if line.One.Y == line.Two.Y {
			// Horizontal line
			for i := line.SortbyX().One.X; i <= line.SortbyX().Two.X; i++ {
				f[line.One.Y][i] += 1
			}
		}
	}
	count := 0
	for i := 0; i <= boundary; i++ {
		for j := 0; j <= boundary; j++ {
			if f[i][j] >= 2 {
				count += 1
			}
			/*
				if f[i][j] == 0 {
					fmt.Printf("%s", ".")

				} else {
					fmt.Printf("%d", f[i][j])
				}
			*/
		}
		//fmt.Printf("\n")
	}
	return count

}

func part2(lines []*Line, boundary int) int {
	f := make(map[int]map[int]int, boundary+1)

	for i := 0; i <= boundary; i++ {
		f[i] = make(map[int]int, boundary+1)
		for j := 0; j <= boundary; j++ {
			f[i][j] = 0
		}
	}

	for _, line := range lines {
		if line.One.X != line.Two.X && line.One.Y != line.Two.Y {
			// Diagonal line
			for i := line.SortbyX().One.X; i <= line.SortbyX().Two.X; i++ {
				j := line.Y(i)
				f[j][i] += 1
			}
		}
		if line.One.X == line.Two.X {
			// Vertical line
			for i := line.SortbyY().One.Y; i <= line.SortbyY().Two.Y; i++ {
				f[i][line.One.X] += 1
			}
		}
		if line.One.Y == line.Two.Y {
			// Horizontal line
			for i := line.SortbyX().One.X; i <= line.SortbyX().Two.X; i++ {
				f[line.One.Y][i] += 1
			}
		}
	}
	count := 0
	for i := 0; i <= boundary; i++ {
		for j := 0; j <= boundary; j++ {
			if f[i][j] >= 2 {
				count += 1
			}
			/*
				if f[i][j] == 0 {
					fmt.Printf("%s", ".")

				} else {
					fmt.Printf("%d", f[i][j])
				}
			*/
		}
		//fmt.Printf("\n")
	}
	return count
}
func main() {
	filename := os.Args[1]
	lines := parseInput(readInput(filename))
	var boundary int
	if filename == "testinput" {
		boundary = 9
	} else {
		boundary = 999
	}
	p1 := part1(lines, boundary)
	fmt.Printf("Part 1 result: %d\n", p1)
	p2 := part2(lines, boundary)
	fmt.Printf("Part 2 result: %d\n", p2)

}
