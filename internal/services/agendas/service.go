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
	agenda, err := repository.GetAgendaById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "agenda not found",
			}
		}
		logrus.Errorf("error retrieving agenda %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving agenda %s", id.String()),
		}
	}

	return agenda, err
}

func DeleteAgenda(id uuid.UUID) error {
	err := repository.DeleteAgendaById(id)
	return err
}

func CreateAgenda(newAgenda *models.Agenda) (*models.Agenda, error) {

	agenda, err := repository.PostAgenda(newAgenda)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "creation impossible",
			}
		}
	}

	return agenda, err
}

func UpdateAgenda(id uuid.UUID, updatedAgenda *models.Agenda) (*models.Agenda, error) {
	agenda, err := repository.PutAgendaById(id, updatedAgenda)
	return agenda, err
}
