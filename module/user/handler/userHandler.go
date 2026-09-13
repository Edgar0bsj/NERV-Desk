package userhandler

import (
	"net/http"

	"github.com/edgar0bsj/nerv-desk/module/user/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserServiceInterface
}

func New(useCases service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		svc: useCases,
	}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.svc.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error do servidor",
		})
		return
	}

	c.JSON(http.StatusOK, users)

}

// Criar um usuario
func (h *UserHandler) UserRegister(c *gin.Context) {
	var req service.UserRegisterDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "corpo da solicitação inválido",
		})
		return
	}

	err := h.svc.SaveUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "falhas ao salvar",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "Usuario salvo com sucesso!",
	})

}

// Editar Usuario
func (h *UserHandler) UserUpdate(c *gin.Context) {
	userUpdate := service.UserUpdateDto{
		ID: c.Param("id"),
	}

	err := c.ShouldBindJSON(&userUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error no corpo da requisição",
		})
		return
	}

	err = h.svc.UpdateUser(&userUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error ao atualizar usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Usuario atualizado com sucesso!",
	})
}

// Deletar usuario
func (h *UserHandler) UserDelete(c *gin.Context) {
	user_id := c.Param("id")

	if err := h.svc.DeleteUser(user_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error ao Deletar usuario",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Atualizar senha
func (h *UserHandler) UserUpdatePassword(c *gin.Context) {
	var newPassword service.UserUpdatePassword
	user_id := c.Param("id")

	if err := c.ShouldBindJSON(&newPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error no corpo da requisição"})
		return
	}

	if err := h.svc.UpdatePassword(user_id, newPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error ao Atualizar Senha"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Senha atualizada com sucesso!"})
}
