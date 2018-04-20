/* ***************************************************************************
# Copyright (c) 2018, Industrial Logic, Inc., All Rights Reserved.
#
# This code is the exclusive property of Industrial Logic, Inc. It may ONLY be
# used by students during Industrial Logic's workshops or by individuals
# who are being coached by Industrial Logic on a project.
#
# This code may NOT be copied or used for any other purpose without the prior
# written consent of Industrial Logic, Inc.
# ****************************************************************************/
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

func (d Dice) Roll() int {
	return d.Rnd.Intn(6) + 1
}

func (d Dice) Rolls(count int) []int {
	rolls := make([]int, count)
	for i := range rolls {
		rolls[i] = d.Roll()
	}
	return rolls
}
