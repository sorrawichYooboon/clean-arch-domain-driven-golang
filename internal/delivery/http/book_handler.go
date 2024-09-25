package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
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

func (h *BookHandler) GetAll(c echo.Context) error {
	books, err := h.UseCase.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) Create(c echo.Context) error {
	book := new(domain.Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	err := h.UseCase.CreateBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.UseCase.GetBookByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	if err := c.Bind(&book); err != nil {
		return err
	}

	h.UseCase.UpdateBook(book)
	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.UseCase.DeleteBook(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
