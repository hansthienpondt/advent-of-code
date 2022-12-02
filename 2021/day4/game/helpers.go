package game

import (
	"math/rand"
	"time"
)

// IsIntInSlice checks wether an int is in the slice
func IsIntInSlice(a int, list *[]int) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}

// Roll creates a random integer in the span of squares on the board, never zero
func Roll(list *[]int) int {
	var random int

	for {
		rand.Seed(int64(time.Now().Nanosecond()))
		random = rand.Intn(CellsPerSide*CellsPerSide + 1)

		if IsIntInSlice(random, list) || random == 0 {
			continue
		}

		*list = append(*list, random)
		break
	}

	return random
}
