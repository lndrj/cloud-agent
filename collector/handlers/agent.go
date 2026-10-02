package handlers

import (
	"collector/models"
	"collector/repository"
	"collector/store"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SubmitMetricsHandler(
	store *store.AgentStore,
	repository *repository.MetricsRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.AgentMetrics
		id := c.Param("id")

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid metrics",
			})
			return
		}

		store.Set(id, request)
		ctx := c.Request.Context()

		if err := repository.Save(ctx, id, request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, request)
	}
}

func GetAgentHandler(
	store *store.AgentStore,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		state, err := store.Get(id)
		if err == false {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "agent not found",
			})
			return
		}

		online := time.Since(state.LastSeen) < time.Second*15

		response := models.AgentState{
			Metrics:  state.Metrics,
			LastSeen: state.LastSeen,
			Online:   online,
		}

		c.JSON(http.StatusOK, response)
	}
}

func GetAgentsHandler(store *store.AgentStore) gin.HandlerFunc {
	return func(c *gin.Context) {

		agents := store.GetAll()

		response := make(map[string]models.AgentState)

		for id, agent := range agents {
			online := time.Since(agent.LastSeen) < time.Second*15
			agent.Online = online
			response[id] = agent
		}

		c.JSON(http.StatusOK, response)
	}
}

func GetAgentHistoryHandler(
	repository *repository.MetricsRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		limit := 100

		if rawLimit := c.Query("limit"); rawLimit != "" {
			parsedLimit, err := strconv.Atoi(rawLimit)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid limit",
				})
				return
			}

			limit = parsedLimit
		}

		history, err := repository.GetHistory(
			c.Request.Context(),
			id,
			limit,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not load history",
			})
			return
		}

		c.JSON(http.StatusOK, history)
	}
}
