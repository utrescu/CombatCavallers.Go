package main

import (
	"math/rand"
	"time"

	"github.com/utrescu/CombatCavallers.Go/combat"
	"github.com/utrescu/CombatCavallers.Go/lluitador"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	lluitador1 := lluitador.CreateLluitadorRandom("AixafaMandonguilles")
	lluitador2 := lluitador.CreateLluitadorRandom("MataCerilles")

	ring := combat.Ring{}

	ring.EntrenJugadors(lluitador1, lluitador2)
	ring.Lluiteu()

}
