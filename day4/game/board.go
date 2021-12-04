package game

import (
	"fmt"
)

// Board consists of rows and cells, filled with randomizes numbers
type Board struct {
	rowCells [CellsPerSide][CellsPerSide]int
	hasBingo bool
}

// Seed creates random numbers to all cells in the board
func (b *Board) Seed(boardNumbers []int) {
	//usedNumbers := make([]int, 0, CellsPerSide*CellsPerSide)
	count := 0
	for row, cellsArray := range b.rowCells {
		// for each pair in the map, print key and value
		for cellKey := range cellsArray {
			b.rowCells[row][cellKey] = boardNumbers[count]
			count += 1
		}
	}

	b.hasBingo = false
}

// PrintBoard prints the board and tells if it has Bingo
func (b *Board) PrintBoard() {
	fmt.Println()

	if b.hasBingo {
		fmt.Println("Board got Bingo!")
	}

	for _, cellsArray := range b.rowCells {
		fmt.Println(cellsArray)
	}
}

// MarkNumberAsChecked sets the number on the board to zero
func (b *Board) MarkNumberAsChecked(number int) {
	for row, cellsArray := range b.rowCells {
		for cellKey, cellValue := range cellsArray {
			if cellValue == number {
				b.rowCells[row][cellKey] = 0
			}
		}
	}
}

func (b *Board) Sum() (sum int) {
	for row, cellsArray := range b.rowCells {
		for cellKey, _ := range cellsArray {
			sum += b.rowCells[row][cellKey]
		}
	}
	return sum
}

// IsBingo tells if the board has bingo in horizontal or diagonal line
func (b *Board) IsBingo() bool {
	diagonalsIncline := 0
	diagonalsDecline := 0

	for rowIndex, cellsArray := range b.rowCells {
		numbersInARow := 0

		for cellIndex, cellValue := range cellsArray {
			if cellValue == 0 {
				numbersInARow++
				// Decline diagonal bingo
				if rowIndex == cellIndex {
					diagonalsDecline++
				}
				/*
					// Incline diagonal bingo
					if rowIndex+cellIndex+1 == CellsPerSide {
						diagonalsIncline++
					}
				*/
				if numbersInARow == CellsPerSide || diagonalsDecline == CellsPerSide || diagonalsIncline == CellsPerSide {
					b.hasBingo = true
				}
			}
		}
	}
	// do vertical rows
	for i := 0; i < CellsPerSide; i++ {
		numbersInARow := 0

		for j := 0; j < CellsPerSide; j++ {
			if b.rowCells[j][i] == 0 {
				numbersInARow++
			}
		}
		if numbersInARow == CellsPerSide {
			b.hasBingo = true
		}
	}

	return b.hasBingo
}
