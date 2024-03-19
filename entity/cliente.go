package entity

import (
	"encoding/json"
	"regexp"
)

type Cliente struct {
	TipoPessoa                TipoPessoa `json:"tipoPessoa,omitempty"`
	Nome                      string     `json:"nome,omitempty"`
	Email                     string     `json:"email,omitempty"`
	CpfCnpj                   string     `json:"cpfCnpj,omitempty"`
	InscricaoMunicipal        string     `json:"inscricaoMunicipal,omitempty"`
	InscricaoEstadual         string     `json:"inscricaoEstadual,omitempty"`
	IndicadorContribuinteICMS string     `json:"indicadorContribuinteICMS,omitempty"`
	Telefone                  string     `json:"telefone,omitempty"`
	Endereco                  *Endereco  `json:"endereco,omitempty"`
}

func (c Cliente) MarshalJSON() ([]byte, error) {
	tp := c.TipoPessoa
	if string(c.TipoPessoa) == "" {
		tp = tipoPessoa(c.CpfCnpj)
	}

	type Alias Cliente
	return json.Marshal(&struct {
		TipoPessoa TipoPessoa `json:"tipoPessoa"`
		Alias
	}{
		Alias:      (Alias)(c),
		TipoPessoa: tp,
	})
}

func NewCliente(n string, d string, end Endereco) Cliente {
	c := Cliente{
		Nome:     n,
		CpfCnpj:  onlyNumbers(d),
		Endereco: &end,
	}
	c.AutoTipoPessoa()
	return c
}

func (c *Cliente) AutoTipoPessoa() {
	c.TipoPessoa = tipoPessoa(c.CpfCnpj)
}

func tipoPessoa(doc string) TipoPessoa {
	if len(onlyNumbers(doc)) == 14 {
		return TipoPessoa_PessoaJuridica
	}
	return TipoPessoa_PessoaFisica
}

func onlyNumbers(s string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(s, "")
}
