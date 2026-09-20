package repository

import (
	"time"

	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
	"gorm.io/gorm"
)

type TickerUserQuery struct {
	ID          string
	Title       string
	Description string
	Status      model.TicketStatus
	Priority    model.TicketPriority

	User_id    string
	User_name  string
	User_email string

	Attendant_id    string
	Attendant_name  string
	Attendant_email string

	Created_at time.Time
	Updated_at time.Time
}

type TicketRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (t *TicketRepository) Save(ticket *model.TicketModel) error {
	err := t.db.Create(&ticket).Error

	if err != nil {
		return err
	}
	return nil
}

func (t *TicketRepository) FindAll() ([]*TickerUserQuery, error) {
	var tickets []*TickerUserQuery

	err := t.db.
		Table("ticket_models AS ticket").
		Select(`
		ticket.id,
		ticket.title,
		ticket.description,
		ticket.status,
		ticket.priority,
		ticket.created_at,
		ticket.updated_at,

		usuario.id AS user_id,
		usuario.name AS user_name,
		usuario.email AS user_email,

		atendente.id AS attendant_id,
		atendente.name AS attendant_name,
		atendente.email AS attendant_email
	`).
		Joins("JOIN users AS usuario ON usuario.id = ticket.user_id").
		Joins("LEFT JOIN users AS atendente ON atendente.id = ticket.attendant_id").
		Scan(&tickets).Error

	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (t *TicketRepository) FindByID(id string) (*model.TicketModel, error) {
	var ticket *model.TicketModel

	err := t.db.Where("id = ?", id).First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return ticket, nil

}

func (t *TicketRepository) Update(ticket *model.TicketModel) error {
	var ticketEntity *model.TicketModel

	err := t.db.Where("id = ?", ticket.ID).First(&ticketEntity).Error

	if err != nil {
		return err
	}

	t.db.Model(&ticketEntity).Updates(ticket)

	return nil
}

func (t *TicketRepository) Delete(id string) error {
	var ticket *model.TicketModel

	err := t.db.Where("id = ?", id).First(&ticket).Error
	if err != nil {
		return err
	}

	t.db.Delete(&ticket)

	return nil
}

func (t *TicketRepository) FindByTitle(title string) (*model.TicketModel, error) {
	var ticket *model.TicketModel

	err := t.db.Where("title LIKE ?", "%"+title+"%").First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (t *TicketRepository) ListMyTickets(userId string) ([]*TickerUserQuery, error) {
	var tickets []*TickerUserQuery

	err := t.db.
		Table("ticket_models AS ticket").
		Select(`
		ticket.id,
		ticket.title,
		ticket.description,
		ticket.status,
		ticket.priority,
		ticket.created_at,
		ticket.updated_at,

		usuario.id AS user_id,
		usuario.name AS user_name,
		usuario.email AS user_email,

		atendente.id AS attendant_id,
		atendente.name AS attendant_name,
		atendente.email AS attendant_email
	`).
		Joins("JOIN users AS usuario ON usuario.id = ticket.user_id").
		Joins("LEFT JOIN users AS atendente ON atendente.id = ticket.attendant_id").
		Where("ticket.user_id = ?", userId).
		Scan(&tickets).Error

	if err != nil {
		return nil, err
	}
	return tickets, nil
}
