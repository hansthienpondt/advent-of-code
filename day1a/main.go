package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	TotalMeasurements int
	SonarIncreased    int
	SonarDecreased    int
)

type Sonar struct {
	current  int
	previous int
}

func (s *Sonar) NewMeasurement(i int) (bool, error) {
	s.previous = s.current
	s.current = i

	if s.current > s.previous {
		return true, nil
	}
	if s.current < s.previous {
		return false, nil
	}
	return false, fmt.Errorf("incorrect input!")
}

func NewSonar() *Sonar { return &Sonar{current: 0, previous: 0} }

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	s := NewSonar()
	TotalMeasurements = 0
	SonarIncreased = 0
	SonarDecreased = 0

	// fmt.Println("There are a total of " + lines.len() + " numbers")
	for k, v := range lines {
		measurement, _ := strconv.Atoi(v)
		increased, _ := s.NewMeasurement(measurement)
		if err != nil {
			fmt.Println("ERROR!!")
		}
		if increased && k != 0 {
			SonarIncreased++
		}
		if !increased {
			SonarDecreased++
		}
		TotalMeasurements++
	}
	// Printing the Stdout seperator
	fmt.Println(strings.Repeat("#", 64))
	fmt.Printf("There are %d total measurements\n", TotalMeasurements)
	fmt.Printf("%d measurements were increased vs the previous value\n", SonarIncreased)
	fmt.Printf("%d measurements were decreased vs the previous value\n", SonarDecreased)
}
