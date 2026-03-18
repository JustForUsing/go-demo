package engine

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/fx"
	"item-manager-new/internal/pkg/global"
	"log"
	"net"
	"net/http"
	"os"
)

func NewHttp(lc fx.Lifecycle, engine *Engine) *http.Server {
	cfg := global.GetServerConfig()
	fmt.Printf("server_cfg: %v", cfg)
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      engine.Engine,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Printf("服务启动开始，监听地址: %s", server.Addr)
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go func() {
				err := server.Serve(ln)
				if err == nil {
					return
				}
				if errors.Is(err, http.ErrServerClosed) {
					fmt.Printf("服务正常关闭: %v", err)
					return
				}
				_, _ = fmt.Fprintf(os.Stderr, "服务启动失败: %v", err)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("服务关闭开始")
			if err := server.Shutdown(ctx); err != nil {
				return err
			}
			return nil
		},
	})
	return server
}
