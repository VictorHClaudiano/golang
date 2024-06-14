// TESTE DE UNIDADE
// T do test da função tem que ser maiusculo.
// O arquivo tem que possuir nome_test.go
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}


func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{ "Rua ABC", "rua" },
		{ "Avenida Paulista", "avenida" },
		{ "Rodovia dos Imigrantes", "rodovia" },
		{ "Praça das Rosas", "Tipo Inválido"},
		{ "Estrada Qualquer", "estrada" },
		{ "RUA DOS BOBOS", "rua" },
		{ "", "Tipo Inválido"},
	}
	
	for _, cenario := range cenariosDeTeste {
		TipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if TipoDeEnderecoRecebido != cenario.retornoEsperado{
			t.Errorf("O tipo recebido %s é diferente do esperado %s", 
				TipoDeEnderecoRecebido, 
				cenario.retornoEsperado,
			)
		}
	}
	
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}

//go test --coverprofile resultado.txt (para criar uma linha de relatorio)
//para ler o arquivo txt usar na linha de comando 'go tool cover --func=resultado.txt'

//go test -v (verboso)
//go test --cover