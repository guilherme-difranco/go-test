package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-difranco/go-test/usecase"
)

type PriorityController struct {
	useCase usecase.PriorityUseCase
}

func NewPriorityController(useCase usecase.PriorityUseCase) *PriorityController {
	return &PriorityController{useCase}
}

func (ctrl *PriorityController) FetchAll(c *gin.Context) {
	priorities, err := ctrl.useCase.FetchAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, priorities)
}
