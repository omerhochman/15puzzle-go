package game

import "testing"

func TestDirectionFromString_valid(t *testing.T) {
	assertDirectionFromStringValid("Up", UP, t)
	assertDirectionFromStringValid("UP", UP, t)
	assertDirectionFromStringValid("up", UP, t)
	assertDirectionFromStringValid("DOwn", DOWN, t)
	assertDirectionFromStringValid("down", DOWN, t)
	assertDirectionFromStringValid("left", LEFT, t)
	assertDirectionFromStringValid("Left", LEFT, t)
	assertDirectionFromStringValid("Right", RIGHT, t)
	assertDirectionFromStringValid("rigHt", RIGHT, t)
}

func assertDirectionFromStringValid(str string, expectedDirection Direction, t *testing.T) {
	dir := DirectionFromString(str)
	if dir != expectedDirection {
		t.Errorf("Wrong direction. Actual: %v. Expected: %v.", dir, expectedDirection)
	}
}

func TestDirectionFromString_invalid(t *testing.T) {
	assertDirectionFromStringInvalid("Upp", t)
	assertDirectionFromStringInvalid("Lleft", t)
	assertDirectionFromStringInvalid("downn", t)
	assertDirectionFromStringInvalid("dow", t)
}

func assertDirectionFromStringInvalid(str string, t *testing.T) {
	dir := DirectionFromString(str)
	if dir != "" {
		t.Errorf("Wrong direction. Actual: %v. Expected: %v.", dir, "''")
	}
}
