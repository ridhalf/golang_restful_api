package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/ridhalf/belajar-golang-restful-api/exception"
	"github.com/ridhalf/belajar-golang-restful-api/helper"
	"github.com/ridhalf/belajar-golang-restful-api/model/domain"
	"github.com/ridhalf/belajar-golang-restful-api/model/web"
	"github.com/ridhalf/belajar-golang-restful-api/repository"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		DB:                 DB,
		Validate:           validate}
}

func (service CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	customer := domain.Customer{
		Name:    request.Name,
		Address: request.Address,
	}
	customer = service.CustomerRepository.Save(ctx, tx, customer)
	return helper.ToCustomerResponse(customer)
}

func (service CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customer.Name = request.Name
	customer.Address = request.Address

	customer = service.CustomerRepository.Update(ctx, tx, customer)
	return helper.ToCustomerResponse(customer)
}

func (service CustomerServiceImpl) Delete(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.CustomerRepository.Delete(ctx, tx, customer)
}

func (service CustomerServiceImpl) FindById(ctx context.Context, customerId int) web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCustomerResponse(category)
}

func (service CustomerServiceImpl) FindAll(ctx context.Context) []web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := service.CustomerRepository.FindAll(ctx, tx)

	return helper.ToCustomerResponses(customer)
}
