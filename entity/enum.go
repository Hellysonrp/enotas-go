package entity

type Ambiente string
type TipoPessoa string
type FinalidadeNFe string
type IndicadorPresencaConsumidor string

const (
	Ambiente_Homologacao Ambiente = "Homologacao"
	Ambiente_Producao    Ambiente = "Producao"

	TipoPessoa_PessoaFisica   TipoPessoa = "F"
	TipoPessoa_PessoaJuridica TipoPessoa = "J"

	FinalidadeNFe_Normal              FinalidadeNFe = "Normal"              // 1
	FinalidadeNFe_Complementar        FinalidadeNFe = "Complementar"        // 2 - não tenho certeza de que a string é essa mesmo; não possui documentação na enotas
	FinalidadeNFe_Ajuste              FinalidadeNFe = "Ajuste"              // 3 - não tenho certeza de que a string é essa mesmo; não possui documentação na enotas
	FinalidadeNFe_DevolucaoMercadoria FinalidadeNFe = "DevolucaoMercadoria" // 4

	IndicadorPresencaConsumidor_NaoSeAplica          IndicadorPresencaConsumidor = "NaoSeAplica"          // 0
	IndicadorPresencaConsumidor_OperacaoPresencial   IndicadorPresencaConsumidor = "OperacaoPresencial"   // 1
	IndicadorPresencaConsumidor_OperacaoPelaInternet IndicadorPresencaConsumidor = "OperacaoPelaInternet" // 2
	// TODO outros indicadores de presença
)
