package domain

const (
	createMutasi = `INSERT INTO mutasi (
		mutasi_id,
		no_rekening,
		saldo_before,
		saldo_after,
		kode_transaksi,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	getMutasiByNoRek = `SELECT
		mutasi_id,
		no_rekening,
		saldo_before,
		saldo_after,
		kode_transaksi,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM mutasi WHERE no_rekening = $1`
)
