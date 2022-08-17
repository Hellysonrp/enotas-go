package entity

type Transporte struct {
	Frete           Frete           `json:"frete"`
	EnderecoEntrega EnderecoEntrega `json:"enderecoEntrega"`
	Transportadora  Transportadora  `json:"transportadora"`
	Veiculo         Veiculo         `json:"veiculo"`
	Volume          Volume          `json:"volume"`
}

type EnderecoEntrega struct {
	Endereco
	TipoPessoaDestinatario TipoPessoa `json:"tipoPessoaDestinatario"` // ?
	CpfCnpjDestinatario    string     `json:"cpfCnpjDestinatario"`
}

type Frete struct {
	Modalidade ModalidadeFrete `json:"modalidade"`
}

type Transportadora struct {
	UsarDadosEmitente bool       `json:"usarDadosEmitente"`
	TipoPessoa        TipoPessoa `json:"tipoPessoa"`
	CpfCnpj           string     `json:"cpfCnpj"`
	Nome              string     `json:"nome"`
	InscricaoEstadual string     `json:"inscricaoEstadual"`
	EnderecoCompleto  string     `json:"enderecoCompleto"`
	Cidade            string     `json:"cidade"`
	Uf                string     `json:"uf"`
}

type Veiculo struct {
	Placa string `json:"placa"`
	Uf    string `json:"uf"`
	Rntc  string `json:"rntc,omitempty"`
}

type Volume struct {
	Quantidade  float64 `json:"quantidade"`
	Especie     string  `json:"especie,omitempty"`
	Marca       string  `json:"marca,omitempty"`
	Numeracao   string  `json:"numeracao"`
	PesoLiquido float64 `json:"pesoLiquido"`
	PesoBruto   float64 `json:"pesoBruto"`
}
