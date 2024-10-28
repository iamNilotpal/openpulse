package users_handler

type User struct {
	Id              int             `json:"id"`
	FirstName       string          `json:"firstName"`
	LastName        string          `json:"lastName"`
	Email           string          `json:"email"`
	Phone           string          `json:"phoneNumber"`
	AvatarUrl       string          `json:"avatarURL"`
	Designation     string          `json:"designation"`
	IsEmailVerified bool            `json:"isEmailVerified"`
	Role            Role            `json:"role"`
	AccountStatus   string          `json:"accountStatus"`
	OAuthAccounts   []OAuthAccount  `json:"oauthAccounts"`
	Team            *Team           `json:"team"`
	AccessControl   []AccessControl `json:"accessControl"`
	CreatedAt       string          `json:"createdAt"`
	UpdatedAt       string          `json:"updatedAt"`
}

type Team struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logoURL"`
}

type Role struct {
	Id           int    `json:"id"`
	IsSystemRole bool   `json:"isSystemRole"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Role         string `json:"role"`
}

type OAuthAccount struct {
	Id         int    `json:"id"`
	Provider   string `json:"provider"`
	ExternalId string `json:"externalId"`
	Scope      string `json:"scope"`
	Metadata   string `json:"metadata"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type Resource struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
}

type Permission struct {
	Id          int    `json:"id"`
	Enabled     bool   `json:"enabled"`
	Name        string `json:"Name"`
	Description string `json:"description"`
	Action      string `json:"action"`
}

type AccessControl struct {
	Resource   Resource   `json:"resource"`
	Permission Permission `json:"permission"`
}
