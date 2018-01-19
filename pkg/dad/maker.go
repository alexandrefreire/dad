package dad

type Character struct {
	Strength int
}

type Maker struct {
	Character Character
}

type Roller interface{
	Roll() int
}

func (maker *Maker) StrengthFrom(dice Roller) {
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
