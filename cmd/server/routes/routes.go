package routes

import (
	"github.com/jgbz/wasm-task/cmd/server/routes/handlers"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	App     *fiber.App
	Handler *handlers.Handler
}

func NewRoutes() *Routes {
	return &Routes{
		App:     fiber.New(),
		Handler: handlers.NewHandlers(),
	}
}

func (r *Routes) RegisterRoutesGet(group string) {
	api := r.App.Group(group)
	api.Get("tasks", r.Handler.GetTasks)
	api.Post("tasks", r.Handler.NewTask)
	api.Patch("tasks", r.Handler.UpdateTask)
	api.Delete("tasks", r.Handler.DeleteTask)
}
