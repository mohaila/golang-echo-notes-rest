package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mohaila/echo-notes-rest/config"
	"github.com/mohaila/echo-notes-rest/service"
)

type (
	// Server for REST note API
	Server struct {
		*config.ServerConfig
		*echo.Echo
		service.NoteService
	}
)

// NewServer creates a new server
func NewServer(c *config.ServerConfig, ns service.NoteService) *Server {
	s := &Server{c, echo.New(), ns}

	s.Logger.SetLevel(log.Lvl(log.DEBUG))
	s.Pre(middleware.RemoveTrailingSlash())
	s.HideBanner = true
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())

	// Note handlers
	s.GET("/api/notes", s.GetAllNotes)
	s.GET("/api/notes/:id", s.GetNote)
	s.POST("/api/notes", s.CreateNote)
	s.PUT("/api/notes/:id", s.UpdateNote)
	s.DELETE("/api/notes/:id", s.DeleteNote)

	return s
}

func (s *Server) Start() error {
	a := fmt.Sprintf("%s:%d", s.Address, s.Port)
	return s.Echo.Start(a)
}
