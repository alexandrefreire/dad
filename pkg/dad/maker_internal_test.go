package dad

import (
	"testing"
)

func TestFourOnesIsThree(t *testing.T) {
	if best3OutOf4([]int{1,1,1,1}) != 3 {
		t.Fail()
	}
}

func TestFourSixesIsEighteen(t *testing.T) {
	if best3OutOf4([]int{6,6,6,6}) != 18 {
		t.Fail()
	}
}

func TestIgnoresLowestRoll(t *testing.T) {
	// try doing [1,6,6,6], [2,6,6,6],[3,6,6,6], etc., results always 18
}

func TestPositionOfMinDoesNotMatter(t *testing.T) {
	// here, let's do [1,6,6,6], [6,1,6,6], [6,6,1,6], and [6,6,6,1]
}
