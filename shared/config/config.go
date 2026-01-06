package config

import (
	"crypto-platform/shared/vault"

	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

type Exchange struct {
	APIKey     string
	APISecret  string
	Passphrase string
}

type Vault struct {
	Address   string `mapstructure:"address"`
	Token     string `mapstructure:"token"`
	MountPath string `mapstructure:"mount_path"`
}

type Gateway struct {
	Port string `mapstructure:"port"`
}

type Exchanges map[string]Exchange

type Service struct {
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	GRPCPort string `mapstructure:"grpc_port"`
}

type Config struct {
	Vault           Vault   `mapstructure:"vault"`
	CryptoService   Service `mapstructure:"crypto-service"`
	CurrencyService Service `mapstructure:"currency-service"`
	Postgres        Postgres
	Gateway         Gateway `mapstructure:"gateway"`
	Exchanges       Exchanges
}

func Load() (*Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("shared/config/")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	vaultClient, err := vault.NewClient(vault.Config{
		Address:   config.Vault.Address,
		Token:     config.Vault.Token,
		MountPath: config.Vault.MountPath,
	})
	if err != nil {
		return nil, err
	}

	postgresConfig, err := GetPostgresCredentials(vaultClient)
	if err != nil {
		return nil, err
	}

	binance, err := GetExchangeCredentials(vaultClient, "exchange/binance")
	if err != nil {
		return nil, err
	}

	kucoin, err := GetExchangeCredentials(vaultClient, "exchange/kucoin")
	if err != nil {
		return nil, err
	}

	config.Postgres = postgresConfig
	config.Exchanges = Exchanges{
		"binance": binance,
		"kucoin":  kucoin,
	}

	return &config, nil
}

func GetPostgresCredentials(vaultClient *vault.Client) (Postgres, error) {
	data, err := vaultClient.GetSecret("database/postgres")
	if err != nil {
		return Postgres{}, err
	}

	return Postgres{
		Host:     GetValue[string](data, "host"),
		Port:     GetValue[string](data, "port"),
		Username: GetValue[string](data, "username"),
		Password: GetValue[string](data, "password"),
		Database: GetValue[string](data, "database"),
		SSLMode:  GetValue[string](data, "ssl_mode"),
	}, nil
}

func GetExchangeCredentials(vaultClient *vault.Client, path string) (Exchange, error) {
	data, err := vaultClient.GetSecret(path)
	if err != nil {
		return Exchange{}, err
	}

	return Exchange{
		APIKey:     GetValue[string](data, "api_key"),
		APISecret:  GetValue[string](data, "api_secret"),
		Passphrase: GetValue[string](data, "passphrase"),
	}, nil
}

func GetValue[T any](data map[string]any, key string) T {
	if val, ok := data[key].(T); ok {
		return val
	}
	var zero T
	return zero
}
