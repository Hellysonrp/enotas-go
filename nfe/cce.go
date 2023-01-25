package nfe

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/padmoney/enotas-go/config"
	"github.com/padmoney/enotas-go/entity"
	"github.com/padmoney/enotas-go/internal/rest"
)

type CCe struct {
	client rest.Client
}

func NewCCe(client rest.Client) *CCe {
	return &CCe{
		client: client,
	}
}

func (c *CCe) Create(empresaID string, cce *entity.CCe) error {
	url := fmt.Sprintf("%s/empresas/%s/nf-e/cartaCorrecao", config.EndpointV2, empresaID)
	body, err := json.Marshal(cce)
	if err != nil {
		fmt.Printf("error in emitir: body %v\r\n", err.Error())
		return errors.New("error in body: " + err.Error())
	}
	response := c.client.Post(url, body)
	if response.Error != nil {
		// fmt.Printf("error in emitir: response %v", response.Error.Error())
		return errors.New("error in response: " + response.Error.Error())
	}
	if !response.Ok() {
		return errors.New("erro ao emitir a CCe; code: " + strconv.FormatInt(int64(response.Code), 10))
	}
	return nil
}

func (c *CCe) Consultar(empresaId, cceId string) (*entity.CCeResponse, error) {
	url := fmt.Sprintf("%s/empresas/%s/nf-e/cartaCorrecao/%s", config.EndpointV2, empresaId, cceId)
	response := c.client.Get(url)
	if response.Error != nil {
		return nil, response.Error
	}
	if !response.Ok() {
		return nil, errors.New("erro no retorno da consulta da Carta de Correção: " + strconv.FormatInt(int64(response.Code), 10))
	}
	resp := &entity.CCeResponse{}
	err := json.Unmarshal(response.Body, resp)
	if err != nil {
		return nil, errors.New("erro ao extrair dados do response consulta CC-e: " + string(response.Body[:]))
	}
	return resp, nil
}

func (c *CCe) ConsultarXML(empresaId, cceId string) ([]byte, error) {
	url := fmt.Sprintf("%s/empresas/%s/nf-e/cartaCorrecao/%s/xml", config.EndpointV2, empresaId, cceId)
	response := c.client.Send("GET", url, nil)
	if response.Error != nil {
		return nil, response.Error
	}
	if !response.Ok() {
		return nil, errors.New("erro no retorno da consulta da Carta de Correção: " + strconv.FormatInt(int64(response.Code), 10))
	}
	return response.Body, nil
}
