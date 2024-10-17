package sessions

import "time"

type NewSession struct {
	UserId       int
	Token        string
	UserAgent    string
	IpAddress    string
	DeviceInfo   map[string]any
	LocationInfo map[string]any
}

type Session struct {
	Id             int
	UserId         int
	Token          string
	UserAgent      string
	IsActive       bool
	IpAddress      string
	RevokedAt      string
	DeviceInfo     map[string]any
	LocationInfo   map[string]any
	ExpiresAt      time.Time
	CreatedAt      time.Time
	LastActivityAt time.Time
}
