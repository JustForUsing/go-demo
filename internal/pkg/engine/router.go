package engine

import (
	"github.com/gin-gonic/gin"
	"item-manager-new/internal/api/handler"
)

func bindRouter(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	{
		apiGroup.POST("/login", handler.Instance.User.Login)
	}
}
