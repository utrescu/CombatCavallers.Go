package lluitador

// ILluitador és una interfície pels lluitadors
type ILluitador interface {
	GetNom() string
	Protegeix() []LlocOnPicar
	Pica() LlocOnPicar
}
