package customer

import (
	"customer-service/module/customer/service"
	"github.com/gin-gonic/gin"
)

type HandlerConfig struct {
	R               *gin.Engine
	CustomerService service.CustomerService
}
