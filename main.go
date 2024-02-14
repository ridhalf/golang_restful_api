package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/ridhalf/belajar-golang-restful-api/app"
	"github.com/ridhalf/belajar-golang-restful-api/controller"
	"github.com/ridhalf/belajar-golang-restful-api/exception"
	"github.com/ridhalf/belajar-golang-restful-api/helper"
	"github.com/ridhalf/belajar-golang-restful-api/middleware"
	"github.com/ridhalf/belajar-golang-restful-api/model/web"
	"github.com/ridhalf/belajar-golang-restful-api/repository"
	"github.com/ridhalf/belajar-golang-restful-api/service"
	"net/http"
)

func main() {

	validate := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepositoryImplementation()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	customerRepository := repository.NewCustomerRepositoryImplementation()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
		}

		helper.WriteToResponseBody(writer, webResponse)
	})
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/:customerId", customerController.FindById)
	router.POST("/api/customers", customerController.Create)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
