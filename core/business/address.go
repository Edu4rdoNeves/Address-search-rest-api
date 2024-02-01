package business

import (
	"encoding/json"
	"fmt"

	"github.com/Edu4rdoNeves/Address-search-api/constants"
	"github.com/Edu4rdoNeves/Address-search-api/model"
	"github.com/Edu4rdoNeves/Address-search-api/service"
)

type IAddressBusiness interface {
	FindAddressByViaCep(cep string) (*model.AddressByViaCep, error)
	FindAddressByCorreios(cep string) (*model.AddressByCorreios, error)
}

type AddressBusiness struct {
}

func NewAddressBusiness() IAddressBusiness {
	return &AddressBusiness{}
}

func (c *AddressBusiness) FindAddressByViaCep(cep string) (*model.AddressByViaCep, error) {
	url := fmt.Sprintf(constants.ViacepBaseUrl, cep)

	responseData, err := service.FindAddress(url)
	if err != nil {
		return nil, err
	}

	Address := &model.AddressByViaCep{}
	err = json.Unmarshal(responseData, Address)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. Error: %v", err)
	}

	return Address, nil
}

func (c *AddressBusiness) FindAddressByCorreios(cep string) (*model.AddressByCorreios, error) {
	url := fmt.Sprintf(constants.CorreiosBaseUrl, cep)

	responseData, err := service.FindAddress(url)
	if err != nil {
		return nil, err
	}

	address := &model.AddressByCorreios{}
	err = json.Unmarshal(responseData, address)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. Error: %v", err)
	}

	return address, nil
}
