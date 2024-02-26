package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-difranco/go-test/usecase"
)

type StatusController struct {
	useCase usecase.StatusUseCase
}

func NewStatusController(useCase usecase.StatusUseCase) *StatusController {
	return &StatusController{useCase}
}

func (ctrl *StatusController) FetchAll(c *gin.Context) {
	statuses, err := ctrl.useCase.FetchAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, statuses)
}
