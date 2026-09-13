package service_test

import (
	"testing"

	"github.com/edgar0bsj/nerv-desk/module/ticket/database"
	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
	"github.com/edgar0bsj/nerv-desk/module/ticket/repository"
	"github.com/edgar0bsj/nerv-desk/module/ticket/service"
)

func setup() *service.TicketService {
	db, _ := database.New()
	storage := repository.New(db)
	svc := service.New(storage)

	return svc
}

func TestTicketService_Create(t *testing.T) {
	scv := setup()

	tickeCreate := service.TicketCreateDto{
		User_id:     "536a819d-0e99-4cdd-b124-ac8abbb7e6e4",
		Title:       "Sem acesso a rede NOVAMENTE",
		Description: "Não consigo entrar no portal",
		Priority:    model.PriorityHigh,
	}

	err := scv.SaveTicket(&tickeCreate)
	if err != nil {
		t.Fatalf("Error ao Criar Ticket -> %v", err)
	}
}

func TestTicketService_FindByTitle(t *testing.T) {
	scv := setup()

	ticket, err := scv.FindByTitle("Sem acesso a rede")
	if err != nil {
		t.Fatalf("Error ao buscar por titulo -> %v", ticket)
	}

	t.Log(ticket)
}

func TestTicketService_FindById(t *testing.T) {
	scv := setup()

	ticket, err := scv.FindByTitle("Sem acesso a rede")
	if err != nil {
		t.Fatalf("Error ao buscar por titulo -> %v", err)
	}

	ticketPorId, err := scv.FindByIdTicket(ticket.ID)
	if err != nil {
		t.Fatalf("Error ao buscar por Id -> %v", err)
	}

	t.Log(ticketPorId)
}

func TestTicketService_FindAll(t *testing.T) {
	scv := setup()
	allTickets, err := scv.FindAllTickets()

	if err != nil {
		t.Fatalf("Error ao buscar todos os tickets -> %v", err)
	}

	t.Log(allTickets)
}
func TestTicketService_Update(t *testing.T) {
	scv := setup()
	ticket, err := scv.FindByTitle("Sem acesso a rede")

	if err != nil {
		t.Fatalf("Error ao Buscar por titulo -> %v", err)
	}

	ticketUpdate := service.TicketUpdateDto{
		ID:          ticket.ID,
		Title:       "Titulo alterado",
		Description: ticket.Description,
		Status:      model.StatusInProcess,
		Priority:    model.PriorityCritical,
	}

	if err := scv.UpdateTicket(&ticketUpdate); err != nil {
		t.Fatalf("Error ao atualizar Ticket -> %v", err)
	}

}
func TestTicketService_Delete(t *testing.T) {
	scv := setup()
	ticket, err := scv.FindByTitle("Sem acesso a rede")

	if err != nil {
		t.Fatalf("Error ao Buscar por titulo -> %v", err)
	}

	if err := scv.DeleteTicket(ticket.ID); err != nil {
		t.Fatalf("Error ao Deletar Ticket -> %v", err)
	}

}
