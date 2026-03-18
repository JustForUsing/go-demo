package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"item-manager-new/internal/services"
	"path/filepath"

	"item-manager-new/internal/pkg/global"
	"item-manager-new/internal/pkg/utils"
)

func Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("创建数据库驱动失败: %w", err)
	}

	// 初始化数据库表
	err = initDb(driver)
	utils.HandleError(err)

	//预设数据
	presetData()
	return nil
}

// initDb 初始化数据库迁移
func initDb(driver database.Driver) error {
	migrationPath := "file://" + filepath.Join(global.ItemRootPath(), "migrations")
	fmt.Printf("迁移文件路径: %s\n", migrationPath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		global.GetViperConfigString("database.dbname"),
		driver,
	)
	if err != nil {
		panic(fmt.Sprintf("创建迁移实例失败: %v", err))
	}

	// 添加详细的版本信息输出
	version, dirty, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		panic(fmt.Sprintf("获取版本信息时出错: %v\n", err))
	}
	fmt.Printf("当前数据库版本: %d, 脏状态: %t\n", version, dirty)

	// 处理脏版本
	if dirty {
		fmt.Printf("检测到脏版本 %d，尝试强制重置...\n", version)
		if err := m.Force(int(version)); err != nil {
			panic(fmt.Sprintf("强制重置脏版本失败: %v", err))
		}
		fmt.Println("脏版本已重置成功")
		// 重置后重新获取版本信息
		version, _, _ = m.Version()
		fmt.Printf("重置后的版本: %d\n", version)
	}

	fmt.Println("开始执行数据库迁移...")
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("数据库已是最新版本，无需迁移")
			return nil
		}
		panic(fmt.Sprintf("执行迁移失败: %v", err))
	}

	fmt.Println("数据库迁移完成")
	return nil
}

// presetData 预设数据
func presetData() {
	fmt.Println("预设数据开始")
	//预设管理员
	adminId, _ := services.Instance.User.PresetAdmin()
	fmt.Printf("预设管理员用户ID: %d\n", adminId)
	//预设角色
	services.Instance.Role.PresetAdminRole(adminId)
}
