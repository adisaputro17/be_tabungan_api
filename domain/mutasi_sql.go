package domain

import (
	"be_tabungan_api/entity"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func (d *domain) sqlCreateMutasi(tx *sql.Tx, v entity.Mutasi) (*sql.Tx, error) {
	_, err := tx.Exec(createMutasi,
		v.MutasiID,
		v.NoRekening,
		v.SaldoBefore,
		v.SaldoAfter,
		v.KodeTransaksi,
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

func (d *domain) sqlGetMutasiByNoRek(ctx context.Context, noRek string) ([]entity.Mutasi, error) {
	results := []entity.Mutasi{}
	rows, err := d.DB.QueryContext(ctx, getMutasiByNoRek, noRek)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		data := entity.Mutasi{}
		if err := rows.Scan(
			&data.MutasiID,
			&data.NoRekening,
			&data.SaldoBefore,
			&data.SaldoAfter,
			&data.KodeTransaksi,
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
