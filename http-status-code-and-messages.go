package mop_common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal_server_error")
	ErrRecordNotFound = errors.New("record_not_found")
)

type httpErrorWithErrorSlug struct {
	HumanizedError string `json:"humanized_error"`
	ErrorSlug string `json:"error_slug"`
	Status    int    `json:"status"`
}

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

func ThrowStatusBadRequest(err error, c *gin.Context, optionalMessage ...string) {
	message := "Bad request"
	if len(optionalMessage) > 0 {
		message = optionalMessage[0]
	}

	e := httpErrorWithErrorSlug{
		HumanizedError: message,
		ErrorSlug:      err.Error(),
		Status:         http.StatusBadRequest,
	}

	c.JSON(http.StatusBadRequest, e)
}

func ThrowStatusUnauthorized(c *gin.Context, optionalMessage ...string) {
	message := "Unauthorized"
	if len(optionalMessage) > 0 {
		message = optionalMessage[0]
	}

	e := httpErrorWithErrorSlug{
		HumanizedError: message,
		ErrorSlug:      ErrUnauthorized.Error(),
		Status:         http.StatusUnauthorized,
	}

	c.JSON(http.StatusUnauthorized, e)
}

func ThrowStatusInternalServerError(c *gin.Context, optionalMessage ...string) {
	message := "Internal server error"
	if len(optionalMessage) > 0 {
		message = optionalMessage[0]
	}

	e := httpErrorWithErrorSlug{
		HumanizedError: message,
		ErrorSlug:      ErrInternalServerError.Error(),
		Status:         http.StatusInternalServerError,
	}

	c.JSON(http.StatusInternalServerError, e)
}

func ThrowStatusNotFound(c *gin.Context, optionalMessage ...string) {
	message := "Record not found"
	if len(optionalMessage) > 0 {
		message = optionalMessage[0]
	}

	e := httpErrorWithErrorSlug{
		HumanizedError: message,
		ErrorSlug:      ErrRecordNotFound.Error(),
		Status:         http.StatusNotFound,
	}

	c.JSON(http.StatusNotFound, e)
}

