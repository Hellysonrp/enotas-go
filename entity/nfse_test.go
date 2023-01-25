package entity

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewNFSe(t *testing.T) {
	end := Endereco{}
	cliente := NewCliente("Mony", "89460862039", end)
	servico := NewServico("Serviço prestado", "1234567890", "1234", 5)
	nfse := NewNFSe(cliente, servico, 100, Ambiente_Homologacao)

	if idExterno := nfse.IdExterno; len(idExterno) != 36 {
		t.Errorf("Expected a valid idExterno, got '%s'", idExterno)
	}

	if enviarPorEmail := nfse.EnviarPorEmail; !enviarPorEmail {
		t.Errorf("EnviarPorEmail should not be false")
	}
}

func TestNFSeMarshalJSON(t *testing.T) {
	dataCompetencia, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	n := NFSe{
		ID:              "e93ceafb-6fd9-4543-9106-e0aebac9ef8d",
		Ambiente:        Ambiente_Homologacao,
		DataCompetencia: &dataCompetencia,
		Numero:          "42",
		EnviarPorEmail:  false,
		ValorTotal:      3.47,
		Cliente: Cliente{
			Nome:     "John Doe",
			Email:    "john@doe.com",
			CpfCnpj:  "47142365471",
			Endereco: &Endereco{},
		},
	}
	b, err := json.Marshal(n)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := `{"id":"e93ceafb-6fd9-4543-9106-e0aebac9ef8d","ambienteEmissao":"Homologacao","numero":42,"enviarPorEmail":false,"valorTotal":3.47,"cliente":{"tipoPessoa":"F","nome":"John Doe","email":"john@doe.com","cpfCnpj":"47142365471","endereco":{"logradouro":"","numero":"","bairro":"","cep":"","cidade":"","uf":""}},"servico":{"cnae":"","codigoServicoMunicipio":"","descricao":""},"dataCompetencia":"2006-01-02T15:04Z"}`
	if string(b) != expected {
		t.Errorf("Expected '%s', got %s", expected, string(b))
	}

}
