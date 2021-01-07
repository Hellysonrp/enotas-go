package entity

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/go-uuid"
)

// NFSE representa a Nota Fiscal de Serviço Eletrônica
type NFSe struct {
	ID                string     `json:"id,omitempty"`
	Tipo              string     `json:"tipo,omitempty"`
	IdExterno         string     `json:"idExterno,omitempty"`
	Status            string     `json:"status,omitempty"`
	MotivoStatus      string     `json:"motivoStatus,omitempty"`
	Ambiente          Ambiente   `json:"ambienteEmissao"`
	DataCompetencia   *time.Time `json:"dataCompetencia,omitempty"`
	Numero            int        `json:"numero,omitempty"`
	CodigoVerificacao string     `json:"codigoVerificacao,omitempty"`
	ChaveAcesso       string     `json:"chaveAcesso,omitempty"`
	NumeroRPS         int        `json:"numeroRps,omitempty"`
	SerieRPS          string     `json:"serieRps,omitempty"`
	EnviarPorEmail    bool       `json:"enviarPorEmail"`
	ValorTotal        float64    `json:"valorTotal"`
	Cliente           Cliente    `json:"cliente"`
	Servico           Servico    `json:"servico"`
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
