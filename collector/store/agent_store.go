package store

import (
	"collector/models"
	"log"
	"sync"
	"time"
)

type AgentStore struct {
	mu      sync.RWMutex
	metrics map[string]models.AgentState
}

func NewAgentStore() *AgentStore {
	return &AgentStore{metrics: make(map[string]models.AgentState)}
}

func (s *AgentStore) Set(id string, metrics models.AgentMetrics) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.metrics[id] = models.AgentState{
		Metrics:  metrics,
		LastSeen: time.Now(),
	}

	log.Println("Set")
}

func (s *AgentStore) Get(
	id string,
) (models.AgentState, bool) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	state, ok := s.metrics[id]
	if !ok {
		log.Println("Could not get metrics")
		return models.AgentState{}, false
	}

	return state, true
}

func (s *AgentStore) GetAll() map[string]models.AgentState {
	s.mu.RLock()
	defer s.mu.RUnlock()

	response := make(map[string]models.AgentState)

	for id, state := range s.metrics {
		response[id] = state
	}

	return response

}
