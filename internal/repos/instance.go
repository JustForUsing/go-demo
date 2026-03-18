package repos

import "go.uber.org/fx"

var Instance RepoContext

type RepoContext struct {
	fx.In
	User        *UserRepo
	Role        *RoleRepo
	UserRole    *UserRoleRepo
	Audit       *AuditRepos
	Team        *TeamRepos
	Project     *ProjectRepo
	TeamUser    *TeamUserRepo
	ProjectUser *ProjectUserRepo
}
