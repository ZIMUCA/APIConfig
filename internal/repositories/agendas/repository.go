package repository

import (
	"log"
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
		Id, 
		group_id, 
		calendar_id, 
		created_at, 
		updated_at FROM AGENDA`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Agenda{}

	for rows.Next() {
		var e models.Agenda
		var createdAt, updatedAt, idStr string

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

		parsedID, err := uuid.FromString(idStr)
		if err != nil {
			return nil, err
		}
		e.Id = &parsedID

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

func PostAgenda(newAgenda *models.Agenda) (*models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO agenda (id, group_id,calendar_id,created_at,updated_at) VALUES (?, ?,?,?,?)",
		newAgenda.Id.String(), newAgenda.GroupID, newAgenda.CalendarID, newAgenda.CreatedAt, newAgenda.UpdatedAt)
	helpers.CloseDB(db)

	if err != nil {
		log.Println("DB QUERY ERROR:", err)
		return nil, err
	}
	return nil, nil
}

func PutAgendaById(updatedAgenda *models.Agenda) (*models.Agenda, error) {
	return nil, nil
	/*db, err := helpers.OpenDB()
		if err != nil {
			return nil, err
		}
		_, err = db.Exec("UPDATE agenda SET groupId = ?, calendarId = ?, createdAt = ?, updatedAt WHERE id = ?",
	    updatedAgenda.GroupID.String(),
	    updatedAgenda.Title,
	    updatedAgenda.Date,
	    updatedAgenda.Id.String(),
	)
		helpers.CloseDB(db)

		return nil, err*/
}

func DeleteAgendaById(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM agenda WHERE id = ?", id.String())
	helpers.CloseDB(db)

	return err
}
