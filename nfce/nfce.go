package nfce

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/padmoney/enotas-go/config"
	"github.com/padmoney/enotas-go/entity"
	"github.com/padmoney/enotas-go/internal/rest"
)

type NFCe struct {
	client rest.Client
}

func NewNFCe(client rest.Client) *NFCe {
	return &NFCe{
		client: client,
	}
}

func (n *NFCe) Emitir(empresaID string, nfe *entity.NFCe) (*entity.NFCeResponse, error) {
	url := fmt.Sprintf("%s/empresas/%s/nfc-e", config.EndpointV2, empresaID)
	body, err := json.Marshal(nfe)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s", body)
	response := n.client.Post(url, body)
	if response.Error != nil {
		return nil, response.Error
	}
	if !response.Ok() {
		return nil, errors.New("erro ao emitir a NFCe; code: " + strconv.FormatInt(int64(response.Code), 10))
	}

	resp := &entity.NFCeResponse{}
	err = json.Unmarshal(response.Body, resp)
	if err != nil {
		return nil, errors.New("erro ao extrair dados do response" + string(response.Body[:]))
	}
	return resp, nil
}

func (n *NFCe) Cancelar(empresaID, id string) error {
	url := fmt.Sprintf("%s/empresas/%s/nfc-e/%s", config.EndpointV2, empresaID, id)
	response := n.client.Delete(url)
	if response.Error != nil {
		return response.Error
	}
	if !response.Ok() {
		return errors.New("erro no retorno do cancelamento da Nota Fiscal: " + strconv.FormatInt(int64(response.Code), 10))
	}

	return nil
}
