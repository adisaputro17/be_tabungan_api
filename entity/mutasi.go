package entity

type Mutasi struct {
	MutasiID      string  `json:"mutasi_id"`
	NoRekening    string  `json:"no_rekening"`
	SaldoBefore   float64 `json:"saldo_before"`
	SaldoAfter    float64 `json:"saldo_after"`
	KodeTransaksi string  `json:"kode_transaksi"`
	CommonField
}
