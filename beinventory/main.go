package main

import (
	_ "database/sql"
	"fmt"
	"net/http"

	"github.com/Jehanv60/app"
	"github.com/Jehanv60/controller"
	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/middleware"
	"github.com/Jehanv60/repository"
	"github.com/Jehanv60/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	DB := app.NewDb()
	validate := validator.New()
	barangRepository := repository.NewRepositoryBarang()
	barangService := service.NewBarangService(barangRepository, DB, validate)
	penggunaRepository := repository.NewRepositoryPengguna()
	penggunaService := service.NewPenggunaService(penggunaRepository, DB, validate)
	transaksiRepository := repository.NewRepositoryTransaksi()
	transaksiService := service.NewTransaksiService(transaksiRepository, barangRepository, DB, validate)
	barangController := controller.NewBarangController(barangService, penggunaService)
	penggunaController := controller.NewPenggunaController(penggunaService)
	transaksiController := controller.NewTransaksiController(transaksiService, penggunaService, barangService)
	router := app.NewRouter(barangController, penggunaController, transaksiController)
	Handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true, // ← penting kalau kamu pakai cookies
	}).Handler(middleware.NewAuthMiddleware(router))
	server := http.Server{
		Addr:    ":8080",
		Handler: Handler,
	}
	fmt.Println("Aplikasi Mulai")
	err := server.ListenAndServe()
	helper.PanicError(err)
}
