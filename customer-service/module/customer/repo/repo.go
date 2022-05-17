package repo

import "github.com/npeters-dev/shop-go/customer-service/module/customer/model"

type CustomerRepo interface {
	List() (model.Customers, error)
	Get(id int) (*model.Customer, error)
	Add(customer *model.Customer) (*model.Customer, error)
	Update(form *model.CustomerForm, id int) (*model.Customer, error)
	Patch(form *model.CustomerPatchForm, id int) (*model.Customer, error)
	Delete(id int) (*model.Customer, error)
	GetByEmail(email string) (*model.Customer, error)
}
