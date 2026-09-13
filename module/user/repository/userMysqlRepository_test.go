package repository_test

import (
	"testing"
	"time"

	"github.com/edgar0bsj/nerv-desk/module/user/database"
	"github.com/edgar0bsj/nerv-desk/module/user/model"
	"github.com/edgar0bsj/nerv-desk/module/user/repository"
	"github.com/google/uuid"
)

// ==============================================
// TESTE DE PERSISTENCIA
// ==============================================
func TestUserMysqlRepository_Save(t *testing.T) {
	var user model.User

	db, err := database.New()

	if err != nil {
		t.Fatalf("erro ao conectar no banco: %v", err)
	}

	repository := repository.New(db)

	user = model.User{
		ID:            uuid.New().String(),
		Name:          "Edgar teste",
		Email:         "Edgar@email.com",
		Password_hash: "hashash1234",
		Role:          model.RoleUser,
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}

	err = repository.Save(&user)

	if err != nil {
		t.Fatalf("esperava nil, recebeu: %v", err)
	}

}

// ==============================================
// TESTE DE LISTAR USERS
// ==============================================
func TestUserMysqlRepository_FindAll(t *testing.T) {

	db, err := database.New()

	if err != nil {
		t.Fatalf("Erro ao se conectar")
	}

	repository := repository.New(db)

	users, err := repository.FindAll()

	if err != nil {
		t.Fatalf("Error ao instancia o repository")
	}

	t.Log("v%", users)

}

// ==============================================
// TESTE DE BUSCAR USER
// ==============================================
func TestUserMysqlRepository_FindById(t *testing.T) {

	db, err := database.New()

	if err != nil {
		t.Fatalf("Erro ao se conectar")
	}

	repository := repository.New(db)

	users, err := repository.FindByID("692fe79e-c315-4187-bf5b-cd76742376b7")

	if err != nil {
		t.Fatalf("Error ao buscar user")
	}

	t.Log("v%", users)

}

// ==============================================
// TESTE DE ATUALIZAR USER
// ==============================================
func TestUserMysqlRepository_Update(t *testing.T) {

	db, err := database.New()

	if err != nil {
		t.Fatalf("Erro ao se conectar")
	}

	repository := repository.New(db)

	user, err := repository.FindByID("692fe79e-c315-4187-bf5b-cd76742376b7")
	if err != nil {
		t.Fatalf("Error ao buscar user")
	}

	user.Name = "Jorel da silva"
	user.Role = model.RoleAttendant

	if err := repository.Update(user); err != nil {
		t.Fatalf("Erro ao atualizar o usuario")
	}

	t.Log("v%", user)

}

// ==============================================
// TESTE DE DELETAR USER
// ==============================================
func TestUserMysqlRepository_Delete(t *testing.T) {

	db, err := database.New()
	if err != nil {
		t.Fatalf("Erro ao se conectar")
	}

	repository := repository.New(db)

	if err := repository.Delete("692fe79e-c315-4187-bf5b-cd76742376b7"); err != nil {
		t.Fatalf("Error ao buscar user")
	}

	t.Log("Deletado com sucesso!")

}
