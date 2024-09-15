package config

import (
	"time"
)

type Web struct {
	APIHost         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type Auth struct {
	Issuer string
	Secret string
}

type DB struct {
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
	Name         string
	User         string
	Host         string
	Password     string
}

type OpenpulseApiConfig struct {
	DB   DB
	Web  Web
	Auth Auth
}

func NewOpenpulseConfig() *OpenpulseApiConfig {
	webReadTimeOut, err := time.ParseDuration(GetEnvString("WEB_READ_TIMEOUT", "5s"))
	if err != nil {
		webReadTimeOut = time.Second * 5
	}

	webIdleTimeout, err := time.ParseDuration(GetEnvString("WEB_IDLE_TIMEOUT", "120s"))
	if err != nil {
		webIdleTimeout = time.Second * 120
	}

	webWriteTimeout, err := time.ParseDuration(GetEnvString("WEB_WRITE_TIMEOUT", "10s"))
	if err != nil {
		webWriteTimeout = time.Second * 10
	}

	webShutdownTimeout, err := time.ParseDuration(GetEnvString("WEB_WRITE_TIMEOUT", "20s"))
	if err != nil {
		webShutdownTimeout = time.Second * 20
	}

	return &OpenpulseApiConfig{
		DB: DB{
			MaxIdleConns: GetEnvInt("DB_MAX_IDLE_CONN", 5),
			MaxOpenConns: GetEnvInt("DB_MAX_OPEN_CONN", 20),
			User:         GetEnvString("DB_USER", "postgres"),
			Name:         GetEnvString("DB_NAME", "openpulse"),
			Host:         GetEnvString("DB_HOST", "localhost"),
			Password:     GetEnvString("DB_PASSWORD", "password"),
			DisableTLS:   GetEnvString("DB_TLS", "disable") == "disable",
		},
		Web: Web{
			ReadTimeout:     webReadTimeOut,
			IdleTimeout:     webIdleTimeout,
			WriteTimeout:    webWriteTimeout,
			ShutdownTimeout: webShutdownTimeout,
			APIHost:         GetEnvString("WEB_API_HOST", "localhost:3000"),
		},
		Auth: Auth{
			Issuer: GetEnvString("AUTH_ISSUER", "open-pulse-backend"),
			Secret: GetEnvString(
				"AUTH_SECRET", "92c3ba3f929dc49a3468c0ff6b7340997c04d522af3a5216f43d71a3b5c97788c64484",
			),
		},
	}
}
