package agenda

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	agenda "middleware/example/internal/services/agendas"
	"net/http"

	"github.com/gofrs/uuid"
)

// GetAgenda
// @Tags         agendas
// @Summary      Get a agenda.
// @Description  Get a agenda.
// @Param        id           	path      string  true  "Agenda UUID formatted ID"
// @Success      200            {object}  models.Agenda
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /agendas/{id} [get]
func GetAgenda(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	agendaId, _ := ctx.Value("id").(uuid.UUID) // getting key set in context.go

	agenda, err := agenda.GetAgendaById(agendaId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(agenda)
	_, _ = w.Write(body)
	return
}
