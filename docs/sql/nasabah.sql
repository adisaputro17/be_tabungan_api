CREATE TABLE IF NOT EXISTS nasabah
(
    no_rekening varchar(100) NOT NULL,
    nik varchar(20) NOT NULL,
    nama text NOT NULL,
    no_hp varchar(20) NOT NULL,
    saldo decimal(15,2) NOT NULL,
    created_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    created_by bigint NOT NULL,
    updated_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_by bigint NOT NULL,
    PRIMARY KEY (no_rekening)
)