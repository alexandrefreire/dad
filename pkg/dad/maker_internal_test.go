package dad

import (
	"testing"
)

func checkBest3OutOf4(t *testing.T, strength int, ints ...int) {
	if best3OutOf4(ints) != strength {
		t.Fail()
	}
}

func TestFourOnesIsThree(t *testing.T) {
	checkBest3OutOf4(t,3, 1, 1, 1, 1)
}

func TestFourSixesIsEighteen(t *testing.T) {
	checkBest3OutOf4(t, 18, 6,6,6,6)
}

func TestIgnoresLowestRoll(t *testing.T) {
	checkBest3OutOf4(t, 18, 1,6,6,6)
	checkBest3OutOf4(t, 18, 2,6,6,6)
	checkBest3OutOf4(t, 18, 3,6,6,6)
	checkBest3OutOf4(t, 18, 4,6,6,6)
	checkBest3OutOf4(t, 18, 5,6,6,6)
}

func TestPositionOfMinDoesNotMatter(t *testing.T) {
	checkBest3OutOf4(t, 18, 6,1,6,6)
	checkBest3OutOf4(t, 18, 6,6,1,6)
	checkBest3OutOf4(t, 18, 6,6,6,1)
}
