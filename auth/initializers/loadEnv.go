package initializers

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost         string `mapstructure:"MYSQL_HOST"`
	DBUserName     string `mapstructure:"MYSQL_USER"`
	DBUserPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName         string `mapstructure:"MYSQL_DATABASE"`
	DBPort         string `mapstructure:"MYSQL_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	RedisUri     string `mapstructure:"REDIS_URL"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	// err1 := godotenv.Load()
	// if err1 != nil {
	// 	DBHost1, _ := os.LookupEnv("MYSQL_HOST")
	// 	log.Fatal("Error loading .env file", DBHost1)
	// }

	DBHost, _ := os.LookupEnv("MYSQL_HOST")
	DBUserName, _ := os.LookupEnv("MYSQL_USER")
	DBUserPassword, _ := os.LookupEnv("MYSQL_PASSWORD")
	DBName, _ := os.LookupEnv("MYSQL_DATABASE")
	DBPort, _ := os.LookupEnv("MYSQL_PORT")
	ServerPort, _ := os.LookupEnv("PORT")
	ClientOrigin, _ := os.LookupEnv("CLIENT_ORIGIN")
	RedisUri, _ := os.LookupEnv("REDIS_URL")
	AccessTokenPrivateKey, _ := os.LookupEnv("ACCESS_TOKEN_PRIVATE_KEY")
	AccessTokenPublicKey, _ := os.LookupEnv("ACCESS_TOKEN_PUBLIC_KEY")
	RefreshTokenPrivateKey, _ := os.LookupEnv("REFRESH_TOKEN_PRIVATE_KEY")
	RefreshTokenPublicKey, _ := os.LookupEnv("REFRESH_TOKEN_PUBLIC_KEY")
	AccessTokenExpiresIn_str, _ := os.LookupEnv("ACCESS_TOKEN_EXPIRED_IN")
	RefreshTokenExpiresIn_str, _ := os.LookupEnv("REFRESH_TOKEN_EXPIRED_IN")
	AccessTokenMaxAge_str, _ := os.LookupEnv("ACCESS_TOKEN_MAXAGE")
	RefreshTokenMaxAge_str, _ := os.LookupEnv("REFRESH_TOKEN_MAXAGE")

	AccessTokenExpiresIn, _ := time.ParseDuration(AccessTokenExpiresIn_str)
	RefreshTokenExpiresIn, _ := time.ParseDuration(RefreshTokenExpiresIn_str)
	AccessTokenMaxAge, _ := strconv.Atoi(AccessTokenMaxAge_str)
	RefreshTokenMaxAge, _ := strconv.Atoi(RefreshTokenMaxAge_str)

	return Config{
		DBHost,
		DBUserName,
		DBUserPassword,
		DBName,
		DBPort,
		ServerPort,
		ClientOrigin,
		RedisUri,
		AccessTokenPrivateKey,
		AccessTokenPublicKey,
		RefreshTokenPrivateKey,
		RefreshTokenPublicKey,
		AccessTokenExpiresIn,
		RefreshTokenExpiresIn,
		AccessTokenMaxAge,
		RefreshTokenMaxAge,
	}, nil
}
