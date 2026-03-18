package global

// 其他全局常量可以类似定义
const (
	AppName = "item-manager-new"
	Version = "1.0.0"
)

var (
	// 私有变量，防止外部直接修改
	itemRootPath string
	//可执行文件执行目录
	execPath string
)

// ItemRootPath 返回项目根路径
func ItemRootPath() string {
	return itemRootPath
}

// ExecPath 返回可执行文件执行目录
func ExecPath() string {
	return execPath
}
