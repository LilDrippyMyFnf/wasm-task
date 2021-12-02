package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jgbz/wasm-task/internal/pkg/models/request"
	"github.com/jgbz/wasm-task/internal/pkg/services"
)

type Handler struct {
	// SvrPessoas  *services.PessoasService Aqui vai os serviços
	TaskService *services.TasksService
}

// Função que retorna um handler
func NewHandlers() *Handler {
	return &Handler{}
}

func (h *Handler) NewTask(c *fiber.Ctx) error {
	ctx := c.Context()
	data := c.Request().Body()

	var rqtTask *request.TaskRequest
	err := json.Unmarshal(data, &rqtTask)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	rsp, er := h.TaskService.NewTask(ctx, rqtTask)
	if er != nil {
		return c.Status(er.Status).JSON(er)
	}
	return c.Status(http.StatusOK).JSON(rsp)

}

func (h *Handler) GetTasks(c *fiber.Ctx) error {

	rsp, er := h.TaskService.GetTasks()
	if er != nil {
		return c.Status(er.Status).JSON(er)
	}
	return c.Status(http.StatusOK).JSON(rsp)
}

func (h *Handler) UpdateTask(c *fiber.Ctx) error {

	ctx := c.Context()
	data := c.Request().Body()
	fmt.Println(string(data))

	var rqtTask *request.TaskRequest
	err := json.Unmarshal(data, &rqtTask)
	if err != nil {
		return err
	}

	rsp, er := h.TaskService.UpdateTask(ctx, rqtTask)
	if er != nil {
		return c.Status(er.Status).JSON(er)
	}
	return c.Status(http.StatusOK).JSON(rsp)

}
func (h *Handler) DeleteTask(c *fiber.Ctx) error {

	ctx := c.Context()
	data := c.Request().Body()
	fmt.Println(string(data))

	var rqtTask *request.TaskRequest
	err := json.Unmarshal(data, &rqtTask)
	if err != nil {
		return err
	}

	rsp, er := h.TaskService.DeleteTask(ctx, rqtTask)
	if er != nil {
		return c.Status(er.Status).JSON(er)
	}
	return c.Status(http.StatusOK).JSON(rsp)
}
