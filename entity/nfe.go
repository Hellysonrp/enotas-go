package entity

type NFe struct {
	ID                          string                      `json:"id"`
	Ambiente                    Ambiente                    `json:"ambienteEmissao"`
	NaturezaOperacao            string                      `json:"naturezaOperacao"`
	Finalidade                  FinalidadeNFe               `json:"finalidade,omitempty"`
	ConsumidorFinal             bool                        `json:"consumidorFinal"`
	IndicadorPresencaConsumidor IndicadorPresencaConsumidor `json:"indicadorPresencaConsumidor"`
	Cliente                     Cliente                     `json:"cliente"`
	EnviarPorEmail              bool                        `json:"enviarPorEmail"`
	Itens                       []ItemNFe                   `json:"itens"`
	Transporte                  *Transporte                 `json:"transporte,omitempty"`
	InformacoesAdicionais       string                      `json:"informacoesAdicionais,omitempty"`
	// TODO resto do json completo
}
