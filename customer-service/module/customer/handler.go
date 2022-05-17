package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/npeters-dev/shop-go/customer-service/module/customer/service"
)

type Handler struct {
	CustomerService service.CustomerService
}

func (h *Handler) ListCustomers(c *gin.Context) {
	customers, err := h.CustomerService.List()

	if err != nil {
		InternalServerError(c, "error while listing customers")
	}

	SuccessResponse(c, "customers", customers)
}

func (h *Handler) GetCustomer(c *gin.Context) {
	customer, err := h.CustomerService.Get(c.GetInt("id"))

	if err != nil {
		InternalServerError(c, "error while getting customer")
	}

	SuccessResponse(c, "customer", customer)
}
