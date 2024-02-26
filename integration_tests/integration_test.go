package integration_tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/guilherme-difranco/go-test/cache"
	"github.com/guilherme-difranco/go-test/config"
	"github.com/guilherme-difranco/go-test/domain"
	"github.com/guilherme-difranco/go-test/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func TestTaskCreation(t *testing.T) {
	app := config.App()
	env := app.Env
	defer app.CloseDBConnection()

	// Inicialização do serviço de cache e carregamento dos dados no cache
	cacheService := cache.NewCacheService(app.Mongo.Database(app.Env.DBName))
	cacheService.InitializeCache()

	// Busca por Priority e Status no cache
	priorityName := "Alta"
	priority, found := cacheService.GetPriority(priorityName)
	if !found {
		log.Printf("Priority %s não encontrado no cache.\n", priorityName)
	} else {
		log.Printf("Priority %s encontrado no cache.\n", priorityName)
	}

	statusName := "Pendente"
	status, found := cacheService.GetStatus(statusName)
	if !found {
		log.Printf("Status %s não encontrado no cache.\n", statusName)
	} else {
		log.Printf("Status %s encontrado no cache.\n", statusName)
	}

	// Criação de uma tarefa de teste
	taskRepo := repository.NewTaskRepository(app.Mongo.Database(env.DBName), domain.CollectionTask)
	task := domain.Task{
		Title:        "Integration Test Task",
		Description:  "This is a task created for an integration test.",
		Priority:     priority,
		Status:       status,
		CreationDate: time.Now(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Inserção da tarefa de teste
	var err = taskRepo.Create(ctx, &task)
	if err != nil {
		log.Fatalf("Failed to insert test task: %v", err)
	}

	log.Printf("Task with ID %s created successfully.", task.ID.Hex())
}

func TestFetchTasksIntegration(t *testing.T) {
	// Inicialização e configuração
	app := config.App()
	env := app.Env
	defer app.CloseDBConnection()

	// Inicialização do serviço de cache e carregamento dos dados no cache
	cacheService := cache.NewCacheService(app.Mongo.Database(app.Env.DBName))
	cacheService.InitializeCache()

	// Busca por Priority e Status no cache
	priorityName := "Alta"
	priority, found := cacheService.GetPriority(priorityName)
	if !found {
		log.Printf("Priority %s não encontrado no cache.\n", priorityName)
	} else {
		log.Printf("Priority %s encontrado no cache.\n", priorityName)
	}

	statusName := "Pendente"
	status, found := cacheService.GetStatus(statusName)
	if !found {
		log.Printf("Status %s não encontrado no cache.\n", statusName)
	} else {
		log.Printf("Status %s encontrado no cache.\n", statusName)
	}

	// Criação de uma tarefa de teste
	taskRepo := repository.NewTaskRepository(app.Mongo.Database(env.DBName), domain.CollectionTask)
	task := domain.Task{
		Title:        "Integration Test Task",
		Description:  "This is a task created for an integration test.",
		Priority:     priority,
		Status:       status,
		CreationDate: time.Now(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Inserção da tarefa de teste
	var err error
	err = taskRepo.Create(ctx, &task)
	if err != nil {
		log.Fatalf("Failed to insert test task: %v", err)
	}

	// Busca por tarefas
	projection := bson.M{"title": 1, "description": 1}
	limit := int64(10)
	skip := int64(0)

	tasks, err := taskRepo.FetchTasks(context.TODO(), bson.M{}, projection, limit, skip)
	if err != nil {
		t.Fatalf("Failed to fetch tasks: %v", err)
	}

	if len(tasks) > 0 {
		log.Println("Fetched tasks:")
		for _, t := range tasks {
			log.Printf("Task ID: %s, Title: %s, Description: %s\n", t.ID.Hex(), t.Title, t.Description)
		}
	} else {
		t.Errorf("Expected to fetch at least one task, got 0")
	}
}
