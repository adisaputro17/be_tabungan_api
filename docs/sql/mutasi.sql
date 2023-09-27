CREATE TABLE IF NOT EXISTS mutasi
(
    mutasi_id varchar(100) NOT NULL,
    no_rekening varchar(100) NOT NULL,
    saldo_before decimal(15,2) NOT NULL,
    saldo_after decimal(15,2) NOT NULL,
    kode_transaksi varchar(10) NOT NULL,
    created_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    created_by bigint NOT NULL,
    updated_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_by bigint NOT NULL,
    PRIMARY KEY (mutasi_id)
)