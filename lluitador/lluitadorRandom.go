package lluitador

import "math/rand"
import "github.com/utrescu/CombatCavallers.Go/cops"

const DefaultName = "LluitadorRandom"
// LluitadorRandom es un lluitador que ataca aleatòriament
type LluitadorRandom struct {
	llocs []cops.LlocOnPicar
	nom string
}

func CreateLluitadorRandom(nomLluitador string, onPicar cops.ILlocOnPicar) LluitadorRandom {
	nomReal := nomLluitador
	if (len(nomReal) == 0) {
		nomReal = DefaultName
	}

	lluitador := LluitadorRandom{nom: nomReal, llocs: onPicar.GetLlocsOnPicar()}
	return lluitador
}

// GetNom retorna el nom del lluitador
func (ll LluitadorRandom) GetNom() string {
	return ll.nom
}

// Protegeix retorna els llocs en que es protegeix
func (ll LluitadorRandom) Protegeix() []cops.LlocOnPicar {
	nous := make([]cops.LlocOnPicar, len(ll.llocs)-1, len(ll.llocs)-1)
	eliminar := ll.Pica()
	for _, element := range ll.llocs {
		if element != eliminar {
			nous = append(nous, element)
		}
	}
	return nous
}

// Pica retorna on pica
func (ll LluitadorRandom) Pica() cops.LlocOnPicar {
	return ll.llocs[rand.Intn(len(ll.llocs))]
}
