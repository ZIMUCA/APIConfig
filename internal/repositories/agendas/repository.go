package repository

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"time"

	"github.com/gofrs/uuid"
)

func GetAllAgendas() ([]models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query(`SELECT 
		id, 
		groupId, 
		calendarID, 
		createdAt, 
		updatedAt`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Agenda{}

	for rows.Next() {
		var e models.Agenda
		var createdAt, updatedAt string

		err = rows.Scan(
			&e.Id,
			&e.GroupID,
			&e.CalendarID,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Parsing des dates
		e.CreatedAt, _ = time.Parse("20060102T150405Z", createdAt)
		e.UpdatedAt, _ = time.Parse("20060102T150405Z", updatedAt)

		events = append(events, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetAgendaById(id uuid.UUID) (*models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM agenda WHERE id=?", id.String())
	helpers.CloseDB(db)

	var agenda models.Agenda
	err = row.Scan(&agenda.Id, &agenda.GroupID)
	if err != nil {
		return nil, err
	}
	return &agenda, err
}
