package service

import "github.com/GuilhermeFerza/content-platform-api/internal/model"

type TaskService struct {
	tasks []model.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: []model.Task{
			{ID: 1, Title: "Aprender GO", Done: false},
			{ID: 2, Title: "Criar uma API", Done: true},
		},
	}
}

func (s *TaskService) GetAll() []model.Task {
	return s.tasks
}

func (s *TaskService) Create(task model.Task) model.Task {
	task.ID = len(s.tasks) + 1
	s.tasks = append(s.tasks, task)
	return task
}
