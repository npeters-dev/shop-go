package service

import (
	"github.com/npeters-dev/shop-go/customer-service/module/customer/model"
	"github.com/npeters-dev/shop-go/customer-service/module/customer/repo"
)

type CustomerService interface {
	List() (model.CustomerDtos, error)
	Get(id int) (*model.CustomerDto, error)
	Add(form *model.CustomerForm) (*model.CustomerDto, error)
	Update(form *model.CustomerForm, id int) (*model.CustomerDto, error)
	Patch(form *model.CustomerPatchForm, id int) (*model.CustomerDto, error)
	Delete(id int) (*model.CustomerDto, error)
	Login(form *model.CustomerLoginForm) (*model.Customer, error)
}

type CustomerServiceConfig struct {
	CustomerRepo repo.CustomerRepo
	//Token        Token
}
