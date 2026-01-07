package config

type Config struct {
	Database DatabaseConfig `mapstructure:",squash"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Name     string `mapstructure:"DB_NAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Port     int    `mapstructure:"DB_PORT"`
}
