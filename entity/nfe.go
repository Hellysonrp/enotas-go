package entity

import "time"

type NFe struct {
	ID                          string                      `json:"id"`
	Ambiente                    Ambiente                    `json:"ambienteEmissao"`
	Tipo                        string                      `json:"tipo"`
	TipoOperacao                TipoOperacao                `json:"tipoOperacao,omitempty"`
	NaturezaOperacao            string                      `json:"naturezaOperacao"`
	Finalidade                  FinalidadeNFe               `json:"finalidade,omitempty"`
	ConsumidorFinal             bool                        `json:"consumidorFinal"`
	Pedido                      Pedido                      `json:"pedido,omitempty"`
	Cobranca                    Cobranca                    `json:"cobranca,omitempty"`
	IndicadorPresencaConsumidor IndicadorPresencaConsumidor `json:"indicadorPresencaConsumidor,omitempty"`
	Cliente                     Cliente                     `json:"cliente"`
	EnviarPorEmail              bool                        `json:"enviarPorEmail"`
	Itens                       []ItemNFe                   `json:"itens"`
	Transporte                  *Transporte                 `json:"transporte,omitempty"`
	InformacoesAdicionais       string                      `json:"informacoesAdicionais,omitempty"`
	Numero                      uint64                      `json:"numero"`
	Serie                       string                      `json:"serie"`
	NFeReferenciada             []NFeReferencia             `json:"nfeReferenciada,omitempty"`
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

type NFeResponse struct {
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

type NFeInitulizacao struct {
	Id            string   `json:"id"`
	Ambiente      Ambiente `json:"ambienteEmissao"`
	Serie         string   `json:"serie"`
	NumeroInicial uint64   `json:"numeroInicial"`
	NumeroFinal   uint64   `json:"numeroFinal"`
	Justificativa string   `json:"justificativa"`
}

type NFeInitulizacaoResponse struct {
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

type NFeReferencia struct {
	ChaveAcesso string `json:"chaveAcesso"`
}
