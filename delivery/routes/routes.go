package routes

import (
	"gorm-practice/config"
	"gorm-practice/delivery/controllers"
	"gorm-practice/delivery/middlewares"
	"gorm-practice/managers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRouter(router *gin.Engine) error {

	router.Use(middlewares.LogRequestMiddleware(logrus.New()))

	manager := managers.NewInfraManager(config.Cfg)
	repositories := managers.NewRepoManager(manager)
	services := managers.NewServiceManager(repositories)

	// Auth
	authController := controllers.NewAuthController(services.AuthService())
	// Admin
	adminController := controllers.NewAdminController(services.AdminService())
	// Customer
	customerController := controllers.NewCustomerController(services.CustomerService(), services.TransactionService())

	v1 := router.Group("/api/v1")
	{
		// Auth Route
		auth := v1.Group("/auth")
		{
			auth.POST("/registration", authController.Registration)
			auth.POST("/login", authController.Login)
			auth.POST("/logout", authController.Logout)
		}

		// Admin Route
		admin := v1.Group("/admin", middlewares.AuthMiddleware())
		{
			management := admin.Group("/management")
			{
				management.GET("/users", adminController.FindAllUser)
			}
		}

		// Customer Route
		customer := v1.Group("/customers", middlewares.AuthMiddleware())
		{
			customer.GET("/profile", customerController.Profile)
			customer.GET("/transactions/history", customerController.HistoryTransaction)
			customer.POST("transactions/:id/send", customerController.SendMoney)
			// customer.POST("/:id/card", customerController.AddCard)
		}
	}

	return router.Run(":8080")
}
