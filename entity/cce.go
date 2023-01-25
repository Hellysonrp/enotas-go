package entity

import "time"

type CCeNFe struct {
	ID          string `json:"id,omitempty"`
	ChaveAcesso string `json:"chaveAcesso"`
}

type CCe struct {
	ID       string   `json:"id"`
	Ambiente Ambiente `json:"ambienteEmissao"`
	Numero   uint64   `json:"numero"`
	Correcao string   `json:"string"`
	NFe      CCeNFe   `json:"nfe"`
}

func NewCCe(a Ambiente) *CCe {
	return &CCe{
		Ambiente: a,
	}
}

func NewCCeNFe(chaveAcesso string) CCeNFe {
	return CCeNFe{
		ChaveAcesso: chaveAcesso,
	}
}

type CCeResponse struct {
	ID                   string     `json:"id,omitempty"`
	AmbienteEmissao      string     `json:"ambienteEmissao,omitempty"`
	Status               string     `json:"status,omitempty"`
	MotivoStatus         string     `json:"motivoStatus,omitempty"`
	Numero               uint64     `json:"numero,omitempty"`
	Correcao             string     `json:"correcao,omitempty"`
	CondicoesUso         string     `json:"condicoesUso,omitempty"`
	Nfe                  CCeNFe     `json:"nfe,omitempty"`
	ProtocoloAutorizacao string     `json:"protocoloAutorizacao,omitempty"`
	DataCriacao          *time.Time `json:"dataCriacao,omitempty"`
	DataUltimaAlteracao  *time.Time `json:"dataUltimaAlteracao,omitempty"`
}
