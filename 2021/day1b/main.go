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

func toInt(input []string) (numbers []int) {
	for _, v := range input {
		value, _ := strconv.Atoi(v)
		numbers = append(numbers, value)
	}
	return numbers
}

func sum(a []int) (result int) {
	for _, v := range a {
		result += v
	}
	return result
}
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

	numbers := toInt(lines)
	for k, _ := range numbers {
		measurement := sum(numbers[k : k+3])
		if k+3 > len(numbers) {
			continue
		}
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
