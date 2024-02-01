package server

import (
	"github.com/Edu4rdoNeves/Address-search-api/infrastructure/dependency"
	"github.com/gin-gonic/gin"
)

func ConfigRouters(router *gin.Engine) *gin.Engine {
	cepControllerWithDependencies := dependency.CepDependency()

	main := router.Group("api/v1")
	{
		cep := main.Group("address")
		{
			cep.GET("/:cep", cepControllerWithDependencies.FindAddress)
		}
	}
	return router
}
