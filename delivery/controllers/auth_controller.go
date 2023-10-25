package controllers

import (
	"errors"
	"gorm-practice/models"
	"gorm-practice/models/dto"
	"gorm-practice/services"
	"gorm-practice/utils/common"
	"gorm-practice/utils/exception"
	"gorm-practice/utils/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *authController {
	return &authController{authService: authService}
}

func (c *authController) Registration(ctx *gin.Context) {
	var request models.Customer

	request.ID = common.GenerateUUID()

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Status:  exception.StatusBadRequest,
			Message: exception.FieldErrors(err),
		})
		return
	}

	_, err := c.authService.CreateNewCustomer(&request)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrFailedGeneratePassword) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrUserNameExist) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrEmailExist) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  exception.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Code:    http.StatusCreated,
		Status:  exception.StatusSuccess,
		Message: exception.StatusRegistrationSuccess,
	})
}

func (c *authController) Login(ctx *gin.Context) {
	var request models.Auth

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Status:  exception.StatusBadRequest,
			Message: exception.FieldErrors(err),
		})
		return
	}

	data, err := c.authService.Login(request.UserName, request.Password)

	if err != nil {
		if errors.Is(err, exception.ErrCustomerDoesntExist) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  exception.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseToken{
		Code:    http.StatusOK,
		Status:  exception.StatusSuccess,
		Message: exception.StatusLoginSuccess,
		Token:   data,
	})
}

func (c *authController) Logout(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Status:  exception.StatusUnauthorized,
			Message: exception.ErrTokenNotProvided.Error(),
		})
		return
	}

	tokenString := strings.Split(authorization, " ")[1]

	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Status:  exception.StatusUnauthorized,
			Message: exception.ErrTokenRequired.Error(),
		})
		return
	}

	_, err := security.VerifyAccessToken(tokenString)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Status:  exception.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	security.BlackListedTokens[tokenString] = struct{}{}

	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Status:  exception.StatusSuccess,
		Message: exception.StatusLogoutSuccess,
	})
}
