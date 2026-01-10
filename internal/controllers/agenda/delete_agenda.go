package agenda

import (
	"fmt"
	"middleware/example/internal/helpers"
	agenda "middleware/example/internal/services/agendas"
	"net/http"

	"github.com/gofrs/uuid"
)

// DeleteAgenda
// @Tags         agendas
// @Summary      Delete an agenda.
// @Description  Delete an agenda by UUID.
// @Param        id   path      string  true  "Agenda UUID formatted ID"
// @Success      204  "No content"
// @Failure      422  "Cannot parse id"
// @Failure      500  "Something went wrong"
// @Router       /agendas/{id} [delete]
func DeleteAgenda(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	agendaId, ok := ctx.Value("id").(uuid.UUID)
	if !ok {
		body, status := helpers.RespondError(fmt.Errorf("Invalid agenda ID"))
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	err := agenda.DeleteAgenda(agendaId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
