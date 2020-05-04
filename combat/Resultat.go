package combat

import "github.com/utrescu/CombatCavallers.Go/lluitador"

// IResultat retorna el resultat del combat
type IResultat interface {
	GetNom() string
	GetVida() int
	EsKo() bool
}

// ICombatent definex el que necessita el Ring
type ICombatent interface {
	GetLluitador() lluitador.ILluitador
	TreuVida() int
}

// Resultat és la implementació de les dues interfícies
type Resultat struct {
	lluitador lluitador.ILluitador
	vida      int
}

// CreateResultat crea un objecte de tipus resultat
func CreateResultat(lluitador1 lluitador.ILluitador, vida int) Resultat {
	return Resultat{lluitador1, vida}
}

func (r Resultat) GetNom() string {
	return r.lluitador.GetNom()
}

func (r Resultat) GetVida() int {
	return r.GetVida()
}

func (r Resultat) EsKo() bool {
	return r.GetVida() == 0
}

// GetLluitador retorna el lluitador
func (r Resultat) GetLluitador() lluitador.ILluitador {
	return r.lluitador
}

func (r *Resultat) TreuVida() int {
	r.vida--
	return r.vida
}
