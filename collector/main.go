package main

import (
	"collector/handlers"
	"collector/middleware"
	"collector/repository"
	"collector/store"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	apiKey := os.Getenv("X_API_KEY")
	if apiKey == "" {
		log.Fatal("API Key is not set")
	}

	pool, err := pgxpool.New(
		context.Background(),
		databaseURL,
	)
	if err != nil {
		log.Fatal("failed to create postgres pool:", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("database is not reachable:", err)
	}

	//store
	metricsStore := store.NewAgentStore()
	metricsRepository := repository.NewMetricsRepository(pool)

	router := gin.Default()

	authenticatedPaths := router.Group("/")
	authenticatedPaths.Use(middleware.APIKeyMiddleware(apiKey))

	authenticatedPaths.POST(
		"/agents/:id/metrics",
		handlers.SubmitMetricsHandler(metricsStore, metricsRepository),
	)
	authenticatedPaths.GET(
		"/agents/:id",
		handlers.GetAgentHandler(metricsStore),
	)

	authenticatedPaths.GET(
		"/agents",
		handlers.GetAgentsHandler(metricsStore),
	)

	authenticatedPaths.GET(
		"/agents/:id/history",
		handlers.GetAgentHistoryHandler(metricsRepository),
	)

	if err := router.Run(":8080"); err != nil {
		return
	}
}
