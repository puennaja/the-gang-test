package constant

const (
	VisitorRole   string = "visitor"
	UserRole      string = "user"
	ModeratorRole string = "moderator"
	AdminRole     string = "admin"
)

const (
	AuthHeaderKey string = "x-user-id"
	AuthNameKey   string = "x-user-name"
	AuthRoleKey   string = "x-user-role"
	AuthLevelKey  string = "x-user-level"

	AuthLevel0 string = "LEVEL0"
	AuthLevel1 string = "LEVEL1"
	AuthLevel2 string = "LEVEL2"
	AuthLevel3 string = "LEVEL3"
)
