package controllers

import (
	"net/http"

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
