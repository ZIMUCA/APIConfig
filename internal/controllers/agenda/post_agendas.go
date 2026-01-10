package agenda

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	agenda "middleware/example/internal/services/agendas"
	"net/http"
)

// CreateAgenda
// @Tags         agendas
// @Summary      Create a agenda.
// @Description  Create a new agenda.
// @Accept       json
// @Produce      json
// @Param        agenda      body      agendas.Agenda  true  "Agenda data"
// @Success      201       {object}  agendas.Agenda
// @Failure      400       "Invalid request body"
// @Failure      500       "Something went wrong"
// @Router       /agendas [post]
func PostAgenda(w http.ResponseWriter, r *http.Request) {

	var newAgenda models.Agenda
	if err := json.NewDecoder(r.Body).Decode(&newAgenda); err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	createdAgenda, err := agenda.CreateAgenda(&newAgenda)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(createdAgenda)
	_, _ = w.Write(body)
}
