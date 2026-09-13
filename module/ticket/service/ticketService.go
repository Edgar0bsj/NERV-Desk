package service

import (
	"time"

	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
	"github.com/edgar0bsj/nerv-desk/module/ticket/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TicketService struct {
	repo repository.TicketInterfaceRepository
}

func New(repo repository.TicketInterfaceRepository) *TicketService {
	return &TicketService{
		repo: repo,
	}
}

func (s *TicketService) SaveTicket(ticket *TicketCreateDto) error {
	validation := validator.New()

	if err := validation.Struct(ticket); err != nil {
		return err
	}

	ticketEntity := model.TicketModel{
		ID:          uuid.New().String(),
		UserID:      ticket.User_id,
		Title:       ticket.Title,
		Description: ticket.Description,
		Status:      model.StatusOpen,
		Priority:    ticket.Priority,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}

	if err := s.repo.Save(&ticketEntity); err != nil {
		return err
	}

	return nil

}

func (s *TicketService) FindAllTickets() ([]*model.TicketModel, error) {
	return s.repo.FindAll()
}

func (s *TicketService) FindByIdTicket(id string) (*model.TicketModel, error) {
	return s.repo.FindByID(id)
}

func (s *TicketService) UpdateTicket(ticketDto *TicketUpdateDto) error {
	validation := validator.New()
	if err := validation.Struct(ticketDto); err != nil {
		return err
	}

	ticket, err := s.repo.FindByID(ticketDto.ID)
	if err != nil {
		return err
	}

	ticket.AttendantID = ticketDto.Attendant_id
	ticket.Title = ticketDto.Title
	ticket.Description = ticketDto.Description
	ticket.Status = ticketDto.Status
	ticket.Priority = ticketDto.Priority

	if err := s.repo.Update(ticket); err != nil {
		return err
	}
	return nil
}

func (s *TicketService) DeleteTicket(id string) error {
	return s.repo.Delete(id)
}

func (s *TicketService) FindByTitle(title string) (*model.TicketModel, error) {
	return s.repo.FindByTitle(title)
}

func (s *TicketService) FindAllByUserId(user_id string) ([]*model.TicketModel, error) {
	return s.repo.FindAllByUserId(user_id)
}
