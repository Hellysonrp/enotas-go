package entity

import (
	"encoding/json"
	"regexp"
)

type Cliente struct {
	TipoPessoa                TipoPessoa `json:"tipoPessoa"`
	Nome                      string     `json:"nome"`
	Email                     string     `json:"email"`
	CpfCnpj                   string     `json:"cpfCnpj"`
	InscricaoMunicipal        string     `json:"instricaoMunicipal,omitempty"`
	InscricaoEstadual         string     `json:"inscricaoEstadual,omitempty"`
	IndicadorContribuinteICMS string     `json:"indicadorContribuinteICMS,omitempty"`
	Telefone                  string     `json:"telefone,omitempty"`
	Endereco                  Endereco   `json:"endereco"`
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
		Endereco: end,
	}
	c.AutoTipoPessoa()
	return c
}

func (c *Cliente) AutoTipoPessoa() {
	c.TipoPessoa = tipoPessoa(c.CpfCnpj)
}

func tipoPessoa(doc string) TipoPessoa {
	if len(onlyNumbers(doc)) == 14 {
		return PessoaJuridica
	}
	return PessoaFisica
}

func onlyNumbers(s string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(s, "")
}
