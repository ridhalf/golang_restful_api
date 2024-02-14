package helper

import (
	"github.com/ridhalf/belajar-golang-restful-api/model/domain"
	"github.com/ridhalf/belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		Address: customer.Address,
	}
}
func ToCustomerResponses(customers []domain.Customer) []web.CustomerResponse {
	var customerResponse []web.CustomerResponse
	for _, customer := range customers {
		customerResponse = append(customerResponse, ToCustomerResponse(customer))
	}
	return customerResponse
}
