package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/middleware"
)

type Handler struct {
	moderatorService ModeratorService
}

func NewModeratorHandler(moderatorService ModeratorService) *Handler {
	return &Handler{
		moderatorService: moderatorService,
	}
}

// TODO: проверка ролей
func MapModeratorRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Post("/addHomework/class", mw.Auth(), h.AddHomeworkToClass())
	g.Post("/addHomework/date", mw.Auth(), h.AddHomeworkToDate())
}
