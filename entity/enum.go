package entity

type Ambiente string
type TipoPessoa string
type FinalidadeNFe string
type IndicadorPresencaConsumidor string
type TipoOperacao string
type ModalidadeFrete string
type PedidoPagamentoTipo string
type PagamentoTipo string

type ModalidadeBaseCalculo int
type ModalidadeBaseCalculoST int

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

	PagamentoTipo_Outros                                    PagamentoTipo = "Outros"
	PagamentoTipo_Dinheiro                                  PagamentoTipo = "Dinheiro"
	PagamentoTipo_CartaoDeCredito                           PagamentoTipo = "CartaoDeCredito"
	PagamentoTipo_Cheque                                    PagamentoTipo = "Cheque"
	PagamentoTipo_CartaoDeDebito                            PagamentoTipo = "CartaoDeDebito"
	PagamentoTipo_CreditoLoja                               PagamentoTipo = "CreditoLoja"
	PagamentoTipo_ValeAlimentacao                           PagamentoTipo = "ValeAlimentacao"
	PagamentoTipo_ValeRefeicao                              PagamentoTipo = "ValeRefeicao"
	PagamentoTipo_ValePresente                              PagamentoTipo = "ValePresente"
	PagamentoTipo_ValeCombustivel                           PagamentoTipo = "ValeCombustivel"
	PagamentoTipo_BoletoBancario                            PagamentoTipo = "BoletoBancario"
	PagamentoTipo_DepositoBancario                          PagamentoTipo = "DepositoBancario"
	PagamentoTipo_PagamentoInstantaneoPix                   PagamentoTipo = "PagamentoInstantaneoPix"
	PagamentoTipo_TransferenciaBancariaCarteiraDigital      PagamentoTipo = "TransferenciaBancariaCarteiraDigital"
	PagamentoTipo_ProgramaFidelidadeCashbackCarteiraVirtual PagamentoTipo = "ProgramaFidelidadeCashbackCarteiraVirtual"
	PagamentoTipo_SemPagamento                              PagamentoTipo = "SemPagamento"

	PedidoPagamentoTipo_AVista PedidoPagamentoTipo = "PagamentoAVista"
	PedidoPagamentoTipo_APrazo PedidoPagamentoTipo = "PagamentoAPrazo"
	PedidoPagamentoTipo_Outros PedidoPagamentoTipo = "Outros"

	ModalidadeBaseCalculo_MargemValorAgregado ModalidadeBaseCalculo = 0
	ModalidadeBaseCalculo_Pauta               ModalidadeBaseCalculo = 1
	ModalidadeBaseCalculo_PrecoTabeladoMax    ModalidadeBaseCalculo = 2
	ModalidadeBaseCalculo_ValorDaOperacao     ModalidadeBaseCalculo = 3

	ModalidadeBaseCalculoST_PrecoTabelado       ModalidadeBaseCalculoST = 0
	ModalidadeBaseCalculoST_ListaPositiva       ModalidadeBaseCalculoST = 1
	ModalidadeBaseCalculoST_ListaNegativa       ModalidadeBaseCalculoST = 2
	ModalidadeBaseCalculoST_ListaNeutra         ModalidadeBaseCalculoST = 3
	ModalidadeBaseCalculoST_MargemValorAgregado ModalidadeBaseCalculoST = 4
	ModalidadeBaseCalculoST_Pauta               ModalidadeBaseCalculoST = 5
	ModalidadeBaseCalculoST_ValorDaOperacao     ModalidadeBaseCalculoST = 6
	ModalidadeBaseCalculoST_Nenhum              ModalidadeBaseCalculoST = 7
)
