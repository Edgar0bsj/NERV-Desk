package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgar0bsj/nerv-desk/module/ticket/database"
	"github.com/edgar0bsj/nerv-desk/module/ticket/handler"
	"github.com/edgar0bsj/nerv-desk/module/ticket/repository"
	"github.com/edgar0bsj/nerv-desk/module/ticket/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTicketHandler_FindAllTicketsUser(t *testing.T) {
	// gin modo teste
	gin.SetMode(gin.TestMode)

	// setup
	db, _ := database.New()
	repo := repository.New(db)
	svc := service.New(repo)
	handler := handler.New(svc)

	// iniciando rota teste
	router := gin.New()

	// Middler passando userId
	router.Use(func(c *gin.Context) {
		c.Set("userID", "536a819d-0e99-4cdd-b124-ac8abbb7e6e4")
		c.Next()
	})
	router.GET("/ticket/user", handler.FindAllTicketsUser)

	// Configurando requisição

	req := httptest.NewRequest(
		http.MethodGet,
		"/ticket/user",
		nil,
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Capturando a resposta
	assert.Equal(t, http.StatusOK, rec.Code)
	t.Log(rec.Body.String())
}
