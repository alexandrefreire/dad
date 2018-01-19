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
	rolls []int
	i     int
}

func (fd *FakeDice) Roll() int {
	defer func() {fd.i++}()
	return fd.rolls[fd.i]
}

func checkMakerFor(t *testing.T, result int, rs ...int) {
	fd := FakeDice{rolls: rs}
	maker := dad.Maker{}
	maker.StrengthFrom(&fd)
	if result != maker.Character.Strength {
		t.Fail()
	}
}

func TestFourOnesIsThree(t *testing.T) {
	checkMakerFor(t, 3, 1,1,1,1)
}

func TestFourSixesIsEighteen(t *testing.T) {
	checkMakerFor(t, 18, 6,6,6,6)
}


func TestIgnoresLowestRoll(t *testing.T) {
	checkMakerFor(t, 18, 1,6,6,6)
	checkMakerFor(t, 18, 2,6,6,6)
	checkMakerFor(t, 18, 3,6,6,6)
	checkMakerFor(t, 18, 4,6,6,6)
	checkMakerFor(t, 18, 5,6,6,6)
}

func TestPositionOfMinDoesNotMatter(t *testing.T) {
	checkMakerFor(t, 18, 6,1,6,6)
	checkMakerFor(t, 18, 6,6,1,6)
	checkMakerFor(t, 18, 6,6,6,1)
}
