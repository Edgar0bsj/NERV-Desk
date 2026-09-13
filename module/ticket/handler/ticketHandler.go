package handler

import (
	"net/http"

	"github.com/edgar0bsj/nerv-desk/module/ticket/service"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service service.TicketInterfaceService
}

func New(svc service.TicketInterfaceService) *TicketHandler {
	return &TicketHandler{
		service: svc,
	}
}

func (h *TicketHandler) FindAllTicketsUser(c *gin.Context) {
	userID := c.GetString("userID")

	userTickets, err := h.service.FindAllByUserId(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error ao buscar lista de Tickets"})
		return
	}

	c.JSON(http.StatusOK, userTickets)
}
