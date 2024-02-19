package configs

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var instance Config

func Load() *Config {
	if err := env.Parse(&instance); err != nil {
		panic(err)
	}

	return &instance
}

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`

	Server   Server
	Logger   Logger
	Postgres Postgres
	Casbin   Casbin
	JWT      JWT
	Setup    Setup
	Telesign Telesign
	Files    Files
	Admin_RPC Admin_RPC 
	Auth_RPC Auth_RPC
	HTTP_Port HTTP_Port
}

type (
	Server struct {
		Environment       string `env:"SERVER_ENVIRONMENT"`
		Port              uint16 `env:"ADMIN_PORT"`
		MaxConnectionIdle uint16 `env:"SERVER_MAX_CONNECTION_IDLE"`
		Timeout           uint16 `env:"SERVER_TIMEOUT"`
		Time              uint16 `env:"SERVER_TIME"`
		MaxConnectionAge  uint16 `env:"SERVER_MAX_CONNECTION_AGE"`
	}

	HTTP_Port struct{
		HTTP_Port string `env:"HTTP_PORT"`
	}

	Setup struct {
		EmployeeName     string `env:"SETUP_EMPLOYEE_NAME"`
		EmployeeLastName string `env:"SETUP_EMPLOYEE_LAST_NAME"`
		EmployeePhone    string `env:"SETUP_EMPLOYEE_PHONE"`
		EmployeePassword string `env:"SETUP_EMPLOYEE_PASSWORD"`
	}

	Logger struct {
		Level    string `env:"LOGGER_LEVEL"`
		Encoding string `env:"LOGGER_ENCODING"`
	}

	/* trunk-ignore(golangci-lint/goimports) */
	

	Postgres struct {
		Port     uint16 `env:"POSTGRES_PORT"`
		Host     string `env:"POSTGRES_HOST"`
		Password string `env:"POSTGRES_PASSWORD"`
		User     string `env:"POSTGRES_USER"`
		Database string `env:"POSTGRES_DATABASE"`
		PoolMax  int    `env:"POSTGRES_POOL_MAX"`
	}

	Casbin struct {
		ConfigPath string `env:"CASBIN_CONFIG_PATH_ADMIN"`
		Name       string `env:"CASBIN_NAME_USER"`
		PolicyFile string `env:"CASBIN_POLICY_FILE"`
	}

	JWT struct {
		SecretKeyExpireMinutes   uint16 `env:"JWT_ADMIN_SECRET_KEY_EXPIRE_MINUTES"`
		SecretKey                string `env:"JWT_ADMIN_SECRET_KEY"`
		RefreshKeyExpireHours    uint16 `env:"JWT_ADMIN_REFRESH_KEY_EXPIRE_HOURS"`
		ClientRefreshExpireHours uint16 `env:"JWT_CLIENT_REFRESH_EXPIRE_HOURS"`
		RefreshKey               string `env:"JWT_ADMIN_REFRESH_KEY"`
	}

	Telesign struct {
		Key string `env:"TELESIGN_KEY"`
	}

	Files struct {
		ParcelsReportFilePath string `env:"PARCELS_EXCEL_FILE_PATH"`
	}

	Admin_RPC struct{
		Hosts string `env:"ADMIN_RPC_HOST"`
		Port string `env:"ADMIN_RPC_PORT"`
	}
	Auth_RPC struct{
		Hosts string `env:"AUTH_RPC_HOST"`
		Port string `env:"AUTH_RPC_PORT"`
	}
)