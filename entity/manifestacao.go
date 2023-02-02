package entity

type ConsultaManifestacaoResponse struct {
	HasMoreResults        bool        `json:"hasMoreResults,omitempty"`
	RetrievedRecordsCount int         `json:"retrievedRecordsCount,omitempty"`
	ContinuationToken     interface{} `json:"continuationToken,omitempty"`
	Records               []Records   `json:"records,omitempty"`
}

type HistoricoManifestacao struct {
	Tipo          string `json:"tipo,omitempty"`
	Justificativa string `json:"justificativa,omitempty"`
	Status        string `json:"status,omitempty"`
	MotivoStatus  string `json:"motivoStatus,omitempty"`
	Protocolo     string `json:"protocolo,omitempty"`
}
type Records struct {
	NFeResponse
	Tipo                  string                  `json:"tipo,omitempty"`
	DataCriacao           string                  `json:"dataCriacao,omitempty"`
	DataUltimaAlteracao   string                  `json:"dataUltimaAlteracao,omitempty"`
	DataAutorizacao       string                  `json:"dataAutorizacao,omitempty"`
	ValorTotal            float64                 `json:"valorTotal,omitempty"`
	Transporte            Transporte              `json:"transporte,omitempty"`
	Pedido                Pedido                  `json:"pedido,omitempty"`
	HistoricoManifestacao []HistoricoManifestacao `json:"historicoManifestacao,omitempty"`
}

type ManifestacaoResponse struct {
	Tipo            ManifestacaoTipo `json:"tipo,omitempty"`
	Justificativa   interface{}      `json:"justificativa,omitempty"`
	Status          string           `json:"status,omitempty"`
	MotivoStatus    interface{}      `json:"motivoStatus,omitempty"`
	DataAutorizacao string           `json:"dataAutorizacao,omitempty"`
	Protocolo       string           `json:"protocolo,omitempty"`
}

type ManifestacaoRequest struct {
	Tipo          ManifestacaoTipo `json:"tipo,omitempty"`
	Justificativa string           `json:"justificativa,omitempty"`
}
