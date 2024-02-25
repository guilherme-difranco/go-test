package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-difranco/go-test/api/controller"
	"github.com/guilherme-difranco/go-test/config"
	"github.com/guilherme-difranco/go-test/domain"
	"github.com/guilherme-difranco/go-test/repository"
	"github.com/guilherme-difranco/go-test/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(env *config.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	taskRepo := repository.NewTaskRepository(db, domain.CollectionTask)
	taskUsecase := usecase.NewTaskUsecase(taskRepo, timeout)
	taskController := controller.NewTaskController(taskUsecase)

	taskGroup := group.Group("/tasks")

	taskGroup.POST("/", taskController.CreateBatch)

	taskGroup.GET("/:userID", taskController.FetchByUserID)

}
