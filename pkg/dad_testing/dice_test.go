package dad_testing

import (
	"github.com/alexandrefreire/dad/pkg/dad"
	"math/rand"
	"testing"
)

func TestCorrectCounts(t *testing.T) {
	dice := &dad.Dice{}
	dice.Init(42)

	rolls := dice.Rolls(0)
	if len(rolls) != 0 {
		t.Fail()
	}

	rolls = dice.Rolls(2)
	if len(rolls) != 2 {
		t.Fail()
	}
}

func TestCorrectSequence(t *testing.T) {
	const count = 10
	const seed = 13
	var randoms []int

	src := rand.NewSource(seed)
	rnd := rand.New(src)

	for i := 0; i < count; i++ {
		randoms = append(randoms, rnd.Intn(6)+1)
	}

	dice := &dad.Dice{}
	dice.Init(seed)
	rolls := dice.Rolls(count)

	for i := 0; i < count; i++ {
		if rolls[i] != randoms[i] {
			t.Fail()
		}
	}
}
