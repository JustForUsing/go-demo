package engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"item-manager-new/internal/pkg/global"
)

type Engine struct {
	*gin.Engine
}

func New() *Engine {
	fmt.Printf("engin_server_mod: %v", global.GetViperConfigString("server.mode"))

	gin.SetMode(global.GetViperConfigString("server.mode"))
	engine := gin.New()
	engine.Use(gin.Recovery())

	bindRouter(engine)

	return &Engine{
		Engine: engine,
	}
}
