package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

func parseInput(input []string) (numbers []uint64) {
	for _, v := range input {
		bitNum, err := strconv.ParseUint(v, 2, len(v))
		if err != nil {
			panic(err) // currently panics with "value out of range"
		}
		numbers = append(numbers, bitNum)
	}

	return numbers
}

func part1(input []uint64, bitlen uint64) (gamma, epsilon uint64) {
	numlen := uint64(len(input) / 2)
	bitCount := make(map[uint64]uint64)
	for i := uint64(0); i < bitlen; i++ {
		idx := uint64(math.Pow(float64(2), float64(i)))
		bitCount[idx] = 0
	}
	for _, number := range input {
		for key, _ := range bitCount {
			if number&key == key {
				bitCount[key] += 1
			}
		}
	}
	//fmt.Println(bitCount)
	gamma = uint64(0)

	for key, value := range bitCount {
		if value > numlen {
			gamma += key
		}
	}

	mask := uint64(math.Pow(float64(2), float64(bitlen))) - uint64(1)
	epsilon = ^gamma & mask

	return gamma, epsilon
}

func matchNum(input []uint64, match uint64) (result []uint64) {

	for _, number := range input {
		if number&match == match {
			result = append(result, number)
		}
	}
	return result
}
func filterNum(input []uint64, match uint64) (result []uint64) {

	for _, number := range input {
		if number&^match == number {
			result = append(result, number)
		}
	}
	return result
}

func calcBits(input []uint64, bitlen uint64) (bitCount, nobitCount map[uint64]uint64) {
	bitCount = make(map[uint64]uint64)
	nobitCount = make(map[uint64]uint64)
	for i := uint64(0); i < bitlen; i++ {
		idx := uint64(math.Pow(float64(2), float64(i)))
		bitCount[idx] = 0
		nobitCount[idx] = 0
	}
	for _, number := range input {
		for key, _ := range bitCount {
			if number&key == key {
				bitCount[key] += 1
			} else {
				nobitCount[key] += 1
			}
		}
	}
	return
}
func part2(input []uint64, bitlen uint64) (oxygen, co2 uint64) {
	var matchCount, nomatchCount map[uint64]uint64
	matchCount, nomatchCount = calcBits(input, bitlen)

	oxyNumbers := input
	co2Numbers := input

	for i := uint64(bitlen); i > 0; i-- {
		idx := uint64(math.Pow(float64(2), float64(i-1)))

		if len(oxyNumbers) == 1 {
			continue
		}
		if i == uint64(1) {
			oxyNumbers = matchNum(oxyNumbers, idx)
			matchCount, nomatchCount = calcBits(oxyNumbers, bitlen)
			continue
		}
		if matchCount[idx] >= nomatchCount[idx] {
			oxyNumbers = matchNum(oxyNumbers, idx)
			matchCount, nomatchCount = calcBits(oxyNumbers, bitlen)
		}
		if matchCount[idx] < nomatchCount[idx] {
			oxyNumbers = filterNum(oxyNumbers, idx)
			matchCount, nomatchCount = calcBits(oxyNumbers, bitlen)
		}
	}
	for i := uint64(bitlen); i > 0; i-- {
		idx := uint64(math.Pow(float64(2), float64(i-1)))
		matchCount, nomatchCount = calcBits(co2Numbers, bitlen)

		if len(co2Numbers) == 1 {
			continue
		}
		if i == uint64(1) {
			co2Numbers = matchNum(co2Numbers, idx)
			continue
		}
		if matchCount[idx] < nomatchCount[idx] {
			co2Numbers = matchNum(co2Numbers, idx)
		}
		if matchCount[idx] >= nomatchCount[idx] {
			co2Numbers = filterNum(co2Numbers, idx)
		}
	}

	return oxyNumbers[0], co2Numbers[0]
}

func testmatchNum(match uint64) (result bool) {

	number := uint64(836)

	if number&match == match {
		result = true
	}
	return result
}

func main() {
	strings := readInput("input")
	input := parseInput(strings)
	gamma, epsilon := part1(input, uint64(len(strings[0])))
	oxygen, co2 := part2(input, uint64(len(strings[0])))
	fmt.Printf("Part 1 - gamma: %d, epsilon: %d ; result = %d\n", gamma, epsilon, gamma*epsilon)
	fmt.Printf("Part 2 - oxygen: %d, co2: %d ; result = %d\n", oxygen, co2, oxygen*co2)

}
