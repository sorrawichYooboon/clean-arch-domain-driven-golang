package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/dto"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
)

type BookHandler struct {
	UseCase *usecase.BookUseCase
}

func NewBookHandler(uc *usecase.BookUseCase) *BookHandler {
	return &BookHandler{
		UseCase: uc,
	}
}

// GetAll godoc
// @Summary Get all books
// @Description Get all books from the database
// @Tags books
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Book
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (h *BookHandler) GetAll(c echo.Context) error {
	books, err := h.UseCase.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, books)
}

// Create godoc
// @Summary Create a new book
// @Description Create a new book in the database
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param book body dto.BookCreateDTO true "Book data"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (h *BookHandler) Create(c echo.Context) error {
	bookDTO := new(dto.BookCreateDTO)
	if err := c.Bind(bookDTO); err != nil {
		return err
	}

	err := h.UseCase.CreateBook(bookDTO.Title, bookDTO.Author, bookDTO.Category, bookDTO.PublishedYear)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "create book successfully"})
}

// Update godoc
// @Summary Update an existing book
// @Description Update the book with the given ID
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Book ID"
// @Param book body dto.BookCreateDTO true "Book data"
// @Success 200 {object} domain.Book
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [put]
func (h *BookHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.UseCase.GetBookByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	bookDTO := new(dto.BookCreateDTO)
	if err := c.Bind(&bookDTO); err != nil {
		return err
	}

	book = &domain.Book{
		ID:            book.ID,
		Title:         bookDTO.Title,
		Author:        bookDTO.Author,
		PublishedYear: bookDTO.PublishedYear,
		Category:      bookDTO.Category,
	}

	h.UseCase.UpdateBook(book)
	return c.JSON(http.StatusOK, book)
}

// Delete godoc
// @Summary Delete a book
// @Description Delete the book with the given ID
// @Tags books
// @Security ApiKeyAuth
// @Param id path int true "Book ID"
// @Success 204 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [delete]
func (h *BookHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.UseCase.DeleteBook(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "delete book successfully"})
}
