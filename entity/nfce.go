package entity

import "time"

type NFCe struct {
	ID                          string                      `json:"id"`
	Ambiente                    Ambiente                    `json:"ambienteEmissao"`
	Tipo                        string                      `json:"tipo"`
	NaturezaOperacao            string                      `json:"naturezaOperacao"`
	EnviarPorEmail              bool                        `json:"enviarPorEmail"`
	Csc                         CSC                         `json:"csc"`
	Pedido                      Pedido                      `json:"pedido"`
	InformacoesAdicionais       string                      `json:"informacoesAdicionais,omitempty"`
	IndicadorPresencaConsumidor IndicadorPresencaConsumidor `json:"indicadorPresencaConsumidor,omitempty"`
	Numero                      uint64                      `json:"numero"`
	Serie                       string                      `json:"serie"`
	Itens                       []ItemNFe                   `json:"itens"`
	Cliente                     *Cliente                    `json:"cliente,omitempty"`
}

type CSC struct {
	IdToken string `json:"id"`
	Csc     string `json:"codigo"`
}

func NewNFCe(c *Cliente, itens []ItemNFe, a Ambiente) *NFCe {
	return &NFCe{
		Tipo:     "NFC-e",
		Cliente:  c,
		Itens:    itens,
		Ambiente: a,
	}
}

type QrCode struct {
	Conteudo string `json:"conteudo"`
}

type Protocolo struct {
	Numero      string `json:"numero"`
	DigestValue string `json:"digestValue"`
}

type NFCeResponse struct {
	Id                         string    `json:"id"`
	Ambiente                   Ambiente  `json:"ambienteEmissao"`
	Status                     string    `json:"status"`
	MotivoStatus               string    `json:"motivoStatus"`
	Numero                     uint64    `json:"numero"`
	Serie                      string    `json:"serie"`
	ChaveAcesso                string    `json:"chaveAcesso"`
	LinkDownloadXml            string    `json:"linkDownloadXml"`
	LinkDanfe                  string    `json:"linkDanfe"`
	LinkConsultaPorChaveAcesso string    `json:"linkConsultaPorChaveAcesso"`
	QrCode                     QrCode    `json:"qrCode"`
	Protocolo                  Protocolo `json:"protocolo"`
	InformacoesAdicionais      string    `json:"informacoesAdicionais"`
	DataEmissao                string    `json:"dataEmissao"`
}

type NFCeInitulizacao struct {
	Id            string   `json:"id"`
	Ambiente      Ambiente `json:"ambienteEmissao"`
	Serie         string   `json:"serie"`
	NumeroInicial uint64   `json:"numeroInicial"`
	NumeroFinal   uint64   `json:"numeroFinal"`
	Justificativa string   `json:"justificativa"`
}

type NFCeInitulizacaoResponse struct {
	Id                   string     `json:"id"`
	Ambiente             Ambiente   `json:"ambienteEmissao"`
	Status               string     `json:"status"`
	MotivoStatus         string     `json:"motivoStatus"`
	Serie                string     `json:"serie"`
	NumeroInicial        uint64     `json:"numeroInicial"`
	NumeroFinal          uint64     `json:"numeroFinal"`
	Justificativa        string     `json:"justificativa"`
	ProtocoloAutorizacao string     `json:"protocoloAutorizacao"`
	DataCriacao          *time.Time `json:"dataCriacao"`
	DataUltimaAlteracao  *time.Time `json:"dataUltimaAlteracao"`
}
