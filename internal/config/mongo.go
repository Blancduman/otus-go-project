package config

type Mongo struct {
	DSN   string `envconfig:"MONGO_DSN" default:"mongodb://127.0.0.1:27017/twirler"`
	DB    string `envconfig:"MONGO_DB"`
	DSNRO string `envconfig:"MONGO_DSN_RO" default:"mongodb://127.0.0.1:27017"`
	DBRO  string `envconfig:"MONGO_DB_RO"`
}
