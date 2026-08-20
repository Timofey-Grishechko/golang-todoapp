package tasks_service

import (
	"context"

	"github.com/Timofey-Grishechko/golang-todoapp/internal/core/domain"
)

type TasksService struct {
	taskRepository TaskRepository
}

type TaskRepository interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)

	GetTasks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Task, error)

	GetTask(
		ctx context.Context,
		id int,
	) (domain.Task, error)

	DeleteTask(
		ctx context.Context,
		id int,
	) error

	PatchTask(
		ctx context.Context,
		id int,
		task domain.Task,
	) (domain.Task, error)
}

func NewTaskService(
	taskRepository TaskRepository,
) *TasksService {
	return &TasksService{
		taskRepository: taskRepository,
	}
}
