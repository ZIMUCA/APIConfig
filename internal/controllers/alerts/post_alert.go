package alerts

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	alert "middleware/example/internal/services/alerts"
	"net/http"
)

// CreateAlerts
// @Tags         alerts
// @Summary      Create an alert.
// @Description  Create a new alert.
// @Accept       json
// @Produce      json
// @Param        alert      body      models.Alerts  true  "Alert data"
// @Success      201       {object}  models.Alerts
// @Failure      400       "Invalid request body"
// @Failure      500       "Something went wrong"
// @Router       /alert [post]
func PostAlert(w http.ResponseWriter, r *http.Request) {

	var newAlert models.Alerts
	if err := json.NewDecoder(r.Body).Decode(&newAlert); err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	createdAlert, err := alert.CreateAlerts(&newAlert)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(createdAlert)
	_, _ = w.Write(body)
}
