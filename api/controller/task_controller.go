package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-difranco/go-test/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func NewTaskController(taskUsecase domain.TaskUsecase) *TaskController {
	return &TaskController{
		TaskUsecase: taskUsecase,
	}
}

func (tc *TaskController) CreateBatch(c *gin.Context) {
	var tasks []domain.Task

	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range tasks {
		tasks[i].ID = primitive.NewObjectID()
	}

	if err := tc.TaskUsecase.CreateBatch(c.Request.Context(), tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tasks created successfully"})
}

func (tc *TaskController) FetchByUserID(c *gin.Context) {
	userID := c.Param("userID")

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks, err := tc.TaskUsecase.FetchByUserID(c.Request.Context(), userObjectID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// FetchTasks busca tarefas com suporte para filtragem, projeção, paginação e ordenação.
func (tc *TaskController) FetchTasks(c *gin.Context) {

	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "20"), 10, 64)
	if err != nil || limit <= 0 || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'limit' is invalid. Must be between 1 and 100."})
		return
	}

	skip, err := strconv.ParseInt(c.DefaultQuery("skip", "0"), 10, 64)
	if err != nil || skip < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'skip' is invalid. Must be a non-negative integer."})
		return
	}

	userID := c.Query("userID")
	var filter bson.M
	if userID != "" {
		userObjectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'userID' is invalid."})
			return
		}
		filter = bson.M{"userID": userObjectID}
	} else {
		filter = bson.M{}
	}

	projectionFields := c.QueryArray("fields")
	var projection bson.M
	if len(projectionFields) > 0 {
		projection = bson.M{}
		for _, field := range projectionFields {
			projection[field] = 1
		}
	}

	tasks, err := tc.TaskUsecase.FetchUserTasks(c.Request.Context(), filter, projection, limit, skip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks. Please try again later."})
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No tasks found."})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
