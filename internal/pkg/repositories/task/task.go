package task

import (
	"strconv"

	"github.com/jgbz/wasm-task/internal/pkg/models/request"
	"github.com/jgbz/wasm-task/internal/pkg/models/response"
	"github.com/jgbz/wasm-task/internal/pkg/repositories/config"
)

type TasksRepository struct {
	instance *config.ConfigRepository
}

func NewTasksRepository() *TasksRepository {
	return &TasksRepository{
		instance: config.GetConfigRepository(),
	}
}

func (r *TasksRepository) Get() (*response.TasksResponse, error) {

	sql := "SELECT * FROM tasks"

	results, err := r.instance.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	rspTask := new(response.TasksResponse)
	for results.Next() {
		task := new(response.TaskResponse)

		err = results.Scan(&task.Id, &task.Description, &task.Status)
		if err != nil {
			continue
		}
		rspTask.Tasks = append(rspTask.Tasks, task)
	}
	results.Close()
	return rspTask, nil
}

func (r *TasksRepository) Insert(rqtTask *request.TaskRequest) (*response.TaskResponse, error) {

	sql := `INSERT INTO tasks(description, status)
			VALUES ('` + rqtTask.Description + `', ` + rqtTask.Status + `);`

	res, err := r.instance.DB.Exec(sql)
	if err != nil {
		return nil, err
	}

	if lastId, err := res.LastInsertId(); err != nil {
		return nil, err
	} else {
		rsp := new(response.TaskResponse)
		rsp.Id = strconv.Itoa(int(lastId))
		rsp.Status = rqtTask.Status
		return rsp, nil
	}

}

func (r *TasksRepository) Update(rqtTask *request.TaskRequest) (*response.TaskResponse, error) {

	sql := "UPDATE tasks SET status = " + rqtTask.Status + `
			WHERE id = ` + rqtTask.Id

	if _, err := r.instance.DB.Exec(sql); err != nil {
		return nil, err
	} else {
		rsp := new(response.TaskResponse)
		rsp.Id = rqtTask.Id
		rsp.Status = rqtTask.Status
		return rsp, nil
	}
}

func (r *TasksRepository) Delete(rqtTask *request.TaskRequest) (*response.TaskResponse, error) {

	sql := "delete from tasks where id  = " + rqtTask.Id

	if _, err := r.instance.DB.Exec(sql); err != nil {
		return nil, err
	} else {
		rsp := new(response.TaskResponse)
		rsp.Id = rqtTask.Id
		rsp.Status = rqtTask.Status
		return rsp, nil
	}
}
