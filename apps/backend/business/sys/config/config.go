package config

import (
	"strings"
	"time"
)

type Web struct {
	APIHost         string
	AllowedOrigins  []string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type Auth struct {
	Issuer              string
	AccessTokenSecret   string
	RefreshTokenSecret  string
	Audience            string
	AccessTokenExpTime  time.Duration
	RefreshTokenExpTime time.Duration
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
	webReadTimeOut, err := time.ParseDuration(GetEnvString("WEB_READ_TIMEOUT", "10s"))
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

	accessTokenExp, err := time.ParseDuration(GetEnvString("ACCESS_TOKEN_EXPIRATION_TIME", "3600s"))
	if err != nil {
		accessTokenExp = time.Second * 3600
	}

	refreshTokenExp, err := time.ParseDuration(GetEnvString("REFRESH_TOKEN_EXPIRATION_TIME", "2190h"))
	if err != nil {
		refreshTokenExp = time.Hour * 2190
	}

	origins := GetEnvString("ALLOWED_ORIGINS", "http://localhost:3000")
	allowedOrigins := strings.Split(origins, ",")

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
			AllowedOrigins:  allowedOrigins,
			WriteTimeout:    webWriteTimeout,
			ShutdownTimeout: webShutdownTimeout,
			APIHost:         GetEnvString("WEB_API_HOST", "localhost:3001"),
		},
		Auth: Auth{
			AccessTokenExpTime:  accessTokenExp,
			RefreshTokenExpTime: refreshTokenExp,
			Audience:            GetEnvString("AUDIENCE", "localhost:3001"),
			Issuer:              GetEnvString("AUTH_ISSUER", "open-pulse-backend"),
			AccessTokenSecret: GetEnvString(
				"ACCESS_TOKEN_SECRET", "92c3ba3f929dc49a3468c0ff6b7340997c04d522af3a5216f43d71a3b5c97788c64484",
			),
			RefreshTokenSecret: GetEnvString(
				"REFRESH_TOKEN_SECRET", "92c3ba3f929dc49a3468c0ff6b7340997c04d522af3a5216f43d71a3b5c97788c64484",
			),
		},
	}
}
