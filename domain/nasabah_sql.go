package domain

import (
	"be_tabungan_api/entity"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func (d *domain) sqlCreateNasabah(tx *sql.Tx, v entity.Nasabah) (*sql.Tx, error) {
	_, err := tx.Exec(createNasabah,
		v.NoRekening,
		v.NIK,
		v.Nama,
		v.NoHP,
		v.Saldo,
		v.CreatedAt,
		v.CreatedBy,
		v.UpdatedAt,
		v.UpdatedBy,
	)
	if err != nil {
		return tx, err
	}

	return tx, nil
}

func (d *domain) sqlUpdateSaldoNasabah(tx *sql.Tx, v entity.UpdateSaldoNasabah) (*sql.Tx, error) {
	_, err := tx.Exec(updateSaldoNasabah,
		v.Saldo,
		v.UpdatedAt,
		v.UpdatedBy,
		// where
		v.NoRekening,
	)
	if err != nil {
		return tx, err
	}

	return tx, err
}

func (d *domain) sqlGetNasabahByID(ctx context.Context, noRekening string) (entity.Nasabah, error) {
	result := entity.Nasabah{}
	row := d.DB.QueryRowContext(ctx, getNasabahByNoRekening, noRekening)

	err := row.Scan(
		&result.NoRekening,
		&result.NIK,
		&result.Nama,
		&result.NoHP,
		&result.Saldo,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (d *domain) sqlGetNasabahByNIKOrHP(ctx context.Context, nik string, noHP string) ([]entity.Nasabah, error) {
	results := []entity.Nasabah{}
	rows, err := d.DB.QueryContext(ctx, getNasabahByNIKOrHP, nik, noHP)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		data := entity.Nasabah{}
		if err := rows.Scan(
			&data.NoRekening,
			&data.NIK,
			&data.Nama,
			&data.NoHP,
			&data.Saldo,
			&data.CreatedAt,
			&data.CreatedBy,
			&data.UpdatedAt,
			&data.UpdatedBy,
		); err != nil {
			return results, err
		}
		results = append(results, data)
	}

	return results, nil
}
