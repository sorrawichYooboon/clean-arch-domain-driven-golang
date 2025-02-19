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

type BookHandlerImpl struct {
	bookUsecase usecase.BookUseCase
}

func NewBookHandler(bookUsecase usecase.BookUseCase) BookHandler {
	return &BookHandlerImpl{
		bookUsecase: bookUsecase,
	}
}

// GetAll godoc
// @Summary Get all books
// @Description Get all books from the database
// @Tags books
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponse{data=[]dto.GetAllBookResponse}
// @Failure 500 {object} response.APIResponse
// @Router /books [get]
func (h *BookHandlerImpl) GetAll(c echo.Context) error {
	books, err := h.bookUsecase.GetAllBooks()
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	booksResp := []dto.GetAllBookResponse{}
	for _, book := range books {
		booksResp = append(booksResp, dto.GetAllBookResponse{
			ID:            book.ID,
			Title:         book.Title,
			Author:        book.Author,
			PublishedYear: book.PublishedYear,
			Category:      book.Category,
		})
	}

	return response.Success(c, http.StatusOK, response.SuccessGetAllBook, booksResp)
}

// Create godoc
// @Summary Create a new book
// @Description Create a new book in the database
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param book body dto.CreateBookRequest true "Book data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /books [post]
func (h *BookHandlerImpl) Create(c echo.Context) error {
	var req dto.CreateBookRequest
	if err := c.Bind(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	err := h.bookUsecase.CreateBook(req.Title, req.Author, req.Category, req.PublishedYear)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, response.SuccessCreateBook, nil)
}

// Update godoc
// @Summary Update an existing book
// @Description Update the book with the given ID
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Book ID"
// @Param book body dto.UpdateBookRequest true "Updated Book data"
// @Success 200 {object} response.APIResponse{data=dto.UpdateBookResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /books/{id} [put]
func (h *BookHandlerImpl) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.bookUsecase.GetBookByID(uint(id))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	var req dto.UpdateBookRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	book = &domain.Book{
		ID:            book.ID,
		Title:         req.Title,
		Author:        req.Author,
		PublishedYear: req.PublishedYear,
		Category:      req.Category,
	}

	err = h.bookUsecase.UpdateBook(book)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, response.SuccessUpdateBook, nil)
}

// Delete godoc
// @Summary Delete a book
// @Description Delete the book with the given ID
// @Tags books
// @Security ApiKeyAuth
// @Param id path int true "Book ID"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /books/{id} [delete]
func (h *BookHandlerImpl) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.bookUsecase.DeleteBook(uint(id)); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, response.SuccessDeleteBook, nil)
}
