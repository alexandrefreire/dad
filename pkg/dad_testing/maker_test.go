package dad_testing

import (
	"github.com/alexandrefreire/dad/pkg/dad"
	"math/rand"
	"sort"
	"testing"
)

func TestRandomToGetRolling(t *testing.T) {
	//predict result based on fixed seed
	const count = 4
	const seed = 42

	src := rand.NewSource(seed)
	rnd := rand.New(src)

	var randoms []int
	for i := 0; i < count; i++ {
		randoms = append(randoms, rnd.Intn(6)+1)
	}
	sort.Ints(randoms)

	sumOfBestThree := randoms[1] + randoms[2] + randoms[3]

	// use fixed seed in fresh maker run
	dice := dad.Dice{}
	dice.Init(seed)

	maker := dad.Maker{}
	maker.StrengthFrom(&dice)


	if sumOfBestThree != maker.Character.Strength {
		t.Fail()
	}
}

type FakeDice struct {
}

func (fd *FakeDice) Roll() int {
	return 1
}

func TestFourOnesIsThree(t *testing.T) {
	fd := FakeDice{}

	// Use this to prove that 4 1's yields a result of 3.
	maker := dad.Maker{}
	maker.StrengthFrom(&fd)

	if 3 != maker.Character.Strength {
		t.Fail()
	}
}

func TestFourSixesIsEighteen(t *testing.T) {
	// Use this to prove that 4 6's yields a result of 18.
}

func TestIgnoresLowestRoll(t *testing.T) {
	// try doing [1,6,6,6], [2,6,6,6],[3,6,6,6], etc., results always 18
}

func TestPositionOfMinDoesNotMatter(t *testing.T) {
	// here, let's do [1,6,6,6], [6,1,6,6], [6,6,1,6], and [6,6,6,1]
}
