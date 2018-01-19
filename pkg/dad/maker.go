package dad

type Character struct {
	Strength int
}

type Maker struct {
	Character Character
}

func (maker *Maker) StrengthFrom(dice Dice) {
	var rolls []int
	sum := 0
	min := 6

	for i := 0; i < 4; i++ {
		rolls = append(rolls, dice.Roll())
		roll := rolls[i]
		sum += roll
		if min > roll {
			min = roll
		}
	}

	maker.Character = Character{sum - min}
}
