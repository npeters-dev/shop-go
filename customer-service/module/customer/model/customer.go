package model

import (
	"gorm.io/gorm"
	"time"
)

type Customers []*Customer

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CustomerDtos []*CustomerDto

type CustomerDto struct {
	ID        uint      `json:"id" example:"123"`
	FirstName string    `json:"first_name" example:"Max"`
	LastName  string    `json:"last_name" example:"Mustermann"`
	Email     string    `json:"email" example:"name@email.com"`
	CreatedAt time.Time `json:"created_at" example:"2021-02-02T02:52:24Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-02-02T02:52:24Z"`
}

type CustomerForm struct {
	FirstName string `json:"first_name" label:"first name" example:"Max" valid:"Required;MinSize(2);MaxSize(255)"`
	LastName  string `json:"last_name" label:"last name" example:"Mustermann" valid:"MaxSize(255)"`
	Email     string `json:"email" label:"email" example:"name@email.com" valid:"Required;Email;"`
	Password  string `json:"password" label:"password"  example:"max123" valid:"MinSize(5);MaxSize(20)"`
}

type CustomerLoginForm struct {
	Email    string `json:"email" label:"email" example:"name@email.com" valid:"Required;Email;"`
	Password string `json:"password" label:"password"  example:"max123" valid:"MinSize(5);MaxSize(20)"`
}

type CustomerPatchForm struct {
	Email     string `json:"email" example:"name@email.com" valid:"Email;"`
	FirstName string `json:"firstName" example:"John" valid:"MinSize(2);MaxSize(255)"`
	LastName  string `json:"lastName" example:"Doe" valid:"MaxSize(255)"`
}

func (c *Customer) ToDto() *CustomerDto {
	return &CustomerDto{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (c Customers) ToDto() CustomerDtos {
	dtos := make([]*CustomerDto, len(c))

	for i, customer := range c {
		dtos[i] = customer.ToDto()
	}

	return dtos
}

func (form *CustomerForm) ToModel() *Customer {
	return &Customer{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Email:     form.Email,
	}
}

func (form *CustomerPatchForm) ToModel() *Customer {
	return &Customer{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Email:     form.Email,
	}
}
