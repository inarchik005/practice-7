package storage

import (
	"sync"
	"task-api/internal/models"
)

type TaskStore struct {
	mu     sync.Mutex
	tasks  map[int]models.Task
	nextID int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *TaskStore) Create(task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	s.tasks[s.nextID] = task
	s.nextID++

	return task
}

func (s *TaskStore) GetAll() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []models.Task
	for _, t := range s.tasks {
		result = append(result, t)
	}
	return result
}

func (s *TaskStore) GetByID(id int) (models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.tasks[id]
	return t, ok
}

func (s *TaskStore) Update(id int, done bool) (models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.tasks[id]
	if !ok {
		return models.Task{}, false
	}

	t.Done = done
	s.tasks[id] = t
	return t, true
}
