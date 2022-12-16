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
var height = flag.Int("height", 2000000, "specify height to match on")
var yMax = flag.Int("yMax", 4000000, "specify yMax for tuning freq")

var sand = image.Point{500, 0}

type sensorBeacon struct {
	sensor image.Point
	beacon image.Point
}

func (sb *sensorBeacon) Distance() int {

	distance := abs(sb.sensor.X-sb.beacon.X) + abs(sb.sensor.Y-sb.beacon.Y)
	return distance
}

func (sb *sensorBeacon) isSensor(p image.Point) bool {
	if p == sb.sensor && p != sb.beacon {
		return true
	}
	return false
}
func (sb *sensorBeacon) isBeacon(p image.Point) bool {
	if p == sb.beacon && p != sb.sensor {
		return true
	}
	return false
}

func NewSensorBeacon() *sensorBeacon {
	return &sensorBeacon{sensor: image.Point{}, beacon: image.Point{}}
}

type grid struct {
	min    image.Point
	max    image.Point
	m      map[image.Point]*sensorBeacon
	sbList []*sensorBeacon
}

func (g *grid) Sensors() (l []*sensorBeacon) {
	for item, sb := range g.m {
		if sb.isSensor(item) {
			l = append(l, sb)
		}
	}
	return
}
func (g *grid) Draw() string {
	var s string = ""
	for i := g.min.Y - 1; i < g.max.Y+1; i++ {
		if i < 0 {
			s += fmt.Sprintf("   ")
		} else {
			s += fmt.Sprintf("%02d ", i)
		}
		for j := g.min.X; j < g.max.X+1; j++ {
			p := image.Point{j, i}

			if i < g.min.Y {
				if j == 0 {
					s += fmt.Sprintf("%d", j)
				} else if j%5 == 0 {
					s += fmt.Sprintf("%d", 5)
				} else {
					s += fmt.Sprintf(" ")
				}
			} else if v, ok := g.m[p]; ok && v.isSensor(p) {
				s += fmt.Sprintf("S")
			} else if v, ok := g.m[p]; ok && v.isBeacon(p) {
				s += fmt.Sprintf("B")
			} else {
				s += fmt.Sprintf(".")
			}
		}
		s += fmt.Sprintf("\n")
	}
	return s
}

func NewGrid() *grid {
	return &grid{
		min:    image.Point{math.MaxInt, math.MaxInt},
		max:    image.Point{0, 0},
		m:      make(map[image.Point]*sensorBeacon, 0),
		sbList: make([]*sensorBeacon, 0)}

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
func parseInput(input []string) *grid {

	g := NewGrid()

	for _, l := range input {
		sb := NewSensorBeacon()
		var sensor, beacon image.Point
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)
		sb.sensor = sensor
		sb.beacon = beacon
		log.Debugf("Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d ; distance=%d\n", sb.sensor.X, sb.sensor.Y, sb.beacon.X, sb.beacon.Y, sb.Distance())

		g.m[sensor] = sb
		g.m[beacon] = sb
		g.sbList = append(g.sbList, sb)
	}

	for item, _ := range g.m {
		if item.X > g.max.X {
			g.max.X = item.X
		}
		if item.Y > g.max.Y {
			g.max.Y = item.Y
		}
		if item.X < g.min.X {
			g.min.X = item.X
		}
		if item.Y < g.min.Y {
			g.min.Y = item.Y
		}
	}
	return g
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func day15a(g *grid, h int) int {
	match := make(map[int]bool, 0)
	for _, s := range g.Sensors() {
		diff := s.Distance() - abs(h-s.sensor.Y)

		for x := s.sensor.X - diff; x <= s.sensor.X+diff; x++ {
			if !(s.beacon.X == x && s.beacon.Y == h) {
				match[x] = true
			}
		}
	}

	return len(match)
}

func day15b(g *grid, ymax int) int {
	for y := 0; y <= ymax; y++ {
	loop:
		for x := 0; x <= ymax; x++ {
			for _, s := range g.Sensors() {
				dx := s.sensor.X - x
				dy := s.sensor.Y - y
				if abs(dx)+abs(dy) <= s.Distance() {
					x += s.Distance() - abs(dy) + dx
					continue loop
				}
			}
			return (x*4000000 + y)
		}
	}

	//var count int = 0

	return 0
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

	g := parseInput(readInput(*filename))

	fmt.Printf("Part 1: %d\n", day15a(g, *height))
	fmt.Printf("Part 2: %d\n", day15b(g, *yMax))
}
