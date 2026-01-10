package repository

import (
	"log"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllAlerts() ([]models.Alerts, error) {

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query(`SELECT id, agenda_id,mail FROM alert`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	alerts := []models.Alerts{}

	for rows.Next() {
		var e models.Alerts

		err = rows.Scan(
			&e.Id,
			&e.AgendaId,
			&e.Mail,
		)

		if err != nil {
			return nil, err
		}

		alerts = append(alerts, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}

func PostAlert(newAlert *models.Alerts) (*models.Alerts, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO alert (id, agenda_id,mail) VALUES (?, ?,?)",
		newAlert.Id.String(), newAlert.AgendaId, newAlert.Mail)
	helpers.CloseDB(db)

	if err != nil {
		log.Println("DB QUERY ERROR:", err)
		return nil, err
	}
	return newAlert, nil
}
