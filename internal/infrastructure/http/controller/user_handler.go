package controller

import (
	"net/http"
	"time"

	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/dto"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
	secretKey   string
}

func NewUserHandler(userUseCase *usecase.UserUseCase, secretKey string) *UserHandler {
	return &UserHandler{userUseCase: userUseCase, secretKey: secretKey}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CredentialsDTO true "User information"
// @Success 201 {object} dto.ResponseDTO "User registered successfully"
// @Failure 400 {object} dto.ErrorDTO "Invalid input"
// @Router /register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var credentials dto.CredentialsDTO
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := h.userUseCase.CreateUser(credentials.Username, credentials.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, dto.ResponseDTO{Message: "User created successfully"})
}

// Login godoc
// @Summary Login user
// @Description Login user and return JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body dto.CredentialsDTO true "User login credentials"
// @Success 200 {object} dto.TokenResponseDTO "JWT token"
// @Failure 401 {object} dto.ErrorDTO "Unauthorized"
// @Router /login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var credentials dto.CredentialsDTO
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	authUser, err := h.userUseCase.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       authUser.ID,
		"username": authUser.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.secretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, dto.TokenResponseDTO{Token: tokenString})
}
