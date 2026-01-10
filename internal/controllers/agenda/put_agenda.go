package agenda

import (
	"encoding/json"
	"fmt"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	agenda "middleware/example/internal/services/agendas"
	"net/http"

	"github.com/gofrs/uuid"
)

// UpdateUser
// @Tags         agendas
// @Summary      Update an genda.
// @Description  Update an agenda by UUID.
// @Param        id    path      string      true  "Agenda UUID formatted ID"
// @Param        agenda  body      agendas.Agenda  true  "Updated agenda data"
// @Success      200   {object}  agendas.Agenda
// @Failure      400   "Invalid request body"
// @Failure      422   "Cannot parse id"
// @Failure      500   "Something went wrong"
// @Router       /agendas/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value("Id").(uuid.UUID)
	if !ok {
		body, status := helpers.RespondError(fmt.Errorf("Invalid user ID"))
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	var updatedData models.Agenda
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	updatedAgenda, err := agenda.UpdateAgenda(userId, &updatedData)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	// Réponse HTTP 200 OK avec l'utilisateur mis à jour
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(updatedAgenda)
	_, _ = w.Write(body)
}
