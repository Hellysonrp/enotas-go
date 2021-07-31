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
	if err != nil {
		return err
	}
	response := n.client.Post(url, body)
	if response.Error != nil {
		return response.Error
	}
	if !response.Ok() {
		return errors.New("erro ao emitir a NFe; code: " + strconv.FormatInt(int64(response.Code), 10))
	}
	return nil
}
