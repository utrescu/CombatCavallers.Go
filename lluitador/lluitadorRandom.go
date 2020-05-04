package lluitador

// LluitadorRandom es un lluitador que ataca aleatòriament
type LluitadorRandom struct {
	nom string
}

func CreateLluitadorRandom(nomLluitador string) LluitadorRandom {
	lluitador := LluitadorRandom{nom: nomLluitador}
	return lluitador
}

// GetNom retorna el nom del lluitador
func (ll LluitadorRandom) GetNom() string {
	return ll.nom
}

// Protegeix retorna els llocs en que es protegeix
func (ll LluitadorRandom) Protegeix() []LlocOnPicar {
	var llocs = CreateLlocsOnPicar()
	return llocs.GetWithoutOne()
}

// Pica retorna on pica
func Pica() LlocOnPicar {
	var llocs = CreateLlocsOnPicar()
	return llocs.GetRandomLloc()
}
