package entity

type Ambiente string
type TipoPessoa string
type FinalidadeNFe string
type IndicadorPresencaConsumidor string
type TipoOperacao string
type ModalidadeFrete string

const (
	Ambiente_Homologacao Ambiente = "Homologacao"
	Ambiente_Producao    Ambiente = "Producao"

	TipoPessoa_PessoaFisica   TipoPessoa = "F"
	TipoPessoa_PessoaJuridica TipoPessoa = "J"

	TipoOperacao_Saida   TipoOperacao = "Saida"
	TipoOperacao_Entrada TipoOperacao = "Entrada"

	FinalidadeNFe_Normal              FinalidadeNFe = "Normal"              // 1
	FinalidadeNFe_Complementar        FinalidadeNFe = "Complementar"        // 2 - não tenho certeza de que a string é essa mesmo; não possui documentação na enotas
	FinalidadeNFe_Ajuste              FinalidadeNFe = "Ajuste"              // 3 - não tenho certeza de que a string é essa mesmo; não possui documentação na enotas
	FinalidadeNFe_DevolucaoMercadoria FinalidadeNFe = "DevolucaoMercadoria" // 4

	IndicadorPresencaConsumidor_NaoSeAplica          IndicadorPresencaConsumidor = "NaoSeAplica"             // 0
	IndicadorPresencaConsumidor_OperacaoPresencial   IndicadorPresencaConsumidor = "OperacaoPresencial"      // 1
	IndicadorPresencaConsumidor_OperacaoPelaInternet IndicadorPresencaConsumidor = "OperacaoPelaInternet"    // 2
	IndicadorPresencaConsumidor_Teleatendimento      IndicadorPresencaConsumidor = "OperacaoTeleAtendimento" // 3
	IndicadorPresencaConsumidor_Outros               IndicadorPresencaConsumidor = "Outros"                  // 9
	// TODO outros indicadores de presença

	ModalidadeFrete_PorContaDoEmitente     ModalidadeFrete = "PorContaDoEmitente"     // 0
	ModalidadeFrete_PorContaDoDestinatario ModalidadeFrete = "PorContaDoDestinatario" // 1
	ModalidadeFrete_PorContaDeTerceiros    ModalidadeFrete = "PorContaDeTerceiros"    // 2
	ModalidadeFrete_SemFrete               ModalidadeFrete = "SemFrete"               // 9
	// TODO 3 e 4
)
