package usecase

import (
	"be_tabungan_api/domain"
	"be_tabungan_api/entity"
	"context"

	"github.com/go-playground/validator/v10"
)

type UsecaseItf interface {
	CreateNasabah(ctx context.Context, request entity.CreateNasabahRequest) (entity.Nasabah, error)
	UpdateSaldoNasabah(ctx context.Context, request entity.UpdateSaldoNasabahRequest) (entity.Mutasi, error)
	GetNasabahByID(ctx context.Context, noRekening string) (entity.Nasabah, error)
	GetMutasiByNoRek(ctx context.Context, noRekening string) ([]entity.Mutasi, error)
}

type usecase struct {
	Domain   domain.DomainItf
	Validate *validator.Validate
}

func InitUsecase(domain domain.DomainItf, validate *validator.Validate) UsecaseItf {
	return &usecase{
		Domain:   domain,
		Validate: validate,
	}
}
