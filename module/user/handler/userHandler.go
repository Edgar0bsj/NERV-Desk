package userhandler

import (
	"net/http"
	"time"

	"github.com/google/uuid"

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

func (h *UserHandler) UserRegister(c *gin.Context) {
	validation := validator.New()
	var req *service.UserRegisterDto

	// 1. Receber e validar os dados do novo usuário.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 2. Caso os dados sejam inválidos, retornar Bad Request.
	if err := validation.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 3. Verificar se o usuário já existe.
	user, _ := h.svc.FindByEmail(req.Email)

	// 4. Caso o usuário já exista, impedir o cadastro e retornar erro.
	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already registered"})
		return
	}

	// 5. Criar o novo usuário no banco de dados.
	user, err := h.svc.CreateUser(req)

	// 6. Caso ocorra um erro durante o cadastro, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// 7. Retornar o usuário criado com status de sucesso.
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *UserHandler) ListUser(c *gin.Context) {
	// 1. Buscar todos os usuários cadastrados no banco de dados.
	users, err := h.svc.FindAllUser()

	// 2. Caso ocorra um erro durante a busca, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// 3. Retornar a lista de usuários com status de sucesso.
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) EditUser(c *gin.Context) {
	// 1. Obter o ID do usuário a partir dos parâmetros da rota.
	userId := c.Param("id")

	// 2. Receber e validar os dados enviados para atualização.
	var req *service.UserUpdateDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 3. Caso o ID ou os dados sejam inválidos, retornar Bad Request.
	if _, err := uuid.Parse(userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter"})
		return
	}

	// 4. Buscar o usuário pelo ID.
	user, err := h.svc.FindById(userId)

	// 5. Caso o usuário não exista, retornar Not Found.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// 6. Atualizar os dados do usuário.
	newUser, err := h.svc.UserUpdate(user, req)

	// 7. Caso ocorra um erro durante a atualização, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// 8. Retornar o usuário atualizado com status de sucesso.
	c.JSON(http.StatusOK, gin.H{"user": newUser})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// 1. Obter o ID do usuário a partir dos parâmetros da rota.
	userId := c.Param("id")

	// 2. Validar o ID do usuário.
	if _, err := uuid.Parse(userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter"})
		return
	}

	// 4. Buscar o usuário pelo ID.
	user, err := h.svc.FindById(userId)

	// 5. Caso o usuário não exista, retornar Not Found.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// 6. Excluir o usuário.
	err = h.svc.DeleteUser(user.ID)

	// 7. Caso ocorra um erro durante a exclusão, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// 8. Retornar resposta de sucesso confirmando a exclusão.
	c.JSON(http.StatusOK, gin.H{"msg": "Usuario deletado com sucesso!"})
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	// 1. Obter o ID do usuário a partir dos parâmetros da rota.
	userId := c.Param("id")

	// 2. Receber e validar a nova senha.
	validation := validator.New()
	var req *service.UserPasswordChange

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in the request body"})
		return
	}

	if err := validation.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	// 3. Caso o ID ou a nova senha sejam inválidos, retornar Bad Request.
	if _, err := uuid.Parse(userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter"})
		return
	}

	// 4. Buscar o usuário pelo ID.
	user, err := h.svc.FindById(userId)

	// 5. Caso o usuário não exista, retornar Not Found.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// 6. Codificar a nova senha utilizando bcrypt.
	passhash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error in hash creation"})
		return
	}

	// 7. Atualizar a senha do usuário com a senha codificada.
	err = h.svc.UserUpdatePasswod(user, passhash)

	// 8. Caso ocorra um erro durante a codificação ou atualização, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error in hash creation"})
		return
	}
	// 9. Retornar resposta de sucesso confirmando a alteração da senha.
	c.JSON(http.StatusOK, gin.H{"msg": "senha atualizada com sucesso!"})
}
