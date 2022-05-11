// TESTE DE UNIDADE
package enderecos

import "testing"

//pacote nativo do go pra testing

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

//ponteiro de testing do modulo
//a funcao tem q comecar com o nome Test, seguido do nome da funcao
//o arquivo tem q conter o _test no final do nome
func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		// {"Praça das Rosas", "Tipo Invalido"},
		// {"Estrada Qualquer", "Estrada"},
		// {"RUA DOS BOBOS", "Rua"},
		// {"AVENIDA REBOUÇAS", "Avenida"},
		// {"", "Tipo Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		retonoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retonoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", cenario.enderecoEsperado, cenario.enderecoInserido)
		}
	}
}
