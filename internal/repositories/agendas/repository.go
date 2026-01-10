package repository

import (
	"log"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllAgendas() ([]models.Agenda, error) {

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query(`SELECT id, agenda_id,name FROM agenda`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Agenda{}

	for rows.Next() {
		var e models.Agenda

		err = rows.Scan(
			&e.Id,
			&e.AgendaId,
			&e.Name,
		)

		if err != nil {
			return nil, err
		}

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

	row := db.QueryRow("SELECT id, agenda_id, name FROM agenda WHERE id=?", id.String())
	helpers.CloseDB(db)

	var agenda models.Agenda
	err = row.Scan(&agenda.Id, &agenda.AgendaId, &agenda.Name)
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

	_, err = db.Exec("INSERT INTO agenda (id, agenda_id,name) VALUES (?, ?,?)",
		newAgenda.Id.String(), newAgenda.AgendaId, newAgenda.Name)
	helpers.CloseDB(db)

	if err != nil {
		log.Println("DB QUERY ERROR:", err)
		return nil, err
	}
	return newAgenda, nil
}

func PutAgendaById(id uuid.UUID, updatedAgenda *models.Agenda) (*models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("UPDATE agenda SET agenda_id = ?, name = ? WHERE id = ?",
		updatedAgenda.AgendaId,
		updatedAgenda.Name,
		updatedAgenda.Id.String(),
	)
	helpers.CloseDB(db)

	return updatedAgenda, err
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
