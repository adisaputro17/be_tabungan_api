package domain

const (
	createNasabah = `INSERT INTO nasabah (
		no_rekening,
		nik,
		nama,
		no_hp,
		saldo,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	updateSaldoNasabah = `UPDATE nasabah SET
		saldo = $1,
		updated_at = $2,
		updated_by = $3
	WHERE no_rekening = $4`

	getNasabahByNoRekening = `SELECT
		no_rekening,
		nik,
		nama,
		no_hp,
		saldo,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM nasabah WHERE no_rekening = $1`

	getNasabahByNIKOrHP = `SELECT
		no_rekening,
		nik,
		nama,
		no_hp,
		saldo,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM nasabah WHERE nik = $1 OR no_hp = $2`
)
