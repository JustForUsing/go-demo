package model

const (
	RoleTypeSystem = 1
	RoleTypeCustom = 2

	RoleAdmin      = "admin"
	RoleTeamLeader = "team leader"
	RoleNormalUser = "normal user"

	RoleNameMaxLength = 128
	RoleDescMaxLength = 512
)

type Role struct {
	ID   int64  `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:128;not null;uniqueIndex"`
	Type int8   `gorm:"type:int;not null;comment:角色类型: 1:系统角色, 2:自定义角色"`
	Desc string `gorm:"size:512"`
}

func (Role) TableName() string {
	return "roles"
}

var SystemRoles = []struct {
	Name string
	Desc string
}{
	{Name: RoleAdmin, Desc: "system administrator"},
	{Name: RoleTeamLeader, Desc: "system team leader"},
	{Name: RoleNormalUser, Desc: "system normal user"},
}
