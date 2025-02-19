package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/dto"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/pkg/response"
)

type UserHandlerImpl struct {
	userUseCase usecase.UserUseCase
	secretKey   string
}

func NewUserHandler(userUseCase usecase.UserUseCase, secretKey string) UserHandler {
	return &UserHandlerImpl{userUseCase: userUseCase, secretKey: secretKey}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserRegisterRequest true "User registration data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /register [post]
func (h *UserHandlerImpl) Register(c echo.Context) error {
	var req dto.UserRegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	if err := h.userUseCase.CreateUser(req.Username, req.Password); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, response.SuccessRegisterUser, nil)
}

// Login godoc
// @Summary Authenticate user
// @Description Authenticate the user and return a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body dto.UserLoginRequest true "User login credentials"
// @Success 200 {object} response.APIResponse{data=dto.UserLoginResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /login [post]
func (h *UserHandlerImpl) Login(c echo.Context) error {
	var req dto.UserLoginRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	authUser, err := h.userUseCase.Authenticate(req.Username, req.Password)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       authUser.ID,
		"username": authUser.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.secretKey))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, response.SuccessLoginUser, dto.UserLoginResponse{Token: tokenString})
}
