package v1

import (
	"github.com/LucasMateus-eng/simple-bank/adapters/primary/api/rest/wallet"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupEchoRouter(handler *wallet.WalletHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")
	{
		wallet := v1.Group("/wallet")
		{
			wallet.GET(":wallet_id", handler.Get)
			wallet.POST("", handler.Add)
			wallet.PUT("", handler.Update)
			wallet.DELETE(":wallet_id", handler.Delete)
		}
	}

	return e
}
