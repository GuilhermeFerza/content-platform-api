package service

import (
	"testing"

	"github.com/GuilhermeFerza/content-platform-api/internal/model"
)

func TestTaskService_GetAll(t *testing.T) {
	service := NewTaskService()

	tasks := service.GetAll()

	if len(tasks) != 2 {
		t.Error("Expected 2 tasks, got", len(tasks))
	}
}

func TestTaskService_Create(t *testing.T) {
	service := NewTaskService()

	task := model.Task{
		Title: "Nova Tarefa",
		Done:  false,
	}

	created := service.Create(task)

	if created.ID != 3 {
		t.Errorf("expected ID 3, got %d", created.ID)
	}

	if len(service.GetAll()) != 3 {
		t.Errorf("expected 3 tasks, got %d", len(service.GetAll()))
	}
}

func TestTaskService_DeleteByID(t *testing.T) {
	service := NewTaskService()

	ok := service.DeleteByID(1)
	if !ok {
		t.Errorf("expected task to be deleted")
	}

	if len(service.GetAll()) != 1 {
		t.Errorf("expected 1 task remaining, got %d", len(service.GetAll()))
	}
}
