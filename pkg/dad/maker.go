package dad

type Character struct {
	Strength int
}

type Maker struct {
	Character Character
}

func (maker *Maker) StrengthFrom(dice Dice) {
	var rolls []int
	for i := 0; i < 4; i++ {
		rolls = append(rolls, dice.Roll())
	}

	strength := best3OutOf4(rolls)

	maker.Character = Character{strength}
}

func best3OutOf4(rolls []int) int {
	sum := 0
	min := 6
	for i := 0; i < 4; i++ {
		roll := rolls[i]
		sum += roll
		if min > roll {
			min = roll
		}
	}
	strength := sum - min
	return strength
}
