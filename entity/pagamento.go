package entity

type Forma struct {
	Tipo      PagamentoTipo `json:"tipo,omitempty"`
	Valor     float64       `json:"valor,omitempty"`
	Descricao string        `json:"descricao,omitempty"`
}

type Pagamento struct {
	Tipo   PedidoPagamentoTipo `json:"tipo,omitempty"`
	Formas []Forma             `json:"formas,omitempty"`
}

type Fatura struct {
	Numero        string  `json:"numero,omitempty"`
	Desconto      float64 `json:"desconto,omitempty"`
	ValorOriginal float64 `json:"valorOriginal,omitempty"`
}

type Parcela struct {
	Numero     string  `json:"numero,omitempty"`
	Valor      float64 `json:"valor,omitempty"`
	Vencimento string  `json:"vencimento,omitempty"`
}

type Cobranca struct {
	Fatura   Fatura    `json:"fatura,omitempty"`
	Parcelas []Parcela `json:"parcelas,omitempty"`
}
