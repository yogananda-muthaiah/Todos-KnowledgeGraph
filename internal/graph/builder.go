package graph

import (
	"fmt"
	"go-knowledge-graph/internal/models"
)

func BuildKnowledgeGraph(todos []models.Todo) {
	userTodos := make(map[int][]models.Todo)

	for _, todo := range todos {
		userTodos[todo.UserID] = append(userTodos[todo.UserID], todo)
	}

	for userID, todos := range userTodos {
		fmt.Printf("User %d:\n", userID)
		for _, todo := range todos {
			status := "Incomplete"
			if todo.Completed {
				status = "Completed"
			}
			fmt.Printf("  - [%s] %s\n", status, todo.Title)
		}
	}
}
