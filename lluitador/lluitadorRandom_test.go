package lluitador

import "testing"
import "github.com/utrescu/CombatCavallers.Go/cops"

type MockCops struct {}

func (c MockCops) GetLlocsOnPicar() []cops.LlocOnPicar {
	return []cops.LlocOnPicar{
		cops.Cap,
	}
}


func TestGetNomRetornaElQueEsPosaAlCreate(t *testing.T) {
	expected := "X11"
	guerrer := CreateLluitadorRandom(expected, MockCops{})

	result := guerrer.GetNom();

	if result != expected {
		t.Errorf("El nom no coincideix, rebut: %s esperava: %s", result, expected)
	}
}

func TestGetNomSiEsBuitElNomEsDefaultName(t *testing.T) {
	expected := DefaultName
	guerrer := CreateLluitadorRandom("", MockCops{})

	result := guerrer.GetNom();

	if result != expected {
		t.Errorf("El nom no coincideix, rebut: %s esperava: %s", result, expected)
	}
}

func TestPicaRetornaElValorAletoriDeUn(t *testing.T) {
	expected := cops.Cap
	guerrer := CreateLluitadorRandom("Nom", MockCops{})

	result := guerrer.Pica()

	if result != expected {
		t.Errorf("El cop ha de ser a %s i és a %s", expected, result)
	}
}

func TestProtegeixRetornaLlistaBuidaSiNomesTeUnValor(t *testing.T) {
	expected := []cops.LlocOnPicar{}

	guerrer := CreateLluitadorRandom("Nom", MockCops{})

	result := guerrer.Protegeix()

	if len(result) != len(expected) {
		t.Errorf("El cop ha de ser a %s i és a %s", expected, result)
	}
}