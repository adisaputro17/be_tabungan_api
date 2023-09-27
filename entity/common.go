package entity

import "time"

type CommonField struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int64     `json:"updated_by"`
}

type AppConfig struct {
	AppPort  string
	DBConfig DatabaseConfig
}

type DatabaseConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

type HTTPErrorResponse struct {
	Remark string `json:"remark"`
}
