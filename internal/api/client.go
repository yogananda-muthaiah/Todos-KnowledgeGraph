package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-knowledge-graph/internal/models"
)

const baseURL = "https://jsonplaceholder.typicode.com"

func FetchTodos() ([]models.Todo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/todos", baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var todos []models.Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}
