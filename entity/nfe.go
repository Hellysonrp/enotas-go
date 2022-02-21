package entity

type NFe struct {
	ID                          string                      `json:"id"`
	Ambiente                    Ambiente                    `json:"ambienteEmissao"`
	Tipo                        string                      `json:"tipo"`
	NaturezaOperacao            string                      `json:"naturezaOperacao"`
	Finalidade                  FinalidadeNFe               `json:"finalidade,omitempty"`
	ConsumidorFinal             bool                        `json:"consumidorFinal"`
	IndicadorPresencaConsumidor IndicadorPresencaConsumidor `json:"indicadorPresencaConsumidor"`
	Cliente                     Cliente                     `json:"cliente"`
	EnviarPorEmail              bool                        `json:"enviarPorEmail"`
	Itens                       []ItemNFe                   `json:"itens"`
	Transporte                  *Transporte                 `json:"transporte,omitempty"`
	InformacoesAdicionais       string                      `json:"informacoesAdicionais,omitempty"`
	Numero                      uint64                      `json:"numero"`
	Serie                       string                      `json:"serie"`
	// TODO resto do json completo
}

func NewNFe(c Cliente, itens []ItemNFe, a Ambiente) *NFe {
	return &NFe{
		Tipo:     "NF-e",
		Cliente:  c,
		Itens:    itens,
		Ambiente: a,
	}
}

func NewNFCe(c Cliente, itens []ItemNFe, a Ambiente) *NFe {
	return &NFe{
		Tipo:     "NFC-e",
		Cliente:  c,
		Itens:    itens,
		Ambiente: a,
	}
}
