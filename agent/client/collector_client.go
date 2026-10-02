package client

import (
	"agent/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendMetrics(collectorURL string,
	agentId string,
	metrics models.AgentMetrics) error {

	jsonMetrics, err := json.Marshal(metrics)
	if err != nil {
		log.Println("Could not format metrics")
		return err
	}

	url := fmt.Sprintf(
		"%s/agents/%s/metrics",
		collectorURL,
		agentId,
	)

	requestBody := bytes.NewBuffer(jsonMetrics)

	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		log.Println("Could not create a request")
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "apikey")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Could not send request:")
		return err
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Println("Collector returned status:", response.StatusCode)
		return err
	}

	return nil
}
