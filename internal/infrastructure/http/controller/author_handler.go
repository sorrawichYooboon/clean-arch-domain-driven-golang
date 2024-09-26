package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/dto"
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

// GetAll godoc
// @Summary Get all authors
// @Description Get all authors from the database
// @Tags authors
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Author
// @Failure 500 {object} map[string]string
// @Router /authors [get]
func (h *AuthorHandler) GetAll(c echo.Context) error {
	authors, err := h.UseCase.GetAllAuthors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, authors)
}

// Create godoc
// @Summary Create a new author
// @Description Create a new author in the database
// @Tags authors
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param author body dto.AuthorCreateDTO true "Author data"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authors [post]
func (h *AuthorHandler) Create(c echo.Context) error {
	authorDTO := new(dto.AuthorCreateDTO)
	if err := c.Bind(&authorDTO); err != nil {
		return err
	}

	err := h.UseCase.CreateAuthor(authorDTO.Name, authorDTO.Bio)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "create author successfully"})
}

// Update godoc
// @Summary Update an existing author
// @Description Update the author with the given ID
// @Tags authors
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Author ID"
// @Param author body dto.AuthorCreateDTO true "Author data"
// @Success 200 {object} domain.Author
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authors/{id} [put]
func (h *AuthorHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	author, err := h.UseCase.GetAuthorByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Author not found"})
	}

	authorDTO := new(dto.AuthorCreateDTO)
	if err := c.Bind(&authorDTO); err != nil {
		return err
	}

	author = &domain.Author{
		ID:   author.ID,
		Name: authorDTO.Name,
		Bio:  authorDTO.Bio,
	}

	h.UseCase.UpdateAuthor(author)
	return c.JSON(http.StatusOK, author)
}

// Delete godoc
// @Summary Delete an author
// @Description Delete the author with the given ID
// @Tags authors
// @Security ApiKeyAuth
// @Param id path int true "Author ID"
// @Success 204 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authors/{id} [delete]
func (h *AuthorHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.UseCase.DeleteAuthor(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "delete author successfully"})
}
