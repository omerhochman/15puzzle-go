package game

import (
	"math/rand"
	"time"
)

func CreateRandomBoard(boardSize int) [][]int {
	numbersArray, k := createSolvableNumbersArray(boardSize), 0
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
		for j := range board[i] {
			board[i][j] = numbersArray[k]
			k++
		}
	}
	return board
}

func createSolvableNumbersArray(boardSize int) []int {
	var array []int
	for !IsArraySolvable(array, boardSize) {
		array = createShuffledNumbersArray(boardSize)
	}

	return array
}

func createShuffledNumbersArray(boardSize int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	arr := make([]int, boardSize*boardSize)

	for i, num := range r.Perm(len(arr)) {
		arr[i] = num
	}

	return arr
}

func IsArraySolvable(array []int, boardSize int) bool {
	if array == nil {
		return false
	}

	inversionsCount := getInversionsCount(array)

	if boardSize%2 == 1 {
		return inversionsCount%2 == 0
	}

	zeroPosition := getZeroDistanceFromBottom(array, boardSize)

	if zeroPosition%2 == 1 {
		return inversionsCount%2 == 0
	}

	return inversionsCount%2 == 1
}

func getZeroDistanceFromBottom(array []int, boardSize int) int {
	zeroIndex := 0

	for i, val := range array {
		if val == 0 {
			zeroIndex = i
		}
	}

	return boardSize - (zeroIndex / boardSize)
}

func getInversionsCount(array []int) int {
	inversionsCount := 0

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > 0 && array[j] > 0 && array[i] > array[j] {
				inversionsCount++
			}
		}
	}

	return inversionsCount
}
