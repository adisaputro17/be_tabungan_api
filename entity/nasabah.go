package entity

import "time"

const (
	KodeTransaksiTabung = "C"
	KodeTransaksiTarik  = "D"
)

type Nasabah struct {
	NoRekening string  `json:"no_rekening"`
	NIK        string  `json:"nik"`
	Nama       string  `json:"nama"`
	NoHP       string  `json:"no_hp"`
	Saldo      float64 `json:"saldo"`
	CommonField
}

type CreateNasabahRequest struct {
	NIK  string `json:"nik" validate:"required"`
	Nama string `json:"nama" validate:"required"`
	NoHP string `json:"no_hp" validate:"required"`
}

type UpdateSaldoNasabahRequest struct {
	KodeTransaksi string  `json:"kode_transaksi" validate:"required"`
	NoRekening    string  `json:"no_rekening" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required"`
}

type UpdateSaldoNasabah struct {
	NoRekening string    `json:"no_rekening"`
	Saldo      float64   `json:"saldo"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int64     `json:"updated_by"`
}
