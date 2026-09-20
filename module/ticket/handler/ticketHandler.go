package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
	"github.com/edgar0bsj/nerv-desk/module/ticket/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TicketHandler struct {
	service *service.TicketService
}

func New(svc *service.TicketService) *TicketHandler {
	return &TicketHandler{
		service: svc,
	}
}

func (h *TicketHandler) FindAllTicketsUser(c *gin.Context) {
	// 1. Obter o ID do usuário autenticado a partir do contexto da requisição.
	userID, exists := c.Get("user_id")

	// 2. Caso o ID do usuário não esteja disponível, retornar Unauthorized.
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Access denied",
		})
		return
	}

	// 3. Fazer um type assertion para garantir a typagem do ID.
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server problem",
		})
		return
	}

	// 4. Buscar os tickets pertencentes ao usuário.
	tickets, err := h.service.ListMyTickets(userIDStr)

	// 5. Caso ocorra um erro durante a busca, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server problem",
		})
		return
	}

	// 6. Retornar a lista de tickets com status de sucesso.
	c.JSON(http.StatusOK, gin.H{"tickets": tickets})

}

func (h *TicketHandler) CreateTicket(c *gin.Context) {
	// 1. Obter o ID do usuário autenticado a partir do contexto da requisição.
	userID, exists := c.Get("user_id")

	// 2. Caso o ID do usuário não esteja disponível, retornar Unauthorized.
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Access denied",
		})
		return
	}

	// 3. Fazer um type assertion para garantir a typagem do ID.
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server problem",
		})
		return
	}

	// 4. Receber e validar os dados do novo ticket.
	validation := validator.New()
	var req *service.TicketCreateDto

	// 5. Caso os dados sejam inválidos, retornar Bad Request.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error in the request body",
		})
		return
	}
	if err := validation.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "validation error",
		})
		return
	}

	// 6. Criar o ticket associado ao usuário autenticado.
	newTicket, err := h.service.CreateTicket(userIDStr, req)

	// 7. Caso ocorra um erro ao criar o ticket, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating a new ticket.",
		})
		return
	}

	// 8. Retornar o ticket criado com status de sucesso.
	c.JSON(http.StatusCreated, gin.H{"ticket": newTicket})

}

func (h *TicketHandler) DeleteTicket(c *gin.Context) {
	// 1. Obter o ID do ticket a partir da requisição.
	ticketId := c.Param("id")

	// 2. Validar o ID do ticket.
	_, err := uuid.Parse(ticketId)

	// 3. Caso o ID seja inválido, retornar Bad Request.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ticket invalid"})
		return
	}

	// 4. Buscar o ticket pelo ID.
	ticket, err := h.service.FindByIdTicket(ticketId)

	// 5. Caso o ticket não exista, retornar Not Found.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
		return
	}

	// 6. Caso o status seja diferente de "OPEN", impedir a exclusão e retornar erro.
	if ticket.Status != model.StatusOpen {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "It is not possible to delete a ticket that is in progress",
			})
		return
	}

	// 7. Excluir o ticket.
	err = h.service.DeleteTicket(ticket.ID)

	// 8. Caso ocorra um erro durante a exclusão, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "It is not possible to delete a ticket that is in progress"})
		return
	}

	// 9. Retornar resposta de sucesso confirmando a exclusão.
	c.JSON(http.StatusOK, gin.H{"msg": "ticket excluido com sucesso!"})

}

func (h *TicketHandler) FindAllTicketsAttendant(c *gin.Context) {
	// 1. Obter os chamados dos usuários.
	allTickets, err := h.service.FindAllTickets()

	// // 2. Caso ocorra um erro durante a busca, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// 3. Retornar a lista de chamados com status de sucesso.
	c.JSON(http.StatusOK, gin.H{"tickets": allTickets})

}

func (h *TicketHandler) AssumeTicket(c *gin.Context) {
	// 1. Obter o ID do ticket a partir dos parâmetros da rota.
	ticketId := c.Param("id")

	// 2. Obter o ID do atendente autenticado a partir do contexto da requisição.
	_attendantId, exists := c.Get("user_id")

	// 3. Validar os IDs recebidos.
	if _, err := uuid.Parse(ticketId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter"})
		return
	}
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// 5. Buscar o ticket pelo ID.
	ticket, err := h.service.FindByIdTicket(ticketId)

	// 6. Caso o ticket não exista, retornar Not Found.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	// 7. Verificar se o ticket possui um atendente vinculado.
	if ticket.AttendantID != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ticket already linked"})
		return
	}

	// 9. Verificar se o status do ticket é "OPEN".
	if ticket.Status != model.StatusOpen {
		c.JSON(http.StatusBadRequest, gin.H{"error": "is not open"})
		return
	}

	// 11. Vincular o atendente autenticado ao ticket.
	_attend, ok := _attendantId.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	resultTicket, err := h.service.SetteAttendant(ticket, _attend)

	// 12. Caso ocorra um erro durante a atualização, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// 13. Retornar o ticket atualizado com status de sucesso.
	c.JSON(http.StatusOK, gin.H{"tickets": resultTicket})

}

func (h *TicketHandler) ChangeStatus(c *gin.Context) {
	// 1. Obter o ID do ticket a partir dos parâmetros da rota.
	ticketId := c.Param("id")

	// 2. Obter o ID do atendente autenticado a partir do contexto da requisição.
	_attendantId, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// 3. Receber e validar o novo status informado.
	validetion := validator.New()
	var req *service.TicketChangeStatusDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// 4. Verificar se o novo status é "RESOLVED" ou "CLOSED".
	if err := validetion.Struct(req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Validation error, only \"RESOLVED\" or \"CLOSED\" statuses are accepted.",
			})
		return
	}

	// 6. Buscar o ticket pelo ID.
	ticket, err := h.service.FindByIdTicket(ticketId)

	// 7. Caso o ticket não exista, retornar Not Found.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
		return
	}

	// 8. Verificar se o ticket está vinculado ao atendente que realizou a requisição.
	_attendantIdStr, ok := _attendantId.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// 9. Caso o ticket pertença a outro atendente ou não possua atendente vinculado, impedir a alteração e retornar Forbidden.
	if ticket.AttendantID != _attendantIdStr {
		c.JSON(http.StatusForbidden, gin.H{"error": "ticket linked to another attendant"})
		return
	}

	// 10. Verificar se o status atual do ticket é "EM_PROCESS".
	if (ticket.Status != model.StatusInProcess) && (ticket.Status != model.StatusOpen) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The ticket is open or in process"})
		return
	}

	// 12. Atualizar o status do ticket para o novo status.
	err = h.service.ChangeStatus(ticket, req)

	// 13. Caso ocorra um erro durante a atualização, retornar Internal Server Error.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	// 14. Retornar o ticket atualizado com status de sucesso.
	response := fmt.Sprintf("ticket %v | status atualizado com sucesso %v", ticket.ID, req.Status)
	c.JSON(http.StatusOK, gin.H{"msg": response})

}
