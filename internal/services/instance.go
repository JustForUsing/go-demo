package services

import "go.uber.org/fx"

var Instance ServiceContext

type ServiceContext struct {
	fx.In
	User        *UserService
	Role        *RoleService
	UserRole    *UserRoleService
	Audit       *AuditService
	Team        *TeamService
	Project     *ProjectService
	TeamUser    *TeamUserService
	ProjectUser *ProjectUserService
}
