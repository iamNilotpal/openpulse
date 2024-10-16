package permissions

import (
	"time"
)

type NewPermission struct {
	Name        string
	Description string
	Action      PermissionAction
}

type Permission struct {
	Id          int
	Name        string
	Description string
	Action      PermissionAction
	CreateAt    time.Time
	UpdatedAt   time.Time
}

type PermissionAccessConfig struct {
	Id     int
	Action PermissionAction
}
