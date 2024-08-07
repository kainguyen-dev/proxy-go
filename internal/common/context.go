package common

type Role string

const (
	SuperAdmin Role = "AESI_SUPER_ADMIN"
	CloudAdmin Role = "AESI_CLOUD_ADMIN"
	SiteAdmin  Role = "AESI_SITE_ADMIN"
)

type RequestContext struct {
	Username   string
	SiteIDs    []int
	CustomerId *int
	Role       []Role
}

func ToRole(string string) Role {
	switch string {
	case "AESI_SUPER_ADMIN":
		return SuperAdmin
	case "AESI_CLOUD_ADMIN":
		return CloudAdmin
	case "AESI_SITE_ADMIN":
		return SiteAdmin
	default:
		panic(ServerError{Code: 500, Message: "Invalid role " + string})
	}
}

func HasRole(roles []Role, checkRole Role) bool {
	for _, role := range roles {
		if role == checkRole {
			return true
		}
	}
	return false
}
