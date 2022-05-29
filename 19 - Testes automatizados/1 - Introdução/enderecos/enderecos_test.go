package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTest := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos imigrantes", "Rodovia"},
		{"Praça das rodas", "Tipo inválido"},
		{"Estrada qualquer", "Estrada"},
		{"RUA T", "Rua"},
		{"Avenida rua estrada", "Avenida"},
		{" ", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTest {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado)
		}
	}

}
