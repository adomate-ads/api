package auth

import "github.com/adomate-ads/api/models"

type Role struct {
	Name        string
	RolesCanAdd []string
}

var Groups = map[string][]string{
	"super-admin": {
		"super-admin",
	},
	"support": {
		"support-ticket",
		"support-billing",
	},
	"admin": {
		"owner",
		"admin",
	},
	"user": {
		"user",
	},
}

var Roles = map[string]Role{
	"super-admin": {
		Name: "super-admin",
		RolesCanAdd: []string{
			"super-admin",
		},
	},
	"support-ticket": {
		Name: "support-ticket",
		RolesCanAdd: []string{
			"super-admin",
		},
	},
	"support-billing": {
		Name: "support-billing",
		RolesCanAdd: []string{
			"super-admin",
		},
	},
	"owner": {
		Name: "owner",
		RolesCanAdd: []string{
			"super-admin",
		},
	},
	"admin": {
		Name: "admin",
		RolesCanAdd: []string{
			"super-admin",
			"support-ticket",
			"support-billing",
			"owner",
			"admin",
		},
	},
	"user": {
		Name: "user",
		RolesCanAdd: []string{
			"super-admin",
			"support-ticket",
			"support-billing",
			"owner",
			"admin",
		},
	},
}

func CanUserModifyRole(user *models.User, role string) bool {
	if _, ok := Roles[role]; !ok {
		return false
	}
	return HasRoleList(user, Roles[role].RolesCanAdd)
}

func InGroup(user *models.User, group string) bool {
	// super admins are always in group, no matter the group.
	if group != "super-admin" && InGroup(user, "super-admin") {
		return true
	}

	// Group does not exist
	if _, ok := Groups[group]; !ok {
		return false
	}

	has := HasRoleList(user, Groups[group])
	return has
}

func HasRoleList(user *models.User, roles []string) bool {
	for _, r := range roles {
		if HasRole(user, r) {
			return true
		}
	}
	return false
}

func HasRole(user *models.User, role string) bool {
	if user.Role == role {
		return true
	}
	return false
}
