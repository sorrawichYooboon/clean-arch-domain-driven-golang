package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/dto"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/pkg/response"
)

type AuthorHandlerImpl struct {
	authorUsecase usecase.AuthorUseCase
}

func NewAuthorHandler(authorUsecase usecase.AuthorUseCase) AuthorHandler {
	return &AuthorHandlerImpl{
		authorUsecase: authorUsecase,
	}
}

// GetAll godoc
// @Summary Get all authors
// @Description Get all authors from the database
// @Tags authors
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponse{data=[]dto.GetAllAuthorResponse}
// @Failure 500 {object} response.APIResponse
// @Router /authors [get]
func (h *AuthorHandlerImpl) GetAll(c echo.Context) error {
	authors, err := h.authorUsecase.GetAllAuthors()
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	authorsResp := []dto.GetAllAuthorResponse{}
	for _, author := range authors {
		authorsResp = append(authorsResp, dto.GetAllAuthorResponse{
			ID:   author.ID,
			Name: author.Name,
			Bio:  author.Bio,
		})
	}

	return response.Success(c, http.StatusOK, response.SuccessGetAllAuthor, authorsResp)
}

// Create godoc
// @Summary Create a new author
// @Description Create a new author in the database
// @Tags authors
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param author body dto.CreateAuthorRequest true "Author data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /authors [post]
func (h *AuthorHandlerImpl) Create(c echo.Context) error {
	var req dto.CreateAuthorRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	err := h.authorUsecase.CreateAuthor(req.Name, req.Bio)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, response.SuccessCreateAuthor, nil)
}

// Update godoc
// @Summary Update an existing author
// @Description Update the author with the given ID
// @Tags authors
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Author ID"
// @Param author body dto.UpdateAuthorRequest true "Updated Author data"
// @Success 200 {object} response.APIResponse{data=dto.UpdateAuthorResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /authors/{id} [put]
func (h *AuthorHandlerImpl) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	author, err := h.authorUsecase.GetAuthorByID(uint(id))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	var req dto.UpdateAuthorRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	author = &domain.Author{
		ID:   author.ID,
		Name: req.Name,
		Bio:  req.Bio,
	}

	err = h.authorUsecase.UpdateAuthor(author)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	authorResp := dto.UpdateAuthorResponse{
		ID:   author.ID,
		Name: author.Name,
		Bio:  author.Bio,
	}

	return response.Success(c, http.StatusOK, response.SuccessUpdateAuthor, authorResp)
}

// Delete godoc
// @Summary Delete an author
// @Description Delete the author with the given ID
// @Tags authors
// @Security ApiKeyAuth
// @Param id path int true "Author ID"
// @Success 204 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /authors/{id} [delete]
func (h *AuthorHandlerImpl) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.authorUsecase.DeleteAuthor(uint(id)); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, response.SuccessDeleteAuthor, nil)
}
