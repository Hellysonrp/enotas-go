package entity

type ItemNFe struct {
	CFOP          string          `json:"cfop"`
	Codigo        string          `json:"codigo"`
	Descricao     string          `json:"descricao"`
	NCM           string          `json:"ncm"`
	Quantidade    float64         `json:"quantidade"`
	UnidadeMedida string          `json:"unidadeMedida"`
	ValorUnitario float64         `json:"valorUnitario"`
	Frete         float64         `json:"frete"`
	Impostos      ItemNFeImpostos `json:"impostos"`
}

type ItemNFeImpostos struct {
	ICMS   *ItemNFeICMS  `json:"icms,omitempty"`
	Pis    ItemNFePis    `json:"pis"`
	Cofins ItemNFeCofins `json:"cofins"`
}

type ItemNFeICMS struct {
	SituacaoTributaria             string   `json:"situacaoTributaria"`
	Origem                         int      `json:"origem"`
	BaseCalculo                    *float64 `json:"baseCalculo,omitempty"`
	PercentualReducaoBaseCalculo   *float64 `json:"percentualReducaoBaseCalculo,omitempty"`
	Aliquota                       *float64 `json:"aliquota,omitempty"`
	PercentualReducaoBaseCalculoST *float64 `json:"percentualReducaoBaseCalculoST,omitempty"`
}

type ItemNFePis struct {
	SituacaoTributaria string              `json:"situacaoTributaria"`
	PorAliquota        *ItemNFePorAliquota `json:"porAliquota,omitempty"`
}

type ItemNFeCofins struct {
	SituacaoTributaria string              `json:"situacaoTributaria"`
	PorAliquota        *ItemNFePorAliquota `json:"porAliquota,omitempty"`
}

type ItemNFePorAliquota struct {
	Aliquota float64 `json:"aliquota"`
}

// TODO percentualAproximadoTributos
// TODO ipi
// TODO combustivel
// TODO outros campos do json completo
