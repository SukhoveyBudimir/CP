package model

type Player struct {
	ID           string `bson,json:"id"`
	Name         string `bson,json:"name"`
	Games        int32  `bson,json:"game"`
	Wins         int32  `bson,json:"wins"`
	Password     string `bson,json:"password"`
	RefreshToken string `bson,json:"refreshToken"`
}

type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL"`
	JwtKey        []byte `env:"JWT-KEY" `
}
