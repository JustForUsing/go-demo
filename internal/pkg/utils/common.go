package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// FindRootPath 向上遍历查找 go.mod 文件，确定项目根目录
func FindRootPath() string {
	// 获取当前代码文件所在的绝对路径
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)

	// 向上遍历直到找到 go.mod
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir // 找到 go.mod，这就是项目根目录
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// 已经到达文件系统根目录，还没找到 go.mod
			log.Fatal("无法找到项目根目录 (go.mod 文件)")
		}
		dir = parentDir
	}
}

// FindExecPath 获取可执行文件的路径
func FindExecPath() string {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("获取可执行文件路径失败: %v", err)
	}
	return filepath.Dir(execPath)
}

// HandleError 错误输出，如果错误还错了，就直接退出
func HandleError(err error) {
	if err != nil {
		fmt.Printf("程序运行发生错误: %v", err)
	}
}
