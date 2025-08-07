package service

import (
	"fmt"
	"strings"
	"todo_api/model"
	"todo_api/store"
)

type TodoService struct {
	store *store.TodoStore
}

func NewTodoService(store *store.TodoStore) *TodoService {
	return &TodoService{store: store}
}

func (s *TodoService) CreateTodo(req model.CreateTodoRequest) (*model.Todo, error) {
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	todo := s.store.Create(req)
	return todo, nil
}

func (s *TodoService) GetAllTodos() []*model.Todo {
	return s.store.GetAll()
}

func (s *TodoService) GetTodoByID(id int) (*model.Todo, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid todo id")
	}

	return s.store.GetByID(id)
}

func (s *TodoService) UpdateTodo(id int, req model.UpdateTodoRequest) (*model.Todo, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid todo id")
	}

	if req.Title != nil {
		*req.Title = strings.TrimSpace(*req.Title)
		if *req.Title == "" {
			return nil, fmt.Errorf("title cannot be empty")
		}
	}

	return s.store.Update(id, req)
}

func (s *TodoService) DeleteTodo(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid todo id")
	}

	return s.store.Delete(id)
}
