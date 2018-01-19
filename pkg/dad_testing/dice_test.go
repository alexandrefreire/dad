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
