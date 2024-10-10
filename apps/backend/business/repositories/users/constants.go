package users

type AccountStatus string

const (
	AccountActiveInt int = iota + 1
	AccountDeletedInt
	AccountSuspendedInt
	AccountDeactivatedInt
)

const (
	AccountActiveString      AccountStatus = "active"
	AccountDeletedString     AccountStatus = "deleted"
	AccountSuspendedString   AccountStatus = "suspended"
	AccountDeactivatedString AccountStatus = "deactivated"
)

var accountStatus = map[int]AccountStatus{
	AccountActiveInt:      AccountActiveString,
	AccountDeletedInt:     AccountDeletedString,
	AccountSuspendedInt:   AccountSuspendedString,
	AccountDeactivatedInt: AccountDeactivatedString,
}

var accountStatusReverse = map[AccountStatus]int{
	AccountActiveString:      AccountActiveInt,
	AccountDeletedString:     AccountDeletedInt,
	AccountSuspendedString:   AccountSuspendedInt,
	AccountDeactivatedString: AccountDeactivatedInt,
}

func ParseStatusString(s AccountStatus) int {
	return accountStatusReverse[s]
}

func ParseStatusInt(v int) AccountStatus {
	return accountStatus[v]
}

type SystemAppearance string

const (
	AppearanceLight  SystemAppearance = "light"
	AppearanceDark   SystemAppearance = "dark"
	AppearanceSystem SystemAppearance = "suspended"
)

func ParseAppearanceString(s string) SystemAppearance {
	return SystemAppearance(s)
}

func ParseAppearance(s SystemAppearance) string {
	return string(s)
}
