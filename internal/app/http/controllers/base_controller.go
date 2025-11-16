package controllers

import (
	"net/http"

	appErr "goframe/internal/core/errors"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// Structs for structured JSON responses
type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type ValidationErrorResponse struct {
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}

type CreatedResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NoContentResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ----------------- Methods -----------------

func (bc *BaseController) Success(c *gin.Context, message string, data interface{}) {
	resp := SuccessResponse{
		Code:    http.StatusOK,
		Data:    data,
		Message: message,
		Status:  true,
	}
	c.JSON(http.StatusOK, resp)
}

func (bc *BaseController) Error(c *gin.Context, message string, err interface{}) {
	resp := ErrorResponse{
		Code:    http.StatusBadRequest,
		Error:   err,
		Message: message,
		Status:  false,
	}
	c.JSON(http.StatusBadRequest, resp)
}

func (bc *BaseController) ValidationError(c *gin.Context, errors interface{}) {
	resp := ValidationErrorResponse{
		Code:    http.StatusUnprocessableEntity,
		Errors:  errors,
		Message: "Validation failed",
		Status:  false,
	}
	c.JSON(http.StatusUnprocessableEntity, resp)
}

func (bc *BaseController) Created(c *gin.Context, message string, data interface{}) {
	resp := CreatedResponse{
		Code:    http.StatusCreated,
		Data:    data,
		Message: message,
		Status:  true,
	}
	c.JSON(http.StatusCreated, resp)
}

func (bc *BaseController) NoContent(c *gin.Context) {
	resp := NoContentResponse{
		Code:    http.StatusNoContent,
		Message: "No content",
		Status:  true,
	}
	c.JSON(http.StatusNoContent, resp)
}

func (bc *BaseController) BadRequest(c *gin.Context, message string, err interface{}) {
	resp := ErrorResponse{
		Code:    http.StatusBadRequest,
		Status:  false,
		Message: message,
		Error:   err,
	}
	c.JSON(http.StatusBadRequest, resp)
}

func (bc *BaseController) Unauthorized(c *gin.Context, message string, err interface{}) {
	resp := ErrorResponse{
		Code:    http.StatusUnauthorized,
		Status:  false,
		Message: message,
		Error:   err,
	}
	c.JSON(http.StatusUnauthorized, resp)
}

func (bc *BaseController) Forbidden(c *gin.Context, message string, err interface{}) {
	resp := ErrorResponse{
		Code:    http.StatusForbidden,
		Status:  false,
		Message: message,
		Error:   err,
	}
	c.JSON(http.StatusForbidden, resp)
}

func (bc *BaseController) NotFound(c *gin.Context, message string, err interface{}) {
	resp := ErrorResponse{
		Code:    http.StatusNotFound,
		Status:  false,
		Message: message,
		Error:   err,
	}
	c.JSON(http.StatusNotFound, resp)
}

func (bc *BaseController) Conflict(c *gin.Context, message string, err interface{}) {
	resp := ErrorResponse{
		Code:    http.StatusConflict,
		Status:  false,
		Message: message,
		Error:   err,
	}
	c.JSON(http.StatusConflict, resp)
}

func (bc *BaseController) InternalError(c *gin.Context, message string, err interface{}) {
	resp := ErrorResponse{
		Code:    http.StatusInternalServerError,
		Status:  false,
		Message: message,
		Error:   err,
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func (bc *BaseController) HandleServiceError(c *gin.Context, err error) {
	if e, ok := err.(*appErr.AppError); ok {

		switch e.Code {
		case "NOT_FOUND":
			c.JSON(http.StatusNotFound, ErrorResponse{
				Code:    http.StatusNotFound,
				Status:  false,
				Message: e.Message,
				Error:   e.Code,
			})
			return

		case "DUPLICATE":
			c.JSON(http.StatusConflict, ErrorResponse{
				Code:    http.StatusConflict,
				Status:  false,
				Message: e.Message,
				Error:   e.Code,
			})
			return

		case "INVALID":
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  false,
				Message: e.Message,
				Error:   e.Code,
			})
			return

		case "UNAUTHORIZED":
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  false,
				Message: e.Message,
				Error:   e.Code,
			})
			return
		}

	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:    http.StatusInternalServerError,
		Status:  false,
		Message: "Internal Server Error",
		Error:   err.Error(),
	})
}
