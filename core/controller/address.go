package controller

import (
	"net/http"

	"time"

	"github.com/Edu4rdoNeves/Address-search-api/core/business"
	"github.com/Edu4rdoNeves/Address-search-api/model"
	"github.com/gin-gonic/gin"
)

type IAddressController interface {
	FindAddress(context *gin.Context)
}

type AddressController struct {
	business business.IAddressBusiness
}

func NewAddressController(Business business.IAddressBusiness) IAddressController {
	return &AddressController{Business}
}

func (c *AddressController) FindAddress(context *gin.Context) {
	cep := context.Param("cep")

	viacep := make(chan *model.AddressByViaCep)
	correios := make(chan *model.AddressByCorreios)

	go func() {
		Address, err := c.business.FindAddressByViaCep(cep)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"Error:": "Failed to find Address " + err.Error(),
			})
			return
		}
		viacep <- Address
	}()

	go func() {
		Address, err := c.business.FindAddressByCorreios(cep)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"Error:": "Failed to find Address " + err.Error(),
			})
			return
		}
		correios <- Address
	}()

	select {
	case addressByViaCep := <-viacep:
		context.JSON(http.StatusOK, addressByViaCep)

	case addressByCorreios := <-correios:
		context.JSON(http.StatusOK, addressByCorreios)

	case <-time.After(time.Second * 3):
		context.JSON(http.StatusOK, gin.H{
			"Error:": "expired waiting time",
		})
	}

}
