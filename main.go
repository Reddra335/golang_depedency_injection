package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Kipas interface {
	NyalaTidak(string) string
}

type Nyala struct{}

func (n *Nyala) NyalaTidak(nyala string) string {
	return "Mantap" + nyala
}

type StopKontak struct {
	Kipas22 Kipas
}

func main() {

	// db := app.NewDB()
	// validate := validator.New()
	// router := httprouter.New()
	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)

	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: middleware.NewAuthMiddleware(router),
	// }
	// err := server.ListenAndServe()
	// helper.ErrorT(err)

	stop := StopKontak{Kipas22: &Nyala{}}
	fmt.Println(stop.Kipas22.NyalaTidak("Nyala"))
}
