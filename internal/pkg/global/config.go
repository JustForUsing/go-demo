package global

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

type ServerConfig struct {
	Host            string
	Port            int
	Mode            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}
type DatabaseConfig struct {
	Driver          string
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	DSN             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}
type AdminConfig struct {
	Username string
	Password string
	Email    string
	Nickname string
}
type LoggerConfig struct {
	Level       string
	Encoding    string
	Development bool
	File        string
	MaxSize     int  // 每个日志文件的最大大小(MB)
	MaxAge      int  // 保留旧日志文件的最大天数
	MaxBackups  int  // 保留的最大旧日志文件数量
	Compress    bool // 是否压缩旧日志文件
}

type Config struct {
	viper    *viper.Viper
	Server   ServerConfig
	Database DatabaseConfig
	Admin    AdminConfig
	Logger   LoggerConfig
	Auth     AuthConfig
}

type AuthConfig struct {
	CookieName string
	SessionTTL time.Duration
}

var (
	LoadConfig Config
)

// ConfigLoad 读取配置文件与环境变量。
func ConfigLoad() {
	LoadConfig.viper = buildViper()
	configFile := os.Getenv("ITEM_CONFIG_FILE")
	if configFile != "" {
		LoadConfig.viper.SetConfigFile(configFile)
	} else {
		LoadConfig.viper.SetConfigFile(filepath.Join(ItemRootPath(), "configs", "default.yaml"))
	}
	if err := LoadConfig.viper.ReadInConfig(); err != nil {
		fmt.Printf("读取配置失败: %v", err)
	}
	if err := LoadConfig.viper.Unmarshal(&LoadConfig); err != nil {
		fmt.Printf("解析配置失败: %v", err)
	}
	fmt.Printf("config: %v\n", LoadConfig)
}

func buildViper() *viper.Viper {
	v := viper.New()

	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.mode", "release")
	v.SetDefault("server.readTimeout", "10s")
	v.SetDefault("server.writeTimeout", "15s")
	v.SetDefault("server.idleTimeout", "60s")
	v.SetDefault("server.shutdownTimeout", "10s")

	v.SetDefault("logger.level", "info")
	v.SetDefault("logger.encoding", "json")
	v.SetDefault("logger.development", false)
	v.SetDefault("logger.file", "./logs/global.log")
	v.SetDefault("logger.maxSize", 100)
	v.SetDefault("logger.maxAge", 7)
	v.SetDefault("logger.maxBackups", 3)
	v.SetDefault("logger.compress", true)

	v.SetDefault("database.driver", "postgres")
	v.SetDefault("database.host", "127.0.0.11")
	v.SetDefault("database.port", "5432")
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.password", "123456")
	v.SetDefault("database.dbname", "item_manager")
	v.SetDefault("database.maxIdleConns", 5)
	v.SetDefault("database.maxOpenConns", 10)
	v.SetDefault("database.connMaxLifetime", "60m")

	v.SetDefault("auth.cookieName", "ITEM_SESSION")
	v.SetDefault("auth.sessionTTL", "24h")

	v.SetDefault("admin.username", "admin")
	v.SetDefault("admin.password", "adminadmin")
	v.SetDefault("admin.email", "")
	v.SetDefault("admin.nickname", "Administrator")

	return v
}

// GetServerConfig 获取服务器配置
func GetServerConfig() ServerConfig {
	return LoadConfig.Server
}

// GetDatabaseConfig 获取数据库配置
func GetDatabaseConfig() DatabaseConfig {
	return LoadConfig.Database
}

// GetAdminConfig 获取管理员配置
func GetAdminConfig() AdminConfig {
	return LoadConfig.Admin
}

// GetLoggerConfig 获取日志配置
func GetLoggerConfig() LoggerConfig {
	return LoadConfig.Logger
}

// GetAuthConfig 获取认证配置
func GetAuthConfig() AuthConfig { return LoadConfig.Auth }

// GetViperConfig 直接从viper获取配置值，返回map类型的数据
func GetViperConfig(key ...string) interface{} {
	// 确保viper实例已初始化
	if LoadConfig.viper == nil {
		return nil
	}

	if len(key) == 0 {
		return LoadConfig.viper.AllSettings()
	}

	// 使用viper获取配置值
	return LoadConfig.viper.Get(key[0])
}

// GetViperConfigString 获取字符串类型的配置值
func GetViperConfigString(key string) string {
	// 确保viper实例已初始化
	if LoadConfig.viper == nil {
		return ""
	}
	return LoadConfig.viper.GetString(key)
}

// GetViperConfigInt 获取整数类型的配置值
func GetViperConfigInt(key string) int {
	// 确保viper实例已初始化
	if LoadConfig.viper == nil {
		return 0
	}
	return LoadConfig.viper.GetInt(key)
}

// GetViperConfigBool 获取布尔类型的配置值
func GetViperConfigBool(key string) bool {
	// 确保viper实例已初始化
	if LoadConfig.viper == nil {
		return false
	}
	return LoadConfig.viper.GetBool(key)
}

// GetViperConfigMap 获取map[string]interface{}类型的配置值
func GetViperConfigMap(key string) map[string]interface{} {
	// 确保viper实例已初始化
	if LoadConfig.viper == nil {
		return make(map[string]interface{})
	}
	return LoadConfig.viper.GetStringMap(key)
}
