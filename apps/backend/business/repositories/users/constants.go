package users

type AccountStatus string
type SystemAppearance string

var (
	AccountActive    AccountStatus = "active"
	AccountDeleted   AccountStatus = "deleted"
	AccountSuspended AccountStatus = "suspended"
)

var (
	AppearanceLight  SystemAppearance = "light"
	AppearanceDark   SystemAppearance = "dark"
	AppearanceSystem SystemAppearance = "suspended"
)

func FromAccountStatus(s AccountStatus) string {
	return string(s)
}

func ToAccountStatus(s string) AccountStatus {
	return AccountStatus(s)
}

func ToSystemAppearance(s string) SystemAppearance {
	return SystemAppearance(s)
}

func FromSystemAppearance(s SystemAppearance) string {
	return string(s)
}
