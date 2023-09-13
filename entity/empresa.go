package entity

type ConfigNFSe struct {
	Sequencial     int    `json:"sequencialNFe"`
	Serie          string `json:"serieNFe"`
	SequencialLote int    `json:"sequencialLoteNFe"`
	Usuario        string `json:"usuarioAcessoProvedor"`
	Senha          string `json:"senhaAcessoProvedor"`
	TokenAcesso    string `json:"tokenAcessoProvedor"`
}

type ConfigNFeAmbiente struct {
	SequencialNFe     int    `json:"sequencialNFe"`
	SerieNFe          string `json:"serieNFe"`
	SequencialLoteNFe int    `json:"sequencialLoteNFe"`
}

type ConfigNFe struct {
	AmbienteProducao    ConfigNFeAmbiente `json:"ambienteProducao"`
	AmbienteHomologacao ConfigNFeAmbiente `json:"ambienteHomologacao"`
}

type ConfigNFCeCsc struct {
	ID     string `json:"id"`
	Codigo string `json:"codigo"`
}

type ConfigNFCeAmbiente struct {
	SequencialNFe int           `json:"sequencialNFe"`
	SerieNFe      string        `json:"serieNFe"`
	Csc           ConfigNFCeCsc `json:"csc"`
}

type ConfigNFCe struct {
	AmbienteProducao    ConfigNFCeAmbiente `json:"ambienteProducao"`
	AmbienteHomologacao ConfigNFCeAmbiente `json:"ambienteHomologacao"`
}

type EmpresaBase struct {
	ID                     string          `json:"id"`
	CNPJ                   string          `json:"cnpj"`
	InscricaoMunicipal     string          `json:"inscricaoMunicipal"`
	InscricaoEstadual      string          `json:"inscricaoEstadual"`
	RazaoSocial            string          `json:"razaoSocial"`
	NomeFantasia           string          `json:"nomeFantasia"`
	OptanteSimplesNacional bool            `json:"optanteSimplesNacional"`
	Email                  string          `json:"email"`
	TelefoneComercial      string          `json:"telefoneComercial"`
	Endereco               EnderecoEmpresa `json:"endereco"`

	// dados retornados na consulta e não enviados no POST
	Status                       string `json:"status,omitempty"`
	DadosObrigatoriosPreenchidos bool   `json:"dadosObrigatoriosPreenchidos,omitempty"`
}

type Empresa struct {
	EmpresaBase

	// dados exclusivos da NFS-e
	IncentivadorCultural         bool       `json:"incentivadorCultural"`
	RegimeEspecialTributacao     string     `json:"regimeEspecialTributacao"`
	CodigoServicoMunicipal       string     `json:"codigoServicoMunicipal"`
	ItemListaServicoLC116        string     `json:"itemListaServicoLC116"`
	CNAE                         string     `json:"cnae"`
	AliquotaISS                  float64    `json:"aliquotaIss"`
	DescricaoServico             string     `json:"descricaoServico"`
	EnviarEmailCliente           bool       `json:"enviarEmailCliente"`
	ConfiguracoesNFSeHomologacao ConfigNFSe `json:"configuracoesNFSeHomologacao"`
	ConfiguracoesNFSeProducao    ConfigNFSe `json:"configuracoesNFSeProducao"`
	NomeEmpresaCertificado       string     `json:"nome"`
	DataVencimentoCertificado    string     `json:"dataVencimento"`
	MEI                          bool       `json:"mei"`
}

type EmpresaV2 struct {
	EmpresaBase

	EmissaoNFeProduto    *ConfigNFe  `json:"emissaoNFeProduto,omitempty"`
	EmissaoNFeConsumidor *ConfigNFCe `json:"emissaoNFeConsumidor,omitempty"`
}
