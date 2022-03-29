package entity

// Endereco do cliente
// Cidade: Nome da cidade ou seu código IBGE
// UF: Sigla do Estado (ES, MG, etc.)
type Endereco struct {
	Logradouro  string `json:"logradouro,omitempty"`
	Numero      string `json:"numero,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	CEP         string `json:"cep,omitempty"`
	Cidade      string `json:"cidade,omitempty"`
	UF          string `json:"uf,omitempty"`
	Pais        string `json:"pais,omitempty"`
}

type EnderecoEmpresa struct {
	Endereco
	CodigoIbgeUF     int `json:"codigoIbgeUf,omitempty"`
	CodigoIbgeCidade int `json:"codigoIbgeCidade,omitempty"`
}
