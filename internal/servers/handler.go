package servers

import (
	"os"
	_tderDomain "tder/internal/domain"
	_usersHttp "tder/internal/users/controllers"
	_usersRepository "tder/internal/users/repositories"
	_usersUsecase "tder/internal/users/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	v1 := s.App.Group("/v1")

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"message": "The server is running properly.",
		})
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")
	jwtService := _tderDomain.NewJWTTokenService(secretKey)

	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	passwordHasher := _tderDomain.NewBcryptPasswordHasher()
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository, passwordHasher, *jwtService)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
