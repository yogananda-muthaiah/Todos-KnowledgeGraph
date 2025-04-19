package main

import (
	"fmt"
	"go-knowledge-graph/internal/api"
	"go-knowledge-graph/internal/graph"
)

func main() {
	todos, err := api.FetchTodos()
	if err != nil {
		fmt.Printf("Error fetching todos: %v\n", err)
		return
	}

	graph.BuildKnowledgeGraph(todos)
}
