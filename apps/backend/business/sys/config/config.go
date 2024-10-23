package config

import (
	"strings"
	"time"

	"github.com/iamNilotpal/openpulse/foundation/validate"
)

type Web struct {
	ClientAPIHost   string
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
	Scheme       string
}

type Cache struct {
	DBName   string
	Protocol string
	User     string
	Host     string
	Scheme   string
	Password string
}

type Email struct {
	Username     string
	Password     string
	Issuer       string
	Secret       string
	Audience     string
	EmailExpTime time.Duration
}

type Onboarding struct {
	Issuer       string
	Secret       string
	Audience     string
	TokenExpTime time.Duration
}

type APIConfig struct {
	DB         *DB
	Web        *Web
	Auth       *Auth
	Cache      *Cache
	Email      *Email
	Onboarding *Onboarding
}

func NewAPIConfig() *APIConfig {
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

	origins := GetEnvString("ALLOWED_ORIGINS", "http://localhost:3000")
	allowedOrigins := strings.Split(origins, ",")

	accessTokenExp, err := time.ParseDuration(GetEnvString("AUTH_ACCESS_TOKEN_EXPIRATION_TIME", "3600s"))
	if err != nil {
		accessTokenExp = time.Second * 3600
	}

	refreshTokenExp, err := time.ParseDuration(GetEnvString("AUTH_REFRESH_TOKEN_EXPIRATION_TIME", "2190h"))
	if err != nil {
		refreshTokenExp = time.Hour * 2190
	}

	emailTokenExpTime, err := time.ParseDuration(GetEnvString("EMAIL_TOKEN_EXPIRATION_TIME", "1800s"))
	if err != nil {
		emailTokenExpTime = time.Second * 1800
	}

	onboardingTokenExpTime, err := time.ParseDuration(
		GetEnvString("ONBOARDING_TOKEN_EXPIRATION_TIME", "1800s"),
	)
	if err != nil {
		onboardingTokenExpTime = time.Second * 1800
	}

	return &APIConfig{
		DB: &DB{
			MaxIdleConns: GetEnvInt("DB_MAX_IDLE_CONN", 5),
			MaxOpenConns: GetEnvInt("DB_MAX_OPEN_CONN", 20),
			User:         GetEnvString("DB_USER", "postgres"),
			Name:         GetEnvString("DB_NAME", "openpulse"),
			Host:         GetEnvString("DB_HOST", "localhost"),
			Scheme:       GetEnvString("DB_SCHEME", "postgres"),
			Password:     GetEnvString("DB_PASSWORD", "password"),
			DisableTLS:   GetEnvString("DB_TLS", "disable") == "disable",
		},
		Cache: &Cache{
			DBName:   GetEnvString("CACHE_DB_NAME", "0"),
			Protocol: GetEnvString("CACHE_DB_PROTOCOL", "3"),
			User:     GetEnvString("CACHE_DB_USER", "redis"),
			Scheme:   GetEnvString("CACHE_DB_SCHEME", "redis"),
			Password: GetEnvString("CACHE_DB_PASSWORD", "redis"),
			Host:     GetEnvString("CACHE_DB_HOST", "localhost:6379"),
		},
		Web: &Web{
			ReadTimeout:     webReadTimeOut,
			IdleTimeout:     webIdleTimeout,
			AllowedOrigins:  allowedOrigins,
			WriteTimeout:    webWriteTimeout,
			ShutdownTimeout: webShutdownTimeout,
			APIHost:         GetEnvString("WEB_API_HOST", "localhost:8000"),
			ClientAPIHost:   GetEnvString("WEB_CLIENT_HOST", "http://localhost:5172"),
		},
		Auth: &Auth{
			AccessTokenExpTime:  accessTokenExp,
			RefreshTokenExpTime: refreshTokenExp,
			AccessTokenSecret:   GetEnvString("AUTH_ACCESS_TOKEN_SECRET", ""),
			RefreshTokenSecret:  GetEnvString("AUTH_REFRESH_TOKEN_SECRET", ""),
			Issuer:              GetEnvString("AUTH_ISSUER", "https://api.openpulse.dev"),
			Audience:            GetEnvString("AUTH_AUDIENCE", "https://api.openpulse.dev"),
		},
		Email: &Email{
			EmailExpTime: emailTokenExpTime,
			Secret:       GetEnvString("EMAIL_SECRET", ""),
			Username:     GetEnvString("EMAIL_USERNAME", ""),
			Password:     GetEnvString("EMAIL_PASSWORD", ""),
			Issuer:       GetEnvString("EMAIL_ISSUER", "https://api.openpulse.dev"),
			Audience:     GetEnvString("EMAIL_AUDIENCE", "https://api.openpulse.dev"),
		},
		Onboarding: &Onboarding{
			TokenExpTime: onboardingTokenExpTime,
			Secret:       GetEnvString("ONBOARDING_SECRET", ""),
			Issuer:       GetEnvString("ONBOARDING_ISSUER", "https://api.openpulse.dev"),
			Audience:     GetEnvString("ONBOARDING_AUDIENCE", "https://api.openpulse.dev"),
		},
	}
}

func Validate(cfg APIConfig) error {
	return validate.Check(cfg)
}
