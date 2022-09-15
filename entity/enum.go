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
	FinalidadeNFe_Complementar        FinalidadeNFe = "Complementar"        // 2
	FinalidadeNFe_Ajuste              FinalidadeNFe = "Ajuste"              // 3
	FinalidadeNFe_DevolucaoMercadoria FinalidadeNFe = "DevolucaoMercadoria" // 4

	IndicadorPresencaConsumidor_NaoSeAplica                             IndicadorPresencaConsumidor = "NaoSeAplica"                             // 0
	IndicadorPresencaConsumidor_OperacaoPresencial                      IndicadorPresencaConsumidor = "OperacaoPresencial"                      // 1
	IndicadorPresencaConsumidor_OperacaoPelaInternet                    IndicadorPresencaConsumidor = "OperacaoPelaInternet"                    // 2
	IndicadorPresencaConsumidor_Teleatendimento                         IndicadorPresencaConsumidor = "OperacaoTeleAtendimento"                 // 3
	IndicadorPresencaConsumidor_EntregaADomicilio                       IndicadorPresencaConsumidor = "EntregaADomicilio"                       // 4
	IndicadorPresencaConsumidor_OperacaoPresencialForaDoEstabelecimento IndicadorPresencaConsumidor = "OperacaoPresencialForaDoEstabelecimento" // 5
	IndicadorPresencaConsumidor_Outros                                  IndicadorPresencaConsumidor = "Outros"                                  // 9

	ModalidadeFrete_PorContaDoEmitente                      ModalidadeFrete = "ContratacaoPorContaDoRemetente"          // 0
	ModalidadeFrete_PorContaDoDestinatario                  ModalidadeFrete = "ContratacaoPorContaDoDestinatario"       // 1
	ModalidadeFrete_PorContaDeTerceiros                     ModalidadeFrete = "ContratacaoPorContaDeTerceiros"          // 2
	ModalidadeFrete_TransporteProprioPorContaDoRemetente    ModalidadeFrete = "TransporteProprioPorContaDoRemetente"    // 3
	ModalidadeFrete_TransporteProprioPorContaDoDestinatario ModalidadeFrete = "TransporteProprioPorContaDoDestinatario" // 4
	ModalidadeFrete_SemFrete                                ModalidadeFrete = "SemOcorrenciaDeTransporte"               // 9
)
