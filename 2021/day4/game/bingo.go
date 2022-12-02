package game

import (
	"fmt"
)

const (
	//CellsPerSide is the amount of rows and cells per row
	CellsPerSide = 5
)

// Bingo is the main game
type Bingo struct {
	Boards  []*Board
	IsBingo bool
}

// Initialize inits boards and seeds with numbers
func (b *Bingo) Initialize(boardNumbers []int) {
	boardCount := len(boardNumbers) / (CellsPerSide * CellsPerSide)

	// Initialize boards
	for i := 0; i < boardCount; i++ {
		board := &Board{}
		board.Seed(boardNumbers[i*25 : ((i + 1) * 25)])
		b.Boards = append(b.Boards, board)
	}
}

// RunGame rolls new numbers and checks boards for bingo
func (b *Bingo) RunGame(drawNumbers []int) {
	isBingo := false
	count := 0
	// Contains drawn numbers
	usedNumbers := make([]int, 0, CellsPerSide*CellsPerSide)

	// Draw numbers until a board gets bingo
	for isBingo == false {
		number := drawNumbers[count]
		usedNumbers = append(usedNumbers, number)

		for _, board := range b.Boards {
			board.MarkNumberAsChecked(number)

			if board.IsBingo() {
				isBingo = true
			}
		}
		count += 1
	}
	sumBingoBoard := 0
	lastNumber := 0

	for _, board := range b.Boards {

		if board.hasBingo {
			board.PrintBoard()
			sumBingoBoard = board.Sum()
		}
	}

	fmt.Println()
	fmt.Println("Drawn numbers: ")
	for _, number := range usedNumbers {
		fmt.Print(number, " ")
	}
	fmt.Println()

	lastNumber = usedNumbers[len(usedNumbers)-1]
	fmt.Println("Puzzle output:")
	fmt.Printf("%d is the sum of the bingo board, %d is the last number drawn\n", sumBingoBoard, lastNumber)
	fmt.Printf("Puzzle Result is %d\n", sumBingoBoard*lastNumber)

}

// RunGame rolls new numbers and checks boards for bingo
func (b *Bingo) RunGame2(drawNumbers []int) {
	isBingo := false
	count := 0
	bingoCount := 0
	lastNumber := 0

	// Contains drawn numbers
	var usedNumbers []int
	// Draw numbers until a board gets bingo
	for isBingo == false {
		number := drawNumbers[count]
		usedNumbers = append(usedNumbers, number)

		for _, board := range b.Boards {
			if board.IsBingo() {
				continue
			}
			board.MarkNumberAsChecked(number)
			if board.IsBingo() {
				bingoCount += 1
			}
			if bingoCount == len(b.Boards) {
				isBingo = true
				board.PrintBoard()
				lastNumber = usedNumbers[len(usedNumbers)-1]
				fmt.Println("Puzzle output:")
				fmt.Printf("%d is the sum of the bingo board, %d is the last number drawn\n", board.Sum(), lastNumber)
				fmt.Printf("Puzzle Result is %d\n", board.Sum()*lastNumber)
			}
		}
		count += 1
	}
}
