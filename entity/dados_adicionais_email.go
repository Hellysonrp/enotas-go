package entity

type DadosAdicionaisEmail struct {
	/*
		multiplos separar por ;
		ex:
		"dadosAdicionaisEmail": {
			"outrosDestinatarios": "contabilidade@mail.com; outroemail@mail.com"

			...
			}
	*/
	OutrosDestinatarios string `json:"outrosDestinatarios,omitempty"`
}
