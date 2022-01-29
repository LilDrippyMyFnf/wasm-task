package services

import (
	"context"
	"github.com/jgbz/wasm-task/internal/pkg/models/request"
	"github.com/jgbz/wasm-task/internal/pkg/models/response"
	"github.com/jgbz/wasm-task/internal/pkg/repositories/task"
	"net/http"
)

type TasksService struct {
}

func NewTasksService() *TasksService {
	return &TasksService{}
}

func (s *TasksService) GetTasks() (*response.TasksResponse, *response.ErrorResponse) {

	r := task.NewTasksRepository()
	rsp, err := r.Get()
	if err != nil {
		return nil, response.NewErrorResponse(http.StatusBadRequest, err.Error())
	}
	return rsp, nil
}

func (s *TasksService) NewTask(ctx context.Context, rqtTask *request.TaskRequest) (*response.TaskResponse, *response.ErrorResponse) {
	if len(rqtTask.Description) < 1 {
		return nil, response.NewErrorResponse(-1, "Task description cannot be empty!")
	}

	r := task.NewTasksRepository()
	rsp, err := r.Insert(rqtTask)
	if err != nil {
		return nil, response.NewErrorResponse(http.StatusBadRequest, err.Error())
	}
	return rsp, nil
}

func (s *TasksService) UpdateTask(ctx context.Context, rqtTask *request.TaskRequest) (*response.TaskResponse, *response.ErrorResponse) {

	r := task.NewTasksRepository()
	rsp, err := r.Update(rqtTask)
	if err != nil {
		return nil, response.NewErrorResponse(http.StatusBadRequest, err.Error())
	}
	return rsp, nil
}

func (s *TasksService) DeleteTask(ctx context.Context, rqtTask *request.TaskRequest) (*response.TaskResponse, *response.ErrorResponse) {

	r := task.NewTasksRepository()
	rsp, err := r.Delete(rqtTask)
	if err != nil {
		return nil, response.NewErrorResponse(http.StatusBadRequest, err.Error())
	}
	return rsp, nil
}
