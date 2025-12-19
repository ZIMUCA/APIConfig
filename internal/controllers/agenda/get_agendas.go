package agenda

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	agenda "middleware/example/internal/services/agendas"
	"net/http"
)

// GetAgendas
// @Tags         users
// @Summary      Get all users.
// @Description  Get all users.
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /users [get]
func GetAgendas(w http.ResponseWriter, _ *http.Request) {
	// calling service
	agendas, err := agenda.GetAllAgendas()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(agendas)
	_, _ = w.Write(body)
	return
}
