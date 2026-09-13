package repository_test

import (
	"testing"
	"time"

	"github.com/edgar0bsj/nerv-desk/module/ticket/database"
	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
	"github.com/edgar0bsj/nerv-desk/module/ticket/repository"
	"github.com/google/uuid"
)

func setup() (*repository.TicketRepository, *model.TicketModel) {
	// Criado um fake ticket
	ticketFake := &model.TicketModel{
		ID:          uuid.New().String(),
		UserID:      "536a819d-0e99-4cdd-b124-ac8abbb7e6e4",
		Title:       "Sem acesso ao sistema",
		Description: "Não consigo entrar no portal de jeito nenhum",
		Status:      model.StatusOpen,
		Priority:    model.PriorityHigh,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}

	// Iniciando database
	db, _ := database.New()

	repo := repository.New(db)
	return repo, ticketFake
}

func TestTicket_Create(t *testing.T) {

	storage, tickeFake := setup()

	err := storage.Save(tickeFake)

	if err != nil {
		t.Fatalf("Error ao SALVAR ticket -> %v", err)
	}

}

func TestTicket_FindByTitle(t *testing.T) {

	storage, tickeFake := setup()

	ticket, err := storage.FindByTitle(tickeFake.Title)

	if err != nil {
		t.Fatalf("Error ao Buscar Ticket -> %v", err)
	}

	t.Logf("Ticket encontrado com sucesso! -> %v", ticket)
}

func TestTicket_FindById(t *testing.T) {

	storage, tickeFake := setup()

	ticket, err := storage.FindByTitle(tickeFake.Title)

	if err != nil {
		t.Fatalf("Error ao Buscar Ticket -> %v", err)
	}

	result, err := storage.FindByID(ticket.ID)

	if err != nil {
		t.Fatalf("Error ao buscar o ID do ticket -> %v", err)
	}

	t.Logf("Ticket encontrado com sucesso! -> %v", result)
}

func TestTicket_FindAll(t *testing.T) {

	storage, _ := setup()

	result, err := storage.FindAll()
	if err != nil {
		t.Fatalf("Erro ao retornar todos os Tickets -> %v", err)
	}

	t.Logf("Total de Ticke retornado! -> %v", len(result))
}

func TestTicket_Update(t *testing.T) {

	storage, ticketFake := setup()

	ticket, err := storage.FindByTitle(ticketFake.Title)
	if err != nil {
		t.Fatalf("Erro ao retornar todos os Tickets -> %v", err)
	}

	ticket.Title = "TESTANDO AQUI"
	ticket.Status = model.StatusInProcess

	if err := storage.Update(ticket); err != nil {
		t.Fatalf("Error ao Atualizar Ticket")
	}

	t.Logf("Total de Ticke retornado! -> %v", ticket)
}

func TestTicket_Delete(t *testing.T) {

	storage, ticketFake := setup()

	_, err := storage.FindByTitle(ticketFake.Title)
	if err != nil {
		t.Fatalf("Erro ao retornar todos os Tickets -> %v", err)
	}

	if err := storage.Delete("d12a8b6d-32b7-479f-9b78-4cfd80b82f38"); err != nil {
		t.Fatalf("Error ao Atualizar Ticket")
	}

	t.Log("Ticket Deletado com sucesso!")
}
