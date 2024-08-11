package main

import (
	"HafidzDaffa/belajar-golang-restful-api/app"
	"HafidzDaffa/belajar-golang-restful-api/controller"
	"HafidzDaffa/belajar-golang-restful-api/helper"
	"HafidzDaffa/belajar-golang-restful-api/middleware"
	"HafidzDaffa/belajar-golang-restful-api/repository"
	"HafidzDaffa/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
