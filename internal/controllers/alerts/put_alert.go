package alerts

import (
	"encoding/json"
	"fmt"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	alert "middleware/example/internal/services/alerts"
	"net/http"

	"github.com/gofrs/uuid"
)

// UpdateAgenda
// @Tags         agendas
// @Summary      Update an genda.
// @Description  Update an agenda by UUID.
// @Param        id    path      string      true  "Agenda UUID formatted ID"
// @Param        agenda  body      models.Agenda  true  "Updated agenda data"
// @Success      200   {object}  models.Agenda
// @Failure      400   "Invalid request body"
// @Failure      422   "Cannot parse id"
// @Failure      500   "Something went wrong"
// @Router       /agendas/{id} [put]
func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, ok := ctx.Value("id").(uuid.UUID)
	if !ok {
		body, status := helpers.RespondError(fmt.Errorf("Invalid alert ID"))
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	var updatedData models.Alerts
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	updatedAlert, err := alert.UpdateAgenda(alertId, &updatedData)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(updatedAlert)
	_, _ = w.Write(body)
}
