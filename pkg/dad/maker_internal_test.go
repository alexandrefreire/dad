package dad

import (
	"testing"
)

func checkBest3OutOf4(ints []int, strength int, t *testing.T) {
	if best3OutOf4(ints) != strength {
		t.Fail()
	}
}

func TestFourOnesIsThree(t *testing.T) {
	ints := []int{1, 1, 1, 1}
	strength := 3
	checkBest3OutOf4(ints, strength, t)
}

func TestFourSixesIsEighteen(t *testing.T) {
	checkBest3OutOf4([]int{6,6,6,6}, 18, t)
}

func TestIgnoresLowestRoll(t *testing.T) {
	// try doing [1,6,6,6], [2,6,6,6],[3,6,6,6], etc., results always 18
}

func TestPositionOfMinDoesNotMatter(t *testing.T) {
	// here, let's do [1,6,6,6], [6,1,6,6], [6,6,1,6], and [6,6,6,1]
}
