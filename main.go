package main

import (
	"math/rand"
	"time"

	"github.com/utrescu/CombatCavallers.Go/combat"
	"github.com/utrescu/CombatCavallers.Go/lluitador"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	copsPermesos := combat.CopsPermesos{}

	lluitador1 := lluitador.CreateLluitadorRandom("AixafaMandonguilles", copsPermesos)
	lluitador2 := lluitador.CreateLluitadorRandom("MataCerilles", copsPermesos)

	ring := combat.Ring{}

	ring.EntrenJugadors(lluitador1, lluitador2)
	ring.Lluiteu()

}
