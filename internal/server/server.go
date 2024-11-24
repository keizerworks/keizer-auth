package server

import (
	"keizer-auth/internal/app"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	container   *app.Container
	controllers *app.ServerControllers
	middlewars  *app.ServerMiddlewares
}

func New() *FiberServer {
	container := app.GetContainer()
	controllers := app.GetControllers(container)
	middlewars := app.GetMiddlewares(container)

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "keizer-auth",
			AppName:      "keizer-auth",
		}),
		container:   container,
		controllers: controllers,
		middlewars:  middlewars,
	}

	return server
}

func (s *FiberServer) Shutdown() error {
	if err := s.App.Shutdown(); err != nil {
		return err
	}

	return s.container.Cleanup()
}
