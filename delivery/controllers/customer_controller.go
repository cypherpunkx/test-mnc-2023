package controllers

import (
	"errors"
	"gorm-practice/models"
	"gorm-practice/models/dto"
	"gorm-practice/services"
	"gorm-practice/utils/common"
	"gorm-practice/utils/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService    services.CustomerService
	transactionService services.TransactionService
}

func NewCustomerController(customerService services.CustomerService, transactionService services.TransactionService) *CustomerController {
	return &CustomerController{customerService: customerService, transactionService: transactionService}
}

func (c *CustomerController) Profile(ctx *gin.Context) {
	username, exists := ctx.Get("username")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: exception.StatusUnauthorized,
		})
		return
	}

	if username, ok := username.(string); ok {
		data, err := c.customerService.GetProfile(username)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:   http.StatusInternalServerError,
				Status: exception.StatusInternalServerError,
			})
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Code:    http.StatusOK,
			Status:  exception.StatusSuccess,
			Message: "Get Profile",
			Data:    data,
		})
	}

}

func (c *CustomerController) SendMoney(ctx *gin.Context) {
	var request models.Transaction

	var id = ctx.Param("id")

	username, exists := ctx.Get("username")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: exception.StatusUnauthorized,
		})
		return
	}

	request.ID = common.GenerateUUID()

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Status:  exception.StatusBadRequest,
			Message: exception.FieldErrors(err),
		})
		return
	}

	if username, ok := username.(string); ok {
		data, err := c.transactionService.SendMoneyToFriend(&request, id, username)

		if err != nil {
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
			Message: exception.StatusTransactionSuccess,
			Data:    data,
		})
	}
}

func (c *CustomerController) HistoryTransaction(ctx *gin.Context) {
	username, exists := ctx.Get("username")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: exception.StatusUnauthorized,
		})
		return
	}

	if username, ok := username.(string); ok {
		data, err := c.transactionService.GetHistory(username)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:   http.StatusInternalServerError,
				Status: exception.StatusInternalServerError,
			})
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Code:    http.StatusOK,
			Status:  exception.StatusSuccess,
			Message: "Get All Transactions",
			Data:    data,
		})

	}
}

// func (c *CustomerController) Pay(ctx *gin.Context) {
// 	var request models.Transaction

// 	var id = ctx.Param("id")

// 	username, exists := ctx.Get("username")

// 	if !exists {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: exception.StatusUnauthorized,
// 		})
// 		return
// 	}

// 	request.ID = common.GenerateUUID()

// 	fmt.Println(username)

// 	if err := ctx.ShouldBindJSON(&request); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
// 			Code:    http.StatusBadRequest,
// 			Status:  exception.StatusBadRequest,
// 			Message: exception.FieldErrors(err),
// 		})
// 		return
// 	}

// 	data, err := c.transactionService.SendMoneyToFriend(&request, id)

// 	if err != nil {
// 		if errors.Is(err, exception.ErrFailedCreate) {
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
// 				Code:    http.StatusInternalServerError,
// 				Status:  exception.StatusInternalServerError,
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		if errors.Is(err, exception.ErrFailedGeneratePassword) {
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
// 				Code:    http.StatusInternalServerError,
// 				Status:  exception.StatusInternalServerError,
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		if errors.Is(err, exception.ErrUserNameExist) {
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
// 				Code:    http.StatusInternalServerError,
// 				Status:  exception.StatusInternalServerError,
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		if errors.Is(err, exception.ErrEmailExist) {
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
// 				Code:    http.StatusInternalServerError,
// 				Status:  exception.StatusInternalServerError,
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
// 			Code:    http.StatusInternalServerError,
// 			Status:  exception.StatusInternalServerError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, dto.Response{
// 		Code:   http.StatusCreated,
// 		Status: exception.StatusRegistrationSuccess,
// 		Data:   data,
// 	})
// }
