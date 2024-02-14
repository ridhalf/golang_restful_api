package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ridhalf/belajar-golang-restful-api/helper"
	"github.com/ridhalf/belajar-golang-restful-api/model/domain"
)

type CustomerRepositoryImplementation struct {
}

func NewCustomerRepositoryImplementation() CustomerRepository {
	return &CustomerRepositoryImplementation{}
}

func (repository *CustomerRepositoryImplementation) Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "INSERT INTO customer(name,address) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Address)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	customer.Id = int(id)
	return customer
}

func (repository *CustomerRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "UPDATE customer SET name = ?, address = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Address, customer.Id)
	helper.PanicIfError(err)
	return customer
}

func (repository *CustomerRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer) {
	SQL := "DELETE FROM customer WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Id)
	helper.PanicIfError(err)
}

func (repository *CustomerRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error) {
	SQL := "SELECT id, name, address FROM customer WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()
	customer := domain.Customer{}
	if rows.Next() {
		err = rows.Scan(&customer.Id, &customer.Name, &customer.Address)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("customer is not found")
	}
}

func (repository *CustomerRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	SQL := "SELECT id, name, address FROM customer"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var customers []domain.Customer
	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}
	return customers
}
