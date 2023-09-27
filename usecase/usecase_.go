package usecase

import (
	"be_tabungan_api/entity"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (u *usecase) CreateNasabah(ctx context.Context, request entity.CreateNasabahRequest) (entity.Nasabah, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.Nasabah{}, err
	}

	currNasabah, err := u.Domain.GetNasabahByNIKOrHP(ctx, request.NIK, request.NoHP)
	if err != nil {
		return entity.Nasabah{}, err
	}

	if len(currNasabah) > 0 {
		return entity.Nasabah{}, errors.New("data already exists")
	}

	timeNowUTC := time.Now().UTC()
	commonField := entity.CommonField{
		CreatedAt: timeNowUTC,
		CreatedBy: 0,
		UpdatedAt: timeNowUTC,
		UpdatedBy: 0,
	}

	nasabahData := entity.Nasabah{
		NoRekening:  fmt.Sprint(timeNowUTC.Unix()),
		NIK:         request.NIK,
		Nama:        request.Nama,
		NoHP:        request.NoHP,
		CommonField: commonField,
	}

	nasabah, err := u.Domain.CreateNasabah(ctx, nasabahData)
	if err != nil {
		return entity.Nasabah{}, err
	}

	return nasabah, nil
}

func (u *usecase) UpdateSaldoNasabah(ctx context.Context, request entity.UpdateSaldoNasabahRequest) (entity.Mutasi, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.Mutasi{}, err
	}

	nasabah, err := u.Domain.GetNasabahByID(ctx, request.NoRekening)
	if err != nil {
		return entity.Mutasi{}, err
	}

	saldoBefore := nasabah.Saldo
	switch request.KodeTransaksi {
	case entity.KodeTransaksiTarik:
		nasabah.Saldo = nasabah.Saldo - request.Nominal
		if nasabah.Saldo < 0 {
			return entity.Mutasi{}, errors.New("balance is not enough")
		}
	case entity.KodeTransaksiTabung:
		nasabah.Saldo = nasabah.Saldo + request.Nominal
	}

	timeNowUTC := time.Now().UTC()
	commonField := entity.CommonField{
		CreatedAt: timeNowUTC,
		CreatedBy: 0,
		UpdatedAt: timeNowUTC,
		UpdatedBy: 0,
	}

	mutasi, err := u.Domain.UpdateSaldoNasabah(ctx, entity.Mutasi{
		MutasiID:      uuid.NewString(),
		NoRekening:    nasabah.NoRekening,
		SaldoBefore:   saldoBefore,
		SaldoAfter:    nasabah.Saldo,
		KodeTransaksi: request.KodeTransaksi,
		CommonField:   commonField,
	})
	if err != nil {
		return entity.Mutasi{}, err
	}

	return mutasi, nil
}

func (u *usecase) GetNasabahByID(ctx context.Context, noRekening string) (entity.Nasabah, error) {
	return u.Domain.GetNasabahByID(ctx, noRekening)
}

func (u *usecase) GetMutasiByNoRek(ctx context.Context, noRekening string) ([]entity.Mutasi, error) {
	return u.Domain.GetMutasiByNoRek(ctx, noRekening)
}
