package combat

import "github.com/utrescu/CombatCavallers.Go/cops"

type CopsPermesos struct {}

func (c CopsPermesos) GetLlocsOnPicar() []cops.LlocOnPicar {
	return []cops.LlocOnPicar{
		cops.Cap,
		cops.CostatEsquerra,
		cops.CostatDret,
		cops.Panxa,
	};
}
