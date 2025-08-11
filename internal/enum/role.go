package enum

import "strings"

type Role string

const (
	ADMIN     Role = "ADMIN"
	INVENTORY Role = "INVENTORY"
	SUPPORT   Role = "SUPPORT"
	GUEST     Role = "GUEST"
)

func RolesToString(roles []Role) string {
	strRoles := make([]string, len(roles))
	for i, r := range roles {
		strRoles[i] = string(r)
	}
	return strings.Join(strRoles, ",")
}

func StringToRoles(r string) []Role {
	var enumRoles []Role
	roles := strings.Split(r, ",")
	for _, roleStr := range roles {
		roleStr = strings.TrimSpace(roleStr)
		if roleStr == "" {
			continue
		}
		enumRoles = append(enumRoles, Role(roleStr))
	}
	return enumRoles
}

func (r Role) IsValid() bool {
	switch r {
	case ADMIN, INVENTORY, SUPPORT, GUEST:
		return true
	}
	return false
}
