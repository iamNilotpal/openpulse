package permissions

import (
	"net/http"
	"slices"

	"github.com/iamNilotpal/openpulse/business/web/errors"
)

type PermissionAction string

const (
	PermissionStringView   PermissionAction = "view"
	PermissionStringCreate PermissionAction = "create"
	PermissionStringUpdate PermissionAction = "update"
	PermissionStringDelete PermissionAction = "delete"
	PermissionStringManage PermissionAction = "manage"
)

const (
	PermissionIntView int = iota + 1
	PermissionIntCreate
	PermissionIntUpdate
	PermissionIntDelete
	PermissionIntManage
)

var permissions = []PermissionAction{
	PermissionStringCreate,
	PermissionStringDelete,
	PermissionStringManage,
	PermissionStringUpdate,
	PermissionStringView,
}

var permissionsMap = map[PermissionAction]int{
	PermissionStringView:   PermissionIntView,
	PermissionStringCreate: PermissionIntCreate,
	PermissionStringUpdate: PermissionIntUpdate,
	PermissionStringDelete: PermissionIntDelete,
	PermissionStringManage: PermissionIntManage,
}

var permissionsReverseMap = map[int]PermissionAction{
	PermissionIntView:   PermissionStringView,
	PermissionIntCreate: PermissionStringCreate,
	PermissionIntUpdate: PermissionStringUpdate,
	PermissionIntDelete: PermissionStringDelete,
	PermissionIntManage: PermissionStringManage,
}

func ParseActionString(action string) (PermissionAction, error) {
	if contains := slices.Contains(permissions, PermissionAction(action)); contains {
		return PermissionAction(action), nil
	}

	return "", errors.NewRequestError(
		"Invalid permission action.",
		http.StatusBadRequest,
		errors.BadRequest,
	)
}

func ParseAction(action PermissionAction) int {
	return permissionsMap[action]
}

func ParseActionInt(action int) PermissionAction {
	return permissionsReverseMap[action]
}
