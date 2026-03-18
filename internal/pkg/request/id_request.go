package request

import "github.com/gin-gonic/gin"

const (
	headerName = "X-Request-ID"
	contextKey = "requestId"
)

// HeaderName 返回请求 ID 使用的 HTTP Header。
func HeaderName() string {
	return headerName
}

// FromContext 获取 Gin Context 中的请求 ID。
func FromContext(c *gin.Context) string {
	if value, exists := c.Get(contextKey); exists {
		if reqID, ok := value.(string); ok {
			return reqID
		}
	}
	return ""
}

// Set 将请求 ID 写入上下文并暴露到 Response Header。
func Set(c *gin.Context, id string) {
	c.Set(contextKey, id)
	c.Writer.Header().Set(headerName, id)
}
