package userhandler

import (
	"net/http"
	"time"

	"github.com/edgar0bsj/nerv-desk/module/user/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	svc *service.UserService
}

func New(useCases *service.UserService) *UserHandler {
	return &UserHandler{
		svc: useCases,
	}
}

func (h *UserHandler) UserLogin(c *gin.Context) {
	// 0. Capturar corpo da requesição
	var req service.UserLoginDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// 1. Validar campos.
	validation := validator.New()
	if err := validation.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// 2. Buscar o usuário pelo e-mail.
	user, err := h.svc.FindByEmail(req.Email)

	// 3. Caso não exista, retornar erro de autenticação 401 Unauthorized.
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication error",
		})
		return
	}

	// 4. Validar a senha informada.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication error",
		})
		return
	}
	// 5. Gerar o JWT com o ID e a role do usuário.
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // expira em 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	assinaturaToken, err := token.SignedString([]byte("nead_desk_secret")) //Senha provisória

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error generating token"})
		return
	}
	// 6. Retornar o JWT.
	c.JSON(http.StatusCreated, gin.H{
		"access_token": assinaturaToken,
		"token_type":   "Bearer",
	})

}
