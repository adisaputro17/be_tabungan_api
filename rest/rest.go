package rest

import (
	"be_tabungan_api/usecase"

	"github.com/labstack/echo/v4"
)

type REST interface {
	CreateNasabah(c echo.Context) error
	TambahSaldo(c echo.Context) error
	TarikSaldo(c echo.Context) error
	GetNasabahByID(c echo.Context) error
	GetMutasiByNoRek(c echo.Context) error
}

type rest struct {
	Usecase usecase.UsecaseItf
}

func Init(usecase usecase.UsecaseItf) REST {
	return &rest{
		Usecase: usecase,
	}
}
