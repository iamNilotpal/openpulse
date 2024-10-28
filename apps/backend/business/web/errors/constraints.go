package errors

var dbUniqueConstraintsMap = map[string]string{
	"resources_resource_key":     "Resource type already exists.",
	"resources_name_key":         "Resource with same name already exists.",
	"permissions_name_key":       "Permission with same name already exists.",
	"permissions_action_key":     "Permission type already exists.",
	"users_email_key":            "User with same email already exists.",
	"roles_name_key":             "Role with same name already exists.",
	"roles_role_key":             "Role type already exists.",
	"organizations_admin_id_key": "One user can create only one organization.",
}

func GetUniqueConstraint(key string) (string, bool) {
	v, ok := dbUniqueConstraintsMap[key]
	return v, ok
}
