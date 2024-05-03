package servers

import (
	"log"

	"tder/configs"
	"tder/pkg/utils"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
	Cfg *configs.Configs
	Db  *gorm.DB
}

func NewServer(cfg *configs.Configs, db *gorm.DB) *Server {
	return &Server{
		App: fiber.New(),
		Cfg: cfg,
		Db:  db,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.Connection("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
