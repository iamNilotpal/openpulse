package permissions

type PermissionAction string

var (
	ViewPermissionAction   PermissionAction = "view"
	CreatePermissionAction PermissionAction = "create"
	UpdatePermissionAction PermissionAction = "update"
	DeletePermissionAction PermissionAction = "delete"
	ManagePermissionAction PermissionAction = "manage"
)

func FromPermissionAction(action PermissionAction) string {
	return string(action)
}

func ToPermissionAction(action string) PermissionAction {
	return PermissionAction(action)
}
