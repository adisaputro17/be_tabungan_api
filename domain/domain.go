package domain

import (
	"be_tabungan_api/entity"
	"context"
	"database/sql"
)

type DomainItf interface {
	CreateNasabah(ctx context.Context, v entity.Nasabah) (entity.Nasabah, error)
	UpdateSaldoNasabah(ctx context.Context, v entity.Mutasi) (entity.Mutasi, error)
	GetNasabahByID(ctx context.Context, noRekening string) (entity.Nasabah, error)
	GetNasabahByNIKOrHP(ctx context.Context, nik string, noHP string) ([]entity.Nasabah, error)

	GetMutasiByNoRek(ctx context.Context, noRek string) ([]entity.Mutasi, error)
}

type domain struct {
	DB  *sql.DB
	Opt Options
}

type Options struct {
}

func InitDomain(
	db *sql.DB,
	opt Options,
) DomainItf {
	return &domain{
		DB:  db,
		Opt: opt,
	}
}
