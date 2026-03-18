package global

import (
	"fmt"
	"item-manager-new/internal/pkg/utils"
	"sync"
)

var (
	once sync.Once
)

// Init 初始化全局变量
func Init() {
	once.Do(func() {
		itemRootPath = utils.FindRootPath()
		execPath = utils.FindExecPath()
		ConfigLoad()
	})
	fmt.Println("全局初始化完成")
}
