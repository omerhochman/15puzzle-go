package game

import "testing"

func TestIsArraySolvableYesOdd(t *testing.T) {
	solvable := IsArraySolvable([]int{1, 8, 2, 0, 4, 3, 7, 6, 5}, 3)

	if !solvable {
		t.Error("Array should be solvable")
	}
}
func TestIsArraySolvableYesEven1(t *testing.T) {
	solvable := IsArraySolvable([]int{6, 13, 7, 10, 8, 9, 11, 0, 15, 2, 12, 5, 14, 3, 1, 4}, 4)

	if !solvable {
		t.Error("Array should be solvable")
	}
}
func TestIsArraySolvableYesEven2(t *testing.T) {
	solvable := IsArraySolvable([]int{13, 2, 10, 3, 1, 12, 8, 4, 5, 0, 9, 6, 15, 14, 11, 7}, 4)

	if !solvable {
		t.Error("Array should be solvable")
	}
}
func TestIsArraySolvableNoEven(t *testing.T) {
	solvable := IsArraySolvable([]int{3, 9, 1, 15, 14, 11, 4, 6, 13, 0, 10, 12, 2, 7, 8, 5}, 4)

	if solvable {
		t.Error("Array should not be solvable")
	}
}
