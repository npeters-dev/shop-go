package repo

import (
	"github.com/npeters-dev/shop-go/customer-service/module/customer/model"
	"gorm.io/gorm"
)

type DefaultCustomerRepo struct {
	Database *gorm.DB
}

func NewDefaultCustomerRepo(database *gorm.DB) DefaultCustomerRepo {
	return DefaultCustomerRepo{
		Database: database,
	}
}

// List @TODO: implement filtering & pagination
func (repo DefaultCustomerRepo) List() (model.Customers, error) {
	customers := make([]*model.Customer, 0)

	if err := repo.Database.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (repo DefaultCustomerRepo) Get(id int) (*model.Customer, error) {
	customer := new(model.Customer)
	err := repo.Database.
		Where("id = ?", id).
		First(&customer).Error

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo DefaultCustomerRepo) Add(customer *model.Customer) (*model.Customer, error) {
	if err := repo.Database.Create(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo DefaultCustomerRepo) Update(form *model.CustomerForm, id int) (*model.Customer, error) {
	customer := form.ToModel()

	if err := repo.Database.Where("id = ?", id).Updates(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo DefaultCustomerRepo) Patch(form *model.CustomerPatchForm, id int) (*model.Customer, error) {
	customer := form.ToModel()

	if err := repo.Database.Where("id = ?", id).Updates(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo DefaultCustomerRepo) Delete(id int) (*model.Customer, error) {
	customer := new(model.Customer)

	if err := repo.Database.Where("id = ?", id).Delete(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo DefaultCustomerRepo) GetByEmail(email string) (*model.Customer, error) {
	customer := new(model.Customer)

	if err := repo.Database.Where("email = ?", email).Delete(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}
