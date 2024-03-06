CREATE TABLE akun (
    no_rekening VARCHAR(20) PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    nik VARCHAR(20) NOT NULL UNIQUE,
    no_hp VARCHAR(20) NOT NULL UNIQUE,
    pin VARCHAR(255) NOT NULL,
    saldo INT NOT NULL DEFAULT 0
);
