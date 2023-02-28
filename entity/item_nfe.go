package entity

type ItemNFe struct {
	CFOP            string              `json:"cfop"`
	Codigo          string              `json:"codigo"`
	Descricao       string              `json:"descricao"`
	NCM             string              `json:"ncm"`
	Quantidade      float64             `json:"quantidade"`
	UnidadeMedida   string              `json:"unidadeMedida"`
	ValorUnitario   float64             `json:"valorUnitario"`
	Descontos       float64             `json:"descontos,omitempty"`
	Frete           float64             `json:"frete"`
	Impostos        ItemNFeImpostos     `json:"impostos"`
	CodigoBeneficio string              `json:"codigoBeneficioFiscal,omitempty"`
	Combustivel     *ItemNFeCombustivel `json:"combustivel,omitempty"`
}

type ItemNFeCombustivel struct {
	CodigoProdutoANP string `json:"codigoProdutoANP,omitempty"`
	UfConsumo        string `json:"ufConsumo,omitempty"`
}

type ItemNFeImpostos struct {
	ICMS   *ItemNFeICMS  `json:"icms,omitempty"`
	Pis    ItemNFePis    `json:"pis"`
	Cofins ItemNFeCofins `json:"cofins"`
}

type ItemNFeICMS struct {
	SituacaoTributaria             string                 `json:"situacaoTributaria"`
	Origem                         int                    `json:"origem"`
	BaseCalculo                    *float64               `json:"baseCalculo,omitempty"`
	PercentualReducaoBaseCalculo   *float64               `json:"percentualReducaoBaseCalculo,omitempty"`
	Aliquota                       *float64               `json:"aliquota,omitempty"`
	PercentualReducaoBaseCalculoST *float64               `json:"percentualReducaoBaseCalculoST,omitempty"`
	ModalidadeBaseCalculo          *ModalidadeBaseCalculo `json:"modalidadeBaseCalculo,omitempty"`
	NaoCalculaDifal                bool                   `json:"naoCalcularDifal,omitempty"`
	NaoCalcularFCP                 bool                   `json:"naoCalcularFCP,omitempty"`
	// 4 - Margem Valor Agregado
	ModalidadeBaseCalculoST *ModalidadeBaseCalculoST `json:"modalidadeBaseCalculoST,omitempty"`
	PercentualMVA           *float64                 `json:"percentualMargemValorAdicionadoST,omitempty"`
	BaseCalculoST           *float64                 `json:"baseCalculoST,omitempty"`
	AliquotaST              *float64                 `json:"aliquotaST,omitempty"`
	ValorST                 *float64                 `json:"valorST,omitempty"`
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
