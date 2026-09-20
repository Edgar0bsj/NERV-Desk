package router

import (
	middle "github.com/edgar0bsj/nerv-desk/middleware"
	"github.com/edgar0bsj/nerv-desk/module/ticket/handler"
	userhandler "github.com/edgar0bsj/nerv-desk/module/user/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *userhandler.UserHandler, ticketHandler *handler.TicketHandler) *gin.Engine {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", userHandler.UserLogin)
	}

	userRoutes := r.Group("/user", middle.RequireUserRole())
	{
		userRoutes.GET("/ticket", ticketHandler.FindAllTicketsUser)
		userRoutes.POST("/ticket", ticketHandler.CreateTicket)
		userRoutes.DELETE("/ticket/:id", ticketHandler.DeleteTicket)
	}

	attendantRoutes := r.Group("/attendant", middle.RequireAttendantRole())
	{
		attendantRoutes.GET("/ticket", ticketHandler.FindAllTicketsAttendant)
		attendantRoutes.PATCH("ticket/assume/:id", ticketHandler.AssumeTicket)
		attendantRoutes.PATCH("ticket/changestatus/:id", ticketHandler.ChangeStatus)
	}

	adminRoutes := r.Group("/admin", middle.RequireAdminRole())
	{
		adminRoutes.POST("/register", userHandler.UserRegister)
		adminRoutes.GET("/user/list", userHandler.ListUser)
		adminRoutes.PUT("/user/edit/:id", userHandler.EditUser)
		adminRoutes.DELETE("/user/delete/:id", userHandler.DeleteUser)
		adminRoutes.PATCH("/user/change/password/:id", userHandler.ChangePassword)
		adminRoutes.GET("/ticket", ticketHandler.FindAllTicketsAttendant)
	}

	return r
}
