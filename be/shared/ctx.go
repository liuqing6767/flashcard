package shared

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Context interface {
	echo.Logger
	context.Context
}

func EchoCtx2LogCtx(ectx echo.Context) Context {
	return &struct {
		echo.Logger
		context.Context
	}{
		Logger:  ectx.Logger(),
		Context: context.TODO(),
	}
}