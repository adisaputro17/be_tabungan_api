package main

import (
	"be_tabungan_api/app"
	"be_tabungan_api/domain"
	"be_tabungan_api/entity"
	"be_tabungan_api/rest"
	"be_tabungan_api/usecase"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appConfig := entity.AppConfig{
		AppPort: os.Getenv("APP_PORT"),
		DBConfig: entity.DatabaseConfig{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBName:     os.Getenv("DB_NAME"),
		},
	}

	validate := validator.New()
	db, err := app.NewDB(appConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	opt := domain.Options{}

	domain := domain.InitDomain(db, opt)
	usecase := usecase.InitUsecase(domain, validate)
	rest := rest.Init(usecase)

	e := echo.New()
	e.POST("/daftar", rest.CreateNasabah)
	e.POST("/tabung", rest.TambahSaldo)
	e.POST("/tarik", rest.TarikSaldo)
	e.GET("/saldo/:no_rekening", rest.GetNasabahByID)
	e.GET("/mutasi/:no_rekening", rest.GetMutasiByNoRek)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appConfig.AppPort)))
}
