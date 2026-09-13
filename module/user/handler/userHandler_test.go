package userhandler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edgar0bsj/nerv-desk/module/user/database"
	userhandler "github.com/edgar0bsj/nerv-desk/module/user/handler"
	"github.com/edgar0bsj/nerv-desk/module/user/repository"
	"github.com/edgar0bsj/nerv-desk/module/user/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Create_user(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, _ := database.New()
	repo := repository.New(db)
	svc := service.New(repo)

	handler := userhandler.New(svc)

	router := gin.New()
	router.POST("/user", handler.UserRegister)

	body := `{
        "name": "Edgar teste",
        "email": "edgarjunior2@email.com",
        "password": "admin123",
        "role": "ADMIN"
    }`

	req := httptest.NewRequest(
		http.MethodPost,
		"/user",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// Status code & body
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "Usuario salvo com sucesso!")

}

func TestUserHandler_ListUsers(t *testing.T) {
	// gin modo teste
	gin.SetMode(gin.TestMode)

	// setup
	db, _ := database.New()
	repo := repository.New(db)
	svc := service.New(repo)
	handler := userhandler.New(svc)

	// iniciando rota teste
	router := gin.New()
	router.GET("/user", handler.ListUsers)

	// Configurando requisição
	req := httptest.NewRequest(
		http.MethodGet,
		"/user",
		nil,
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Capturando a resposta
	assert.Equal(t, http.StatusOK, rec.Code)
	t.Log(rec.Body.String())
}

func TestUserHandler_Update(t *testing.T) {
	// gin modo teste
	gin.SetMode(gin.TestMode)

	// setup
	db, _ := database.New()
	repo := repository.New(db)
	svc := service.New(repo)
	handler := userhandler.New(svc)

	// iniciando rota teste
	router := gin.New()
	router.PUT("/user/:id", handler.UserUpdate)

	// Configurando requisição
	body := `{
        "name": "Marcelo vilela ATUALIZADO",
        "email": "macelinho@email.com",
        "role": "ADMIN"
    }`
	req := httptest.NewRequest(
		http.MethodPut,
		"/user/872e35b5-5464-401d-8376-59ee5283b276",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Capturando a resposta
	assert.Equal(t, http.StatusOK, rec.Code)
	t.Log(rec.Body.String())
}

func TestUserHandler_Delete(t *testing.T) {
	// gin modo teste
	gin.SetMode(gin.TestMode)

	// setup
	db, _ := database.New()
	repo := repository.New(db)
	svc := service.New(repo)
	handler := userhandler.New(svc)

	// iniciando rota teste
	router := gin.New()
	router.DELETE("/user/:id", handler.UserDelete)

	// Configurando requisição
	req := httptest.NewRequest(
		http.MethodDelete,
		"/user/872e35b5-5464-401d-8376-59ee5283b276",
		nil,
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Capturando a resposta
	assert.Equal(t, http.StatusNoContent, rec.Code)
	t.Log(rec.Body.String())
}

func TestUserHandler_UpdatePassword(t *testing.T) {
	// gin modo teste
	gin.SetMode(gin.TestMode)

	// setup
	db, _ := database.New()
	repo := repository.New(db)
	svc := service.New(repo)
	handler := userhandler.New(svc)

	// iniciando rota teste
	router := gin.New()
	router.PATCH("/user/passchange/:id", handler.UserUpdatePassword)

	// Configurando requisição
	body := `{
        "Password": "Segredão"
    }`
	req := httptest.NewRequest(
		http.MethodPatch,
		"/user/passchange/231eac0c-64b2-41a6-8e70-c1805e3577bc",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Capturando a resposta
	assert.Equal(t, http.StatusOK, rec.Code)
	t.Log(rec.Body.String())
}
