package lluitador

import "math/rand"

type LlocOnPicar string

const (
	Cap            LlocOnPicar = "al cap"
	CostatEsquerra LlocOnPicar = "al costat esquerra"
	CostatDret     LlocOnPicar = "al costat dret"
	Panxa          LlocOnPicar = "a la panxa"
)

type LlocsOnPicar struct {
	llocs []LlocOnPicar
}

func CreateLlocsOnPicar() LlocsOnPicar {
	llocsOnPicar := LlocsOnPicar{
		llocs: []LlocOnPicar{
			Cap,
			CostatEsquerra,
			CostatDret,
			Panxa,
		},
	}
	return llocsOnPicar
}

func (c LlocsOnPicar) GetRandomLloc() LlocOnPicar {
	return c.llocs[rand.Intn(len(c.llocs))]
}

func (c LlocsOnPicar) GetWithoutOne() []LlocOnPicar {
	nous := make([]LlocOnPicar, len(c.llocs)-1, len(c.llocs)-1)
	eliminar := c.GetRandomLloc()
	for _, element := range c.llocs {
		if element != eliminar {
			nous = append(nous, element)
		}
	}
	return nous
}
