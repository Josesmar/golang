package enderecos_test

import (
	"introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel() //Executar testes em paralello com outros testes que tenha essa chamada

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
		retornoRecebido := enderecos.TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel() //Executar testes em paralello com outros testes que tenha essa chamada

	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}
