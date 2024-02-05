package entity

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/go-uuid"
)

type MetadataItem struct {
	DiscriminacaoServico   string         `json:"discriminacaoServico,omitempty"`
	Quantidade             float64        `json:"quantidade,omitempty"`
	ValorUnitario          float64        `json:"valorUnitario,omitempty"`
	Tributavel             ItemTributavel `json:"tributavel,omitempty"`
	AliquotaISS            *float64       `json:"aliquotaIss,omitempty"`
	ItemListaServicoLC116  *string        `json:"itemListaServicoLC116,omitempty"`
	CodigoServicoMunicipio *string        `json:"codigoServicoMunicipio,omitempty"`
}

type Metadata struct {
	Itens []MetadataItem `json:"itens"`
}

// NFSE representa a Nota Fiscal de Serviço Eletrônica
type NFSe struct {
	ID                   string                `json:"id,omitempty"`
	Tipo                 string                `json:"tipo,omitempty"`
	IdExterno            string                `json:"idExterno,omitempty"`
	Status               string                `json:"status,omitempty"`
	MotivoStatus         string                `json:"motivoStatus,omitempty"`
	Ambiente             Ambiente              `json:"ambienteEmissao"`
	DataCompetencia      *time.Time            `json:"dataCompetencia,omitempty"`
	Numero               string                `json:"numero,omitempty"`
	CodigoVerificacao    string                `json:"codigoVerificacao,omitempty"`
	ChaveAcesso          string                `json:"chaveAcesso,omitempty"`
	LinkDownloadPDF      string                `json:"linkDownloadPDF,omitempty"`
	LinkDownloadXML      string                `json:"linkDownloadXML,omitempty"`
	NumeroRPS            int                   `json:"numeroRps,omitempty"`
	SerieRPS             string                `json:"serieRps,omitempty"`
	EnviarPorEmail       bool                  `json:"enviarPorEmail"`
	DadosAdicionaisEmail *DadosAdicionaisEmail `json:"dadosAdicionaisEmail,omitempty"`
	ValorTotal           float64               `json:"valorTotal"`
	Cliente              Cliente               `json:"cliente"`
	Servico              Servico               `json:"servico"`
	Metadata             *Metadata             `json:"metadados,omitempty"`
	/* Natureza da Operação
	1 - Dentro do municipio
	2 - Fora do municipio
	se não enviar, o enotas faz a validação corretamente
	*/
	NaturezaOperacao *string `json:"naturezaOperacao,omitempty"`
}

func (n NFSe) MarshalJSON() ([]byte, error) {
	var dataCompetencia string
	if n.DataCompetencia != nil {
		dataCompetencia = n.DataCompetencia.Format("2006-01-02T15:04Z")
	}

	type Alias NFSe
	return json.Marshal(&struct {
		Alias
		DataCompetencia string `json:"dataCompetencia,omitempty"`
	}{
		Alias:           (Alias)(n),
		DataCompetencia: dataCompetencia,
	})
}

// NewNFSe cria um nova nota fiscal
func NewNFSe(c Cliente, s Servico, v float64, a Ambiente) *NFSe {
	return &NFSe{
		IdExterno:      newUUID(),
		Tipo:           "NFS-e",
		Ambiente:       a,
		EnviarPorEmail: true,
		Cliente:        c,
		Servico:        s,
		ValorTotal:     v,
	}
}

func newUUID() (id string) {
	id, _ = uuid.GenerateUUID()
	return
}
