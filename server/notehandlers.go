package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mohaila/echo-notes-rest/model"
)

// GetAllNotes handler
func (s *Server) GetAllNotes(c echo.Context) error {
	n, err := s.NoteService.GetAllNotes()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	if n == nil {
		n = []*model.Note{}
	}

	return c.JSON(http.StatusOK, n)
}

// GetNote handler
func (s *Server) GetNote(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	n, err := s.NoteService.GetNote(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	if n == nil {
		return echo.NewHTTPError(http.StatusNoContent, "note not found")
	}

	return c.JSON(http.StatusOK, n)
}

// CreateNote handler
func (s *Server) CreateNote(c echo.Context) error {
	r := &struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}{}
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	n := &model.Note{
		Title:       r.Title,
		Description: r.Description,
	}
	err := s.NoteService.CreateNote(n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}

	return c.JSON(http.StatusCreated, map[string]int{"id": n.ID})
}

// UpdateNote handler
func (s *Server) UpdateNote(c echo.Context) error {
	n := &model.Note{}
	if err := c.Bind(n); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	err := s.NoteService.UpdateNote(n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}

	return c.NoContent(http.StatusNoContent)
}

// DeleteNote handler
func (s *Server) DeleteNote(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	err = s.NoteService.DeleteNote(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}

	return c.NoContent(http.StatusNoContent)
}
