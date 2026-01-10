package alerts

import (
	"database/sql"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/alerts"

	"github.com/sirupsen/logrus"
)

func getAllAlerts() ([]models.Alerts, error) {

	alerts, err := repository.GetAllAlerts()

	// Gestion des erreurs
	if err != nil {
		logrus.Errorf("error retrieving events: %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving events",
		}
	}

	return alerts, nil
}

func CreateAlerts(newAlert *models.Alerts) (*models.Alerts, error) {
	alert, err := repository.PostAlert(newAlert)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "Error in the creation",
			}
		}
	}

	return alert, err
}
