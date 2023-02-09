package v1

import (
	"github.com/LucasMateus-eng/simple-bank/adapters/primary/api/rest/router"
	"github.com/LucasMateus-eng/simple-bank/adapters/primary/api/rest/wallet"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/LucasMateus-eng/simple-bank/docs"
	echoswagger "github.com/swaggo/echo-swagger"
)

// @title Simple-bank API
// @version 1.0
// @description This is a simple-bank management application.

// @contact.name API Supports
// @contact.email lucas.falecomigo@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func SetupEchoRouter(handler *wallet.WalletHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")
	{
		utilsGroup := v1.Group("/health")
		{
			utilsGroup.GET("", router.Health)
		}
	}
	{
		docs := v1.Group("/docs")
		{
			docs.GET("/swagger/*", echoswagger.WrapHandler)
		}
	}
	{
		wallet := v1.Group("/wallet")
		{
			wallet.GET(":wallet_id", handler.Get)
			wallet.POST("", handler.Add)
			wallet.PUT("", handler.Update)
			wallet.PUT("/transfer", handler.Transfer)
			wallet.PUT("/deposit", handler.Deposit)
			wallet.DELETE(":wallet_id", handler.Delete)
		}
	}

	return e
}
