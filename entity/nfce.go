package entity

type NFCe struct {
	ID                    string    `json:"id"`
	Ambiente              Ambiente  `json:"ambienteEmissao"`
	Tipo                  string    `json:"tipo"`
	NaturezaOperacao      string    `json:"naturezaOperacao"`
	EnviarPorEmail        bool      `json:"enviarPorEmail"`
	Csc                   CSC       `json:"csc"`
	Pedido                Pedido    `json:"pedido"`
	InformacoesAdicionais string    `json:"informacoesAdicionais,omitempty"`
	Numero                uint64    `json:"numero"`
	Serie                 string    `json:"serie"`
	Itens                 []ItemNFe `json:"itens"`
	Cliente               Cliente   `json:"cliente,omitempty"`
}

type CSC struct {
	IdToken string `json:"id"`
	Csc     string `json:"codigo"`
}

type Forma struct {
	Tipo  string  `json:"tipo"`
	Valor float64 `json:"valor"`
}

type Pagamento struct {
	Tipo   string  `json:"tipo"`
	Formas []Forma `json:"formas"`
}

type Pedido struct {
	PresencaConsumidor string    `json:"presencaConsumidor"`
	Pagamento          Pagamento `json:"pagamento"`
}

func NewNFCe(c Cliente, itens []ItemNFe, a Ambiente) *NFCe {
	return &NFCe{
		Tipo:     "NFC-e",
		Cliente:  c,
		Itens:    itens,
		Ambiente: a,
	}
}
