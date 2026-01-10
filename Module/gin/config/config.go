package config

type ServerConfig struct {
	Port  string `yaml:"port"`
	DBDSN string `yaml:"db_dsn"`
}
