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
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
