package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-2021/internal/utils"
)

// numberTracker
type numberTracker struct {
	board int
	row   int
	col   int
}

type board struct {
	boardSize int
	matrix    [][]int
	checked   [][]bool
	hasBingo  bool
}

var numberMap map[int][]numberTracker
var boards []board

func Day4(part string) {
	if part == "example" || part == "all" {
		lines := utils.ReadInput("day-4-example.txt")
		numberMap = map[int][]numberTracker{}
		draw, matrixes := processInput(lines, numberMap)
		boards = make([]board, 0)
		setBoards(&boards, matrixes)
		number, boardNumber := drawNumbers(draw, numberMap, boards, false)
		unmarkedSum := sumUnmarkedNumbers(boards[boardNumber])
		fmt.Printf("Day 4 - example: Score: %d\n", number*unmarkedSum)
	}
	if part != "example" {
		lines := utils.ReadInput("day-4-1.txt")
		numberMap = map[int][]numberTracker{}
		draw, matrixes := processInput(lines, numberMap)
		boards = make([]board, 0)
		setBoards(&boards, matrixes)
		if part == "1" || part == "all" {
			number, boardNumber := drawNumbers(draw, numberMap, boards, false)
			unmarkedSum := sumUnmarkedNumbers(boards[boardNumber])
			fmt.Printf("Day 4 - Part 1: Score: %d\n", number*unmarkedSum)
		}
		if part == "2" || part == "all" {
			number, boardNumber := drawNumbers(draw, numberMap, boards, true)
			unmarkedSum := sumUnmarkedNumbers(boards[boardNumber])
			fmt.Printf("Day 4 - Part 2: Score: %d\n", number*unmarkedSum)
		}
	}

	// Part 1
	// number, boardNumber := drawNumbers(draw, numberMap, boards, false)
	// unmarkedSum := sumUnmarkedNumbers(boards[boardNumber])

	// Part 2
	//number, boardNumber := drawNumbers(draw, numberMap, boards, true)
	//unmarkedSum := sumUnmarkedNumbers(boards[boardNumber])
	//fmt.Printf("BINGO! Last number drawn: %d. Sum of unmarked numbers: %d. Score: %d\n", number, unmarkedSum, number*unmarkedSum)
}

func sumUnmarkedNumbers(brd board) int {
	sum := 0
	for row := 0; row < brd.boardSize; row++ {
		for col := 0; col < brd.boardSize; col++ {
			if !brd.checked[row][col] {
				sum += brd.matrix[row][col]
			}
		}
	}
	return sum
}

// drawNumbers iterate over the numbers in sequence. For each number it looks up in numberMap
// to find where the number appers: the board number, row and column. For each of those positons
// it marks the board.checked matrix as true for the board, row and column. It also verifies
// board.checked to determine if all the elements in that row or column are checked. If they are,
// it returns the last number that was drawn and the sum of unmarked numbers on that board. If
// lastBoard is true it returns the last board to win, instead of the first one
func drawNumbers(numbers []int, numberMap map[int][]numberTracker, boards []board, lastBoard bool) (int, int) {
	N := len((boards)[0].matrix)
	totalBoards := len(boards)
	boardsWithBingo := 0
	for _, number := range numbers {
		for _, position := range numberMap[number] {
			board := &boards[position.board]
			if lastBoard && board.hasBingo {
				continue
			}
			board.checked[position.row][position.col] = true
			checkedNumbersRow := 0
			for i := 0; i < N; i++ {
				if board.checked[position.row][i] {
					checkedNumbersRow++
				} else {
					break
				}
			}
			if checkedNumbersRow == N {
				if !lastBoard {
					return number, position.board
				} else {
					board.hasBingo = true
					boardsWithBingo++
					if boardsWithBingo == totalBoards {
						return number, position.board
					}
					continue
				}
			}

			checkedNumbersCol := 0
			for i := 0; i < N; i++ {
				if board.checked[i][position.col] {
					checkedNumbersCol++
				} else {
					break
				}
			}
			if checkedNumbersCol == N {
				if !lastBoard {
					return number, position.board
				} else {
					board.hasBingo = true
					boardsWithBingo++
					if boardsWithBingo == totalBoards {
						return number, position.board
					}
				}
			}
		}
	}
	return 0, 0
}

func setBoards(boards *[]board, matrixes [][][]int) {
	for _, matrix := range matrixes {
		N := len(matrix)
		checkedMatrix := make([][]bool, N)
		for idx := range checkedMatrix {
			checkedMatrix[idx] = make([]bool, N)
		}
		boardNumbers := board{
			boardSize: N,
			matrix:    matrix,
			checked:   checkedMatrix,
		}
		*boards = append(*boards, boardNumbers)
	}
}

// processInput will parse the lines from the input file and will return a slice with
// the numbers to be drawn and a slice of boards, being each board a 5x5 matrix
func processInput(input []string, numberMap map[int][]numberTracker) ([]int, [][][]int) {
	var draw []int
	var boards [][][]int
	board := [][]int{}
	lineCounter := 0  // to mark which row in the board we are populating
	boardCounter := 0 // to set the board, line and col for each number in numberMap
	for idx, line := range input {
		if idx == 0 {
			parts := strings.Split(line, ",")
			for _, part := range parts {
				if drawNumber, err := strconv.ParseInt(part, 10, 16); err == nil {
					draw = append(draw, int(drawNumber))
				}
			}
		} else {
			if len(line) == 0 {
				if len(board) > 0 {
					boards = append(boards, board)
					boardCounter++
					lineCounter = 0
					board = [][]int{}
				}
				continue
			}
			line = strings.Trim(line, " ")
			line = strings.Replace(line, "  ", " ", -1)
			parts := strings.Split(line, " ")
			board = append(board, make([]int, len(parts)))
			for idx, value := range parts {
				if number, err := strconv.ParseInt(strings.Trim(value, " "), 10, 0); err == nil {
					board[lineCounter][idx] = int(number)
					numberMap[int(number)] = append(numberMap[int(number)], numberTracker{boardCounter, lineCounter, idx})
				}
			}
			lineCounter++
		}
	}
	boards = append(boards, board)
	return draw, boards
}
