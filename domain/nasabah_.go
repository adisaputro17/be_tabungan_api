package domain

import (
	"be_tabungan_api/entity"
	"context"
	"database/sql"
	"log"
)

func (d *domain) CreateNasabah(ctx context.Context, v entity.Nasabah) (entity.Nasabah, error) {
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, err
	}

	tx, err = d.sqlCreateNasabah(tx, v)
	if err != nil {
		return v, err
	}

	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		} else {
			tx.Commit()
		}
	}()

	return v, nil
}

func (d *domain) UpdateSaldoNasabah(ctx context.Context, v entity.Mutasi) (entity.Mutasi, error) {
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, err
	}

	tx, err = d.sqlCreateMutasi(tx, v)
	if err != nil {
		return v, err
	}
	tx, err = d.sqlUpdateSaldoNasabah(tx, entity.UpdateSaldoNasabah{
		NoRekening: v.NoRekening,
		Saldo:      v.SaldoAfter,
		UpdatedAt:  v.UpdatedAt,
		UpdatedBy:  v.UpdatedBy,
	})
	if err != nil {
		return v, err
	}

	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		} else {
			tx.Commit()
		}
	}()

	return v, nil
}

func (d *domain) GetNasabahByID(ctx context.Context, noRekening string) (entity.Nasabah, error) {
	return d.sqlGetNasabahByID(ctx, noRekening)
}

func (d *domain) GetNasabahByNIKOrHP(ctx context.Context, nik string, noHP string) ([]entity.Nasabah, error) {
	return d.sqlGetNasabahByNIKOrHP(ctx, nik, noHP)
}
