package dad

import (
	"math/rand"
)

type Dice struct {
	Rnd *rand.Rand
}

func (d *Dice) Init(seed int64) {
	s := rand.NewSource(seed)
	d.Rnd = rand.New(s)
}

func (d *Dice) Roll() int {
	return d.Rnd.Intn(6) + 1
}

func (d *Dice) Rolls(count int) []int {
	rolls := make([]int, count)
	for i := range rolls {
		rolls[i] = d.Roll()
	}
	return rolls
}
