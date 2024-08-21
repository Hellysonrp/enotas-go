package entity

// Servico representa o objeto contendo os dados do serviço prestado
type Servico struct {
	CNAE                      string   `json:"cnae"`
	CodigoServicoMunicipio    string   `json:"codigoServicoMunicipio"`
	Descricao                 string   `json:"descricao"`
	AliquotaISS               *float64 `json:"aliquotaIss,omitempty"`
	IssRetidoFonte            *bool    `json:"issRetidoFonte,omitempty"`
	ValorPIS                  *float64 `json:"valorPis,omitempty"`
	ValorCOFINS               *float64 `json:"valorCofins,omitempty"`
	ValorCSLL                 *float64 `json:"valorCsll,omitempty"`
	ValorINSS                 *float64 `json:"valorInss,omitempty"`
	ValorIR                   *float64 `json:"valorIr,omitempty"`
	ValorISS                  *float64 `json:"valorIss,omitempty"`
	MunicipioPrestacaoServico *string  `json:"municipioPrestacaoServico,omitempty,string"`
}

// NewService cria um novo serviço
func NewServico(d string, c string, cnae string, iss float64) Servico {
	return Servico{
		Descricao:              d,
		CodigoServicoMunicipio: c,
		CNAE:                   cnae,
		AliquotaISS:            &iss,
	}
}
