package service

import "github.com/labstack/echo/v4"

type Handler interface {
	FilterHasAccount() echo.HandlerFunc
}
