package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guilherme-difranco/go-test/api/controller"
	"github.com/guilherme-difranco/go-test/cache"
	"github.com/guilherme-difranco/go-test/config"
	"github.com/guilherme-difranco/go-test/domain"
	"github.com/guilherme-difranco/go-test/repository"
	"github.com/guilherme-difranco/go-test/usecase"
)

func main() {

	app := config.App()

	cacheService := cache.NewCacheService(app.Mongo.Database(app.Env.DBName))
	cacheService.InitializeCache()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		log.Printf("Recebida requisição para %s", c.Request.URL.Path)
		c.Next()
		log.Printf("Resposta enviada com status %d", c.Writer.Status())
	})

	router.Use(cors.Default())

	router.GET("/testes", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "CORS test successful",
		})
	})

	cacheService = cache.NewCacheService(db)

	// router.Use(cors.Default())

	// Setup task routes directly on the main router
	taskRepo := repository.NewTaskRepository(db, domain.CollectionTask)
	taskUsecase := usecase.NewTaskUsecase(taskRepo, timeout)
	taskController := controller.NewTaskController(taskUsecase)

	priorityRepo := repository.NewPriorityRepository(db, domain.CollectionPriorities)
	priorityUseCase := usecase.NewPriorityUseCase(priorityRepo, cacheService)
	priorityController := controller.NewPriorityController(priorityUseCase)

	statusRepo := repository.NewStatusRepository(db, domain.CollectionStatus)
	statusUseCase := usecase.NewStatusUseCase(statusRepo, cacheService)
	statusController := controller.NewStatusController(statusUseCase)

	router.GET("/tasks", taskController.FetchTasks)
	router.POST("/task", taskController.Create)
	router.POST("/tasks", taskController.CreateBatch)
	router.GET("/tasks/user/:userID", taskController.FetchByUserID)
	router.GET("/tasks/:id", taskController.FetchTaskByID)
	router.PUT("/tasks/:id", taskController.Update)
	router.DELETE("/tasks/:id", taskController.Delete)

	router.GET("/priorities", priorityController.FetchAll)
	router.GET("/statuses", statusController.FetchAll)

	//route.Setup(env, timeout, db, router)
	port := os.Getenv("PORT")
	if port != "" {
		port = ":" + port
	} else {
		port = env.ServerAddress // porta padrão se não houver variável de ambiente PORT
	}
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}

	router.Run(port)
}
