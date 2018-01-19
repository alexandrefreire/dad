package dad

type Character struct {
	Strength int
}

type Maker struct {
	Character Character
}

func (maker *Maker) StrengthFrom(dice Dice) {
	sum := 0
	min := 6

	for i := 0; i < 4; i++ {
		roll := dice.Roll()
		sum += roll
		if min > roll {
			min = roll
		}
	}

	maker.Character = Character{sum - min}
}
