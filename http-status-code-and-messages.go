package mop_common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal_server_error")
)


func ThrowStatusOK(i interface{}, c *gin.Context) {
	if i != nil {
		c.JSON(http.StatusOK, i)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "OK",
	})
}

func ThrowStatusCreated(i interface{}, c *gin.Context) {
	if i != nil {
		c.JSON(http.StatusCreated, i)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "OK",
	})
}

func ThrowStatusBadRequest(err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"error": err.Error(),
	})
}

func ThrowStatusUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  http.StatusUnauthorized,
		"error": ErrUnauthorized.Error(),
	})
}

func ThrowStatusInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"error": ErrInternalServerError.Error(),
	})
}