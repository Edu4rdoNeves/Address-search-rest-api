package dependency

import (
	"github.com/Edu4rdoNeves/Address-search-api/core/business"
	"github.com/Edu4rdoNeves/Address-search-api/core/controller"
)

func CepDependency() controller.IAddressController {
	CepBusiness := business.NewAddressBusiness()
	cepController := controller.NewAddressController(CepBusiness)
	return cepController
}
