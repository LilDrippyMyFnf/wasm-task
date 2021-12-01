package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	// SvrPessoas  *services.PessoasService Aqui vai os serviços
}

// Função que retorna um handler
func NewHandlers() *Handler {
	return &Handler{}
}

func (h *Handler) NewTask(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON("HelloWorld")
}

func (h *Handler) GetTasks(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON("HelloWorld")
}
func (h *Handler) UpdateTask(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON("HelloWorld")
}
func (h *Handler) DeleteTask(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON("HelloWorld")
}
