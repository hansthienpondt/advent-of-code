package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat"
)

func readInput(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)

	return data
}
func parseInput(input string) (crabs []float64) {
	for _, number := range strings.Split(input, ",") {
		n, _ := strconv.ParseFloat(number, 64)
		crabs = append(crabs, n)
	}
	return crabs
}

func dFromMedian(crabs []float64, median float64) float64 {
	var sum float64
	sum = 0
	for _, crab := range crabs {
		if median >= crab {

			sum += (median - crab)
		} else {
			sum += (crab - median)
		}
	}
	return sum
}
func Sigma(n float64) float64 {
	var result int
	result = 0

	for i := 0; i <= int(n); i++ {
		result += i
	}
	return float64(result)
}

func d2FromMedian(crabs []float64, median float64) float64 {
	var sum float64
	sum = 0
	for _, crab := range crabs {
		if median >= crab {
			sum += Sigma((median - crab))

		} else {
			sum += Sigma((crab - median))
		}
	}
	return sum
}

func part1(crabs []float64) (median, fuelSpent float64) {

	sort.Float64s(crabs)
	median = stat.Quantile(0.5, stat.Empirical, crabs, nil)

	fuelSpent = dFromMedian(crabs, median)
	return median, fuelSpent
}

func part2(crabs []float64) (roundedmean, fuelSpent float64) {
	mean := stat.Mean(crabs, nil)
	//roundedmean = math.Round(mean)
	roundedmean = float64(int(mean))

	fuelSpent = d2FromMedian(crabs, roundedmean)
	return roundedmean, fuelSpent
}
func main() {
	filename := os.Args[1]
	crabs := parseInput(readInput(filename))
	p1median, p1fuelSpent := part1(crabs)

	fmt.Printf("Part1: the median is %d, fuel spent is %d \n", int(p1median), int(p1fuelSpent))
	p2median, p2fuelSpent := part2(crabs)
	fmt.Printf("Part2: the mean (rounded) is %d, fuel spent is %d \n", int(p2median), int(p2fuelSpent))

}
