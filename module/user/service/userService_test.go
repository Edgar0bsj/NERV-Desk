package service_test

import (
	"testing"

	"github.com/edgar0bsj/nerv-desk/module/user/database"
	"github.com/edgar0bsj/nerv-desk/module/user/model"
	"github.com/edgar0bsj/nerv-desk/module/user/repository"
	"github.com/edgar0bsj/nerv-desk/module/user/service"
)

func setup() *service.UserService {
	db, _ := database.New()
	repository := repository.New(db)
	service := service.New(repository)

	return service
}

// ==============================================
// TESTE DE PERSISTENCIA
// ==============================================
func TestUserService_Save(t *testing.T) {
	scv := setup()
	userFake := service.UserRegisterDto{
		Name:     "Luciane Barbosa de lima",
		Email:    "Lucianee@email.com",
		Password: "1234567",
		Role:     model.RoleAttendant,
	}

	err := scv.SaveUser(&userFake)

	if err != nil {
		t.Fatalf("[Error] %v", err)
	}

}

// ==============================================
// TESTE DE LISTAR USUARIOS
// ==============================================
func TestUserService_GetAll(t *testing.T) {
	scv := setup()

	users, err := scv.FindAllUsers()

	if err != nil {
		t.Fatalf("[ERROR] %v", err)
	}

	t.Log("Testa feito com sucesso!")
	for _, v := range users {
		t.Log(v)
	}

}

// ==============================================
// TESTE DE BUSCAR UM USUARIO
// ==============================================
func TestUserService_GetOne(t *testing.T) {
	var idx int
	scv := setup()

	users, _ := scv.FindAllUsers()

	for i, v := range users {
		if v.Name != "Luciane Barbosa de lima" {
			continue
		}
		idx = i
	}

	oneUser, err := scv.FindByIdUser(users[idx].ID)

	if err != nil {
		t.Fatalf("Error ao Buscar usuario")
	}

	t.Log("Teste passou com sucesso!")
	t.Logf("%v", oneUser)
}

// ==============================================
// TESTE DE ATUALIZAR USUARIO
// ==============================================
func TestUserService_UpdateUser(t *testing.T) {
	scv := setup()

	_, err := scv.FindByIdUser("9d418d07-8837-403e-942f-e1274abf557b")
	if err != nil {
		t.Fatalf("ERROR AO BUSCAR USUARIO")
	}

	userUpdateDto := service.UserUpdateDto{
		ID:    "9d418d07-8837-403e-942f-e1274abf557b",
		Name:  "Luciane de Lima",
		Email: "Luluzete@email.com",
		Role:  model.RoleUser,
	}

	err = scv.UpdateUser(&userUpdateDto)
	if err != nil {
		t.Fatalf("ERROR AO ATUALIZAR USUARIO: %v", err)
	}

}

// ==============================================
// TESTE DE DELETAR USUARIO
// ==============================================
func TestUserService_DeleteUser(t *testing.T) {
	scv := setup()

	if err := scv.DeleteUser("9d418d07-8837-403e-942f-e1274abf557b"); err != nil {
		t.Fatalf("[Error] %v", err)
	}

}
