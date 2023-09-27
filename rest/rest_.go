package rest

import (
	"be_tabungan_api/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (rest *rest) CreateNasabah(c echo.Context) error {
	req := new(entity.CreateNasabahRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	nasabah, err := rest.Usecase.CreateNasabah(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, nasabah)
}

func (rest *rest) TambahSaldo(c echo.Context) error {
	req := new(entity.UpdateSaldoNasabahRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	req.KodeTransaksi = entity.KodeTransaksiTabung
	mutasi, err := rest.Usecase.UpdateSaldoNasabah(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, mutasi)
}

func (rest *rest) TarikSaldo(c echo.Context) error {
	req := new(entity.UpdateSaldoNasabahRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	req.KodeTransaksi = entity.KodeTransaksiTarik
	mutasi, err := rest.Usecase.UpdateSaldoNasabah(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, mutasi)
}

func (rest *rest) GetNasabahByID(c echo.Context) error {
	noRekening := c.Param("no_rekening")

	nasabah, err := rest.Usecase.GetNasabahByID(c.Request().Context(), noRekening)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, nasabah)
}

func (rest *rest) GetMutasiByNoRek(c echo.Context) error {
	noRekening := c.Param("no_rekening")

	mutasi, err := rest.Usecase.GetMutasiByNoRek(c.Request().Context(), noRekening)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, mutasi)
}
