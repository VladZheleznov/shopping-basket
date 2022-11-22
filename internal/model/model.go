package model

type Product struct {
	ID       string  `bson,json:"id"`
	Name     string  `bson,json:"name"`
	Price    float32 `bson,json:"price"`
	Quantity int32   `bson,json:"quantity"`
}

type Config struct {
	PostgresDBURL string `env:"POSTGRES_DB_URL"`
}
