package model

type Product struct {
	ID       string  `bson,json:"id"`
	NAME     string  `bson,json:"name"`
	PRICE    float32 `bson,json:"price"`
	QUANTITY int32   `bson,json:"quantity"`
}

type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL"`
	MongoDBURL    string `env:"MONGO_DB_URL"`
}
