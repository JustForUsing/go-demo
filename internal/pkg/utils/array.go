package utils

// InStrMap 检查元素是否存在于数组中
func InStrMap(checkStr string, strArr []string) bool {
	// 为了性能优化，预先分配适当的容量
	items := make(map[string]bool, len(strArr))
	for _, item := range strArr {
		items[item] = true
	}
	return items[checkStr]
}
