package entity

// Endereco do cliente
// Cidade: Nome da cidade ou seu código IBGE
// UF: Sigla do Estado (ES, MG, etc.)
type Endereco struct {
	Logradouro  string `json:"logradouro"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro"`
	CEP         string `json:"cep"`
	Cidade      string `json:"cidade"`
	UF          string `json:"uf"`
	Pais        string `json:"pais,omitempty"`
}

type EnderecoEmpresa struct {
	Endereco
	CodigoIbgeUF     int `json:"codigoIbgeUf"`
	CodigoIbgeCidade int `json:"codigoIbgeCidade"`
}
