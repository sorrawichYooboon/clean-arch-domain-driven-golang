package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
)

type AuthorHandler struct {
	UseCase *usecase.AuthorUseCase
}

func NewAuthorHandler(uc *usecase.AuthorUseCase) *AuthorHandler {
	return &AuthorHandler{
		UseCase: uc,
	}
}

func (h *AuthorHandler) GetAll(c echo.Context) error {
	authors, err := h.UseCase.GetAllAuthors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, authors)
}

func (h *AuthorHandler) Create(c echo.Context) error {
	author := new(domain.Author)
	if err := c.Bind(author); err != nil {
		return err
	}

	err := h.UseCase.CreateAuthor(author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	author, err := h.UseCase.GetAuthorByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Author not found"})
	}

	if err := c.Bind(&author); err != nil {
		return err
	}

	h.UseCase.UpdateAuthor(author)
	return c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.UseCase.DeleteAuthor(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
