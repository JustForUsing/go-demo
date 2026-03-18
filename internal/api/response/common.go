package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"item-manager-new/internal/pkg/request"
)

type Builder struct {
	c *gin.Context
}

func NewBuilder(c *gin.Context) *Builder {
	return &Builder{
		c: c,
	}
}

// Response 标准响应格式。
type Response struct {
	RequestId string      `json:"request,omitempty"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

// JSON 输出统一响应。
func (b *Builder) JSON(code int, message string, data interface{}) {
	resp := Response{
		RequestId: request.FromContext(b.c),
		Code:      code,
		Message:   message,
		Data:      data,
	}
	b.c.JSON(code, resp)
}

// OK 输出成功响应。
func (b *Builder) OK(data interface{}) {
	b.JSON(http.StatusOK, "success", data)
}

// Error 输出错误响应并中断请求。
func (b *Builder) Error(code int, message string) {
	resp := Response{
		RequestId: request.FromContext(b.c),
		Code:      code,
		Message:   message,
	}
	logHTTPError(b.c, code, message, resp.RequestId)
	b.c.AbortWithStatusJSON(code, resp)
}

func (b *Builder) BadRequest(message string) {
	resp := Response{
		RequestId: request.FromContext(b.c),
		Code:      http.StatusBadRequest,
		Message:   message,
	}
	b.c.AbortWithStatusJSON(http.StatusBadRequest, resp)
}

func (b *Builder) Unauthorized(message string) {
	resp := Response{
		RequestId: request.FromContext(b.c),
		Code:      http.StatusUnauthorized,
		Message:   message,
	}
	b.c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
}

func (b *Builder) InternalServerError(message string) {
	resp := Response{
		RequestId: request.FromContext(b.c),
		Code:      http.StatusInternalServerError,
		Message:   message,
	}
	b.c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}

// NoContent 输出 204 响应。
func (b *Builder) NoContent() {
	b.c.Status(http.StatusNoContent)
}

func logHTTPError(c *gin.Context, code int, message, request string) {
	fields := []zap.Field{
		zap.Int("code", code),
		zap.String("message", message),
		zap.String("request", request),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
	}
	if len(c.Errors) > 0 {
		fields = append(fields, zap.String("errors", c.Errors.String()))
	}
	logger := zap.L()
	if code >= http.StatusInternalServerError {
		logger.Error("http error response", fields...)
	} else {
		logger.Warn("http error response", fields...)
	}
}
