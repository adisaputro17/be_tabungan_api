package domain

import (
	"be_tabungan_api/entity"
	"context"
)

func (d *domain) GetMutasiByNoRek(ctx context.Context, noRek string) ([]entity.Mutasi, error) {
	return d.sqlGetMutasiByNoRek(ctx, noRek)
}
