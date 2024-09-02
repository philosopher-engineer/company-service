package v1

import (
	"company-service/internal/company"
	"company-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, companyHandler *company.Handler) {
	api := router.Group("/api/v1")
	{
		api.POST("/companies", middleware.JWTAuthMiddleware(), companyHandler.CreateCompany)
		api.GET("/companies/:id", companyHandler.GetCompany)
		api.PATCH("/companies/:id", middleware.JWTAuthMiddleware(), companyHandler.UpdateCompany)
		api.DELETE("/companies/:id", middleware.JWTAuthMiddleware(), companyHandler.DeleteCompany)
	}
}
