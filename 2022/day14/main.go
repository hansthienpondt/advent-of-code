package main

import (
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var filename = flag.String("input", "input", "filename that serves as program input")
var sand = image.Point{500, 0}

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	return lines
}
func parseInput(input []string) (rock map[image.Point]rune, min, max image.Point) {
	rock = make(map[image.Point]rune)
	min.X = math.MaxInt
	min.Y = math.MaxInt
	max.X = 0
	max.Y = 0

	for _, l := range input {
		s := strings.Split(l, " -> ")

		for p := 0; p < len(s)-1; p++ {
			var point1, point2 image.Point

			fmt.Sscanf(s[p], "%d,%d", &point1.X, &point1.Y)
			fmt.Sscanf(s[p+1], "%d,%d", &point2.X, &point2.Y)

			log.Debugf("parsing point pair: %v %v", point1, point2)

			for drawLine := (image.Point{sign(point2.X - point1.X), sign(point2.Y - point1.Y)}); point1 != point2.Add(drawLine); point1 = point1.Add(drawLine) {
				// rune # = 35 ; rune + = 43 ; rune o = 111 ; rune . = 46 ; rune ~ = 126
				rock[point1] = '#'
				if point1.Y > max.Y {
					max.Y = point1.Y
				}
				if point1.X > max.X {
					max.X = point1.X
				}
				if point1.X < min.X {
					min.X = point1.X
				}
				if point1.X < min.Y {
					min.Y = point1.Y
				}
			}
		}
	}
	rock[sand] = '+'
	return
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
func draw(min, max image.Point, rock map[image.Point]rune) {
	for i := 0; i < max.Y+2; i++ {
		for j := min.X - 1; j < max.X+1; j++ {
			p := image.Point{j, i}
			if v, ok := rock[p]; ok {
				fmt.Printf("%s", string(v))
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func dropSandA(rock map[image.Point]rune, sand, min, max image.Point) (map[image.Point]rune, bool) {
	moves := []image.Point{
		{0, 1},
		{-1, 1},
		{1, 1}}

	rest := false
	var position = sand

outer:
	for !rest {
		for _, move := range moves {
			if _, ok := rock[position.Add(move)]; !ok {
				position = position.Add(move)
				if position.Y > max.Y {
					return rock, true
				}
				continue outer
			}
		}
		rest = true
		rock[position] = 'o'
	}

	return rock, false
}
func dropSandB(rock map[image.Point]rune, sand, min, max image.Point) (map[image.Point]rune, bool) {
	moves := []image.Point{
		{0, 1},
		{-1, 1},
		{1, 1}}

	rest := false
	var position = sand

outer:
	for !rest {
		for _, move := range moves {
			if _, ok := rock[position.Add(move)]; !ok && position.Y < max.Y+1 {
				position = position.Add(move)
				continue outer
			}
		}
		rest = true
		rock[position] = 'o'
		if position.Y == sand.Y {
			return rock, true
		}
	}

	return rock, false
}

func day14a(rock map[image.Point]rune, min, max image.Point) int {

	fallIntoVoid := false
	for !fallIntoVoid {
		rock, fallIntoVoid = dropSandA(rock, sand, min, max)
	}

	draw(min, max, rock)

	var count int = 0

	for _, r := range rock {
		// rune o = 111
		if r == 'o' {
			count++
		}
	}

	return count
}

func day14b(rock map[image.Point]rune, min, max image.Point) int {
	reachedTop := false
	var count int = 0

	for !reachedTop {
		rock, reachedTop = dropSandB(rock, sand, min, max)
	}
	min.X -= 5
	max.X += 9
	max.Y += 1
	draw(min, max, rock)

	count = 0

	for _, r := range rock {
		// rune o = 111
		if r == 'o' {
			count++
		}
	}

	return count
}
func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.DebugLevel)
	log.SetLevel(log.WarnLevel)
	// Set Calling method to true
	log.SetReportCaller(true)

}

func main() {
	flag.Parse()

	rock, min, max := parseInput(readInput(*filename))

	fmt.Printf("Part 1: %d\n", day14a(rock, min, max))
	fmt.Printf("Part 2: %d\n", day14b(rock, min, max))
}
