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

type NFe struct {
	client rest.Client
}

func NewNFe(client rest.Client) *NFe {
	return &NFe{
		client: client,
	}
}

func (n *NFe) Emitir(empresaID string, nfe *entity.NFe) error {
	url := fmt.Sprintf("%s/empresas/%s/nf-e", config.EndpointV2, empresaID)
	body, err := json.Marshal(nfe)
	// fmt.Printf("vai la o negocio: %v", string(body[:]))
	if err != nil {
		// fmt.Printf("error in emitir: body %v", err.Error())
		return errors.New("error in body: " + err.Error())
	}
	response := n.client.Post(url, body)
	if response.Error != nil {
		// fmt.Printf("error in emitir: response %v", response.Error.Error())
		return errors.New("error in response: " + response.Error.Error())
	}
	if !response.Ok() {
		return errors.New("erro ao emitir a NFe; code: " + strconv.FormatInt(int64(response.Code), 10))
	}
	return nil
}

func (n *NFe) Cancelar(empresaID, id string) error {
	url := fmt.Sprintf("%s/empresas/%s/nf-e/%s", config.EndpointV2, empresaID, id)
	response := n.client.Delete(url)
	if response.Error != nil {
		return response.Error
	}
	if !response.Ok() {
		return errors.New("erro no retorno do cancelamento da Nota Fiscal: " + strconv.FormatInt(int64(response.Code), 10))
	}

	return nil
}

func (n *NFe) Consultar(empresaID, id string) (*entity.NFeResponse, error) {
	url := fmt.Sprintf("%s/empresas/%s/nf-e/%s", config.EndpointV2, empresaID, id)
	response := n.client.Get(url)
	if response.Error != nil {
		return nil, response.Error
	}
	if !response.Ok() {
		return nil, errors.New("erro no retorno da consulta da Nota Fiscal: " + strconv.FormatInt(int64(response.Code), 10))
	}
	resp := &entity.NFeResponse{}
	err := json.Unmarshal(response.Body, resp)
	if err != nil {
		return nil, errors.New("erro ao extrair dados do response consulta NFe: " + string(response.Body[:]))
	}
	return resp, nil
}

func (n *NFe) Inutilizar(empresaID string, inut *entity.NFeInitulizacao) error {
	url := fmt.Sprintf("%s/empresas/%s/nf-e/inutilizacao", config.EndpointV2, empresaID)
	body, err := json.Marshal(inut)
	if err != nil {
		return err
	}
	response := n.client.Post(url, body)
	if response.Error != nil {
		return response.Error
	}
	if !response.Ok() {
		return errors.New("erro ao inutilizar a NFe; code: " + strconv.FormatInt(int64(response.Code), 10))
	}

	return nil
}

func (n *NFe) ConsultarInutilizacao(empresaID, inutID string) (*entity.NFeInitulizacaoResponse, error) {
	url := fmt.Sprintf("%s/empresas/%s/nf-e/inutilizacao/%s", config.EndpointV2, empresaID, inutID)
	response := n.client.Get(url)
	if response.Error != nil {
		return nil, response.Error
	}
	if !response.Ok() {
		return nil, errors.New("erro ao consultar inutilizacao da NFe; code: " + strconv.FormatInt(int64(response.Code), 10))
	}

	resp := &entity.NFeInitulizacaoResponse{}
	err := json.Unmarshal(response.Body, resp)
	if err != nil {
		return nil, errors.New("erro ao extrair dados do response inutilizacao " + string(response.Body[:]))
	}
	return resp, nil
}

func (n *NFe) ConsultarXmlInutilizacao(empresaID, inutID string) ([]byte, error) {
	url := fmt.Sprintf("%s/empresas/%s/nf-e/inutilizacao/%s/xml", config.EndpointV2, empresaID, inutID)
	response := n.client.Get(url)
	if response.Error != nil {
		return nil, response.Error
	}
	if !response.Ok() {
		return nil, errors.New("erro ao consultar XML inutilizacao da NFe; code: " + strconv.FormatInt(int64(response.Code), 10))
	}

	return response.Body, nil
}
