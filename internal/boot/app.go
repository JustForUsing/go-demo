package boot

import (
	"database/sql"
	"fmt"
	"go.uber.org/fx"
	"item-manager-new/internal/api/handler"
	"item-manager-new/internal/auth"
	"item-manager-new/internal/pkg/database"
	"item-manager-new/internal/pkg/engine"
	"item-manager-new/internal/pkg/global"
	"item-manager-new/internal/repos"
	"item-manager-new/internal/services"

	"net/http"
)

func Run() {
	container := fx.New(
		//组件注入
		fx.Provide(
			engine.New,
			engine.NewHttp,
			database.New,
			auth.NewInMemoryStore,
			func(store *auth.InMemoryStore) auth.Store {
				return store
			},
			auth.NewStoreManager,
		),
		//Repo注入
		fx.Provide(
			repos.NewUserRepo,
			repos.NewRoleRepo,
			repos.NewUserRoleRepo,
			repos.NewAuditRepos,
			repos.NewTeamRepos,
			repos.NewProjectRepo,
			repos.NewTeamUserRepo,
			repos.NewProjectUserRepo,
		),
		//Service注入
		fx.Provide(
			services.NewUserService,
			services.NewRoleService,
			services.NewUserRoleService,
			services.NewAuditService,
			services.NewTeamService,
			services.NewTeamUserService,
			services.NewProjectService,
			services.NewProjectUserService,
		),
		//Handler注入
		fx.Provide(
			handler.NewUserHandler,
		),
		//运行执行
		fx.Invoke(
			global.Init,
			//实例注册
			func(
				servicesInstance services.ServiceContext,
				reposInstance repos.RepoContext,
				handlerInstance handler.HandlersContext,
			) {
				services.Instance = servicesInstance
				repos.Instance = reposInstance
				handler.Instance = handlerInstance
				fmt.Println("所有组件初始化成功")
			},
			//服务注册
			func(
				engine *engine.Engine,
				sqlDB *sql.DB,
				server *http.Server,
			) {
				//数据库迁移
				if err := database.Migrate(sqlDB); err != nil {
					panic(fmt.Sprintf("迁移数据库失败: %v", err))
				}
			},
		),
	)
	container.Run()
}
