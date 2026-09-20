package service

import (
	"time"

	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
	"github.com/edgar0bsj/nerv-desk/module/ticket/repository"
	"github.com/google/uuid"
)

type TicketService struct {
	repo *repository.TicketRepository
}

func New(repo *repository.TicketRepository) *TicketService {
	return &TicketService{
		repo: repo,
	}
}

func (s *TicketService) ListMyTickets(userId string) ([]*model.TicketModel, error) {
	return s.repo.ListMyTickets(userId)
}

func (s *TicketService) CreateTicket(userID string, userCreateDto *TicketCreateDto) (*model.TicketModel, error) {
	newTicket := model.TicketModel{
		ID:          uuid.New().String(),
		UserID:      userID,
		Title:       userCreateDto.Title,
		Description: userCreateDto.Description,
		Status:      model.StatusOpen,
		Priority:    userCreateDto.Priority,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}

	if err := s.repo.Save(&newTicket); err != nil {
		return nil, model.ErrSaveFailedTicket
	}

	return &newTicket, nil
}

func (s *TicketService) FindByIdTicket(ticketId string) (*model.TicketModel, error) {

	return s.repo.FindByID(ticketId)

}

func (s *TicketService) DeleteTicket(ticketId string) error {
	return s.repo.Delete(ticketId)

}

func (s *TicketService) FindAllTickets() ([]*model.TicketModel, error) {
	return s.repo.FindAll()

}

func (s *TicketService) SetteAttendant(ticket *model.TicketModel, attend string) (*model.TicketModel, error) {
	ticket.AttendantID = attend
	ticket.Status = model.StatusInProcess
	ticket.Updated_at = time.Now()

	if err := s.repo.Update(ticket); err != nil {
		return nil, err
	}

	newTicket, err := s.repo.FindByID(ticket.ID)

	if err != nil {
		return nil, err
	}

	return newTicket, nil

}

func (s *TicketService) ChangeStatus(ticket *model.TicketModel, req *TicketChangeStatusDto) error {
	ticket.Status = req.Status
	ticket.Updated_at = time.Now()
	if err := s.repo.Update(ticket); err != nil {
		return err
	}
	return nil
}
