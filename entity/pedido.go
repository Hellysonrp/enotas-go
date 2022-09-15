package entity

type Pedido struct {
	PresencaConsumidor string    `json:"presencaConsumidor,omitempty"`
	Pagamento          Pagamento `json:"pagamento,omitempty"`
}
