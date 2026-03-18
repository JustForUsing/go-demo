package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"item-manager-new/internal/pkg/global"
	"log"
	"os"
	"time"
)

// New 初始化数据库连接
func New() (*gorm.DB, *sql.DB, error) {
	cfg := global.GetDatabaseConfig()
	fmt.Printf("database_new_cfg: %v", cfg)
	//获取gorm对象
	db, sqlDB, err := Open(&cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("初始化数据库失败: %w", err)
	}
	//////初始化数据库表
	//if err := Migrate(sqlDB); err != nil {
	//	return nil, fmt.Errorf("迁移数据库失败: %w", err)
	//}
	return db, sqlDB, nil
}

// Open 打开数据库连接
func Open(cfg *global.DatabaseConfig) (*gorm.DB, *sql.DB, error) {
	var dialector gorm.Dialector
	switch cfg.Driver {
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.DBName,
		)
		fmt.Printf("postgres dsn: %v", dsn)
		dialector = postgres.Open(dsn)
	case "sqlite":
		if cfg.DSN == "" {
			cfg.DSN = "file:item-manager.business?_foreign_keys=on"
		}
		dialector = sqlite.Open(cfg.DSN)
	default:
		return nil, nil, fmt.Errorf("不支持的数据库驱动: %s", cfg.Driver)
	}

	//根据配置环境设置日志级别
	var logLevel logger.LogLevel
	serverMode := global.GetViperConfigString("server.mode")
	if serverMode == "debug" || serverMode == "test" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}
	//配置控制台输出所有Sql查询日志
	newLogger := logger.New(
		// 日志输出目标：控制台（也可改为文件）
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值（超过 1 秒打印警告）
			LogLevel:                  logLevel,    // 日志级别：Info（打印所有 SQL）
			IgnoreRecordNotFoundError: true,        // 忽略 "记录未找到" 错误
			Colorful:                  true,        // 彩色打印（开发环境友好）
		},
	)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	//设置数据库连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("访问数据库连接池失败: %w", err)
	}

	if cfg.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	return db, sqlDB, nil
}
