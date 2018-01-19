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
	// try doing [1,6,6,6], [2,6,6,6],[3,6,6,6], etc., results always 18
}

func TestPositionOfMinDoesNotMatter(t *testing.T) {
	// here, let's do [1,6,6,6], [6,1,6,6], [6,6,1,6], and [6,6,6,1]
}
