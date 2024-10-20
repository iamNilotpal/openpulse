package users

type AccountStatus string

const (
	AccountActiveInt int = iota + 1
	AccountDeletedInt
)

const (
	AccountActiveString  AccountStatus = "active"
	AccountDeletedString AccountStatus = "deleted"
)

var accountStatus = map[int]AccountStatus{
	AccountActiveInt:  AccountActiveString,
	AccountDeletedInt: AccountDeletedString,
}

var accountStatusReverse = map[AccountStatus]int{
	AccountActiveString:  AccountActiveInt,
	AccountDeletedString: AccountDeletedInt,
}

func ParseStatusString(s AccountStatus) int {
	return accountStatusReverse[s]
}

func ParseStatusInt(v int) AccountStatus {
	return accountStatus[v]
}

type SystemAppearance string

const (
	AppearanceDark   SystemAppearance = "dark"
	AppearanceLight  SystemAppearance = "light"
	AppearanceSystem SystemAppearance = "system"
)

func ParseAppearanceString(s string) SystemAppearance {
	return SystemAppearance(s)
}

func ParseAppearance(s SystemAppearance) string {
	return string(s)
}
