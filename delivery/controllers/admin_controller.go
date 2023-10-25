package controllers

import (
	"gorm-practice/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService services.AdminService) *adminController {
	return &adminController{adminService: adminService}
}

func (c *adminController) FindAllUser(ctx *gin.Context) {

	data, err := c.adminService.GetAllUsers()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}
