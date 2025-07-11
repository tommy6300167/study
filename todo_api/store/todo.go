package store

import (
	"fmt"
	"sort"
	"sync"
	"time"
	"todo_api/model"
)

type TodoStore struct {
	mu     sync.RWMutex
	todos  map[int]*model.Todo
	nextID int
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]*model.Todo),
		nextID: 1,
	}
}

func (s *TodoStore) Create(req model.CreateTodoRequest) *model.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := &model.Todo{
		ID:          s.nextID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.todos[s.nextID] = todo
	s.nextID++

	return todo
}

func (s *TodoStore) GetAll() []*model.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*model.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].CreatedAt.Before(todos[j].CreatedAt)
	})

	return todos
}

func (s *TodoStore) GetByID(id int) (*model.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, fmt.Errorf("todo with id %d not found", id)
	}

	return todo, nil
}

func (s *TodoStore) Update(id int, req model.UpdateTodoRequest) (*model.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, fmt.Errorf("todo with id %d not found", id)
	}

	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}

	todo.UpdatedAt = time.Now()

	return todo, nil
}

func (s *TodoStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.todos[id]
	if !exists {
		return fmt.Errorf("todo with id %d not found", id)
	}

	delete(s.todos, id)
	return nil
}
