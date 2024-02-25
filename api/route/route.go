package route

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/guilherme-difranco/go-test/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Env, timeout time.Duration, db *mongo.Database, router *gin.Engine) {
	publicRouter := router.Group("")

	publicRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Bem-vindo à API!"})
	})

	NewTaskRouter(env, timeout, db, publicRouter)
}
