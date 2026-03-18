package handler

import (
	"go.uber.org/fx"
)

var Instance HandlersContext

type HandlersContext struct {
	fx.In
	User *UserHandler
}
