package agenda

import (
	"database/sql"
	"fmt"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/agendas"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllAgendas() ([]models.Agenda, error) {
	agendas, err := repository.GetAllAgendas()

	// Gestion des erreurs
	if err != nil {
		logrus.Errorf("error retrieving events: %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving events",
		}
	}

	return agendas, nil
}

func GetAgendaById(id uuid.UUID) (*models.Agenda, error) {
	user, err := repository.GetAgendaById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "agenda not found",
			}
		}
		logrus.Errorf("error retrieving user %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving user %s", id.String()),
		}
	}

	return user, err
}
