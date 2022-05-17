package service

import (
	"errors"
	"github.com/npeters-dev/shop-go/customer-service/module/customer/model"
	"github.com/npeters-dev/shop-go/customer-service/module/customer/repo"
)

type DefaultCustomerService struct {
	CustomerRepo repo.CustomerRepo
	//Token        Token
}

func NewDefaultCustomerService(config CustomerServiceConfig) CustomerService {
	return &DefaultCustomerService{
		CustomerRepo: config.CustomerRepo,
		//Token:        config.Token,
	}
}

func (s *DefaultCustomerService) List() (model.CustomerDtos, error) {
	customers, err := s.CustomerRepo.List()

	if err != nil {
		return nil, err
	}

	customerDtos := customers.ToDto()

	return customerDtos, nil
}

func (s *DefaultCustomerService) Get(id int) (*model.CustomerDto, error) {
	customer, err := s.CustomerRepo.Get(id)

	if err != nil {
		return nil, err
	}

	customerDto := customer.ToDto()

	return customerDto, nil
}

func (s *DefaultCustomerService) Add(form *model.CustomerForm) (*model.CustomerDto, error) {
	customerModel := form.ToModel()
	customerModel.Password = form.Password
	customer, err := s.CustomerRepo.Add(customerModel)

	if err != nil {
		return nil, err
	}

	customerDto := customer.ToDto()

	return customerDto, nil
}

func (s *DefaultCustomerService) Update(form *model.CustomerForm, id int) (*model.CustomerDto, error) {
	customer, err := s.CustomerRepo.Update(form, id)

	if err != nil {
		return nil, err
	}

	customerDto := customer.ToDto()

	return customerDto, nil
}

func (s *DefaultCustomerService) Patch(form *model.CustomerPatchForm, id int) (*model.CustomerDto, error) {
	customer, err := s.CustomerRepo.Patch(form, id)

	if err != nil {
		return nil, err
	}

	customerDto := customer.ToDto()

	return customerDto, nil
}

func (s *DefaultCustomerService) Delete(id int) (*model.CustomerDto, error) {
	customer, err := s.CustomerRepo.Delete(id)

	if err != nil {
		return nil, err
	}

	customerDto := customer.ToDto()

	return customerDto, nil
}

func (s *DefaultCustomerService) Login(form *model.CustomerLoginForm) (*model.Customer, error) {
	customer, err := s.CustomerRepo.GetByEmail(form.Email)

	if err != nil {
		return nil, err
	}

	//if !password.Compare(customer.Password, form.Password) {
	//	return nil, errors.New("wrong password")
	//}

	if customer.Password != form.Password {
		return nil, errors.New("wrong password")
	}

	return customer, nil
}
