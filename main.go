package main

import (
	"github.com/utrescu/CombatCavallers.Go/combat"
	"github.com/utrescu/CombatCavallers.Go/lluitador"
)

func main() {
	lluitador1 := lluitador.CreateLluitadorRandom("AixafaMandonguilles")
	lluitador2 := lluitador.CreateLluitadorRandom("MataCerilles")

	ring := combat.Ring()

	ring.EntrenJugadors(lluitador1, lluitador2)
	ring.Lluiteu()

}
