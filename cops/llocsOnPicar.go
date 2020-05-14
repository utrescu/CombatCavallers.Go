package cops

type LlocOnPicar string

const (
	Cap            LlocOnPicar = "al cap"
	CostatEsquerra LlocOnPicar = "al costat esquerra"
	CostatDret     LlocOnPicar = "al costat dret"
	Panxa          LlocOnPicar = "a la panxa"
)

type ILlocOnPicar interface {
	GetLlocsOnPicar() []LlocOnPicar
}