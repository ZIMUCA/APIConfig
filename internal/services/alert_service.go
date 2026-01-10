package services

import (
	"database/sql"
	"errors"
	"strings"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetMailsForAlert(alert models.Alert) ([]string, error) {

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// 1. Récupération de tous les agendas
	rows, err := db.Query(`
		SELECT agenda_id, name
		FROM agenda
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type agendaRow struct {
		AgendaId string
		Name     string
	}

	var matchedAgendaIDs []string

	for rows.Next() {
		var res agendaRow
		if err := rows.Scan(&res.AgendaId, &res.Name); err != nil {
			return nil, err
		}

		if strings.Contains(strings.ToLower(alert.Summary), strings.ToLower(res.Name)) ||
			strings.Contains(strings.ToLower(alert.Description), strings.ToLower(res.Name)) {

			matchedAgendaIDs = append(matchedAgendaIDs, res.AgendaId)
		}
	}

	if len(matchedAgendaIDs) == 0 {
		return []string{}, nil
	}

	// 2. Récupération des mails associés
	mails := []string{}

	query := `
		SELECT mail
		FROM alerts
		WHERE agenda_id = ?
	`

	for _, agendaID := range matchedAgendaIDs {
		rows, err := db.Query(query, agendaID)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var mail string
			if err := rows.Scan(&mail); err != nil {
				rows.Close()
				return nil, err
			}
			mails = append(mails, mail)
		}
		rows.Close()
	}

	return mails, nil
}

func getAgendaIdByName(db *sql.DB, name string) (string, error) {
	var agendaId string

	err := db.QueryRow(`
		SELECT agenda_id
		FROM agenda
		WHERE name = ?
	`, name).Scan(&agendaId)

	if err == sql.ErrNoRows {
		return "", errors.New("agenda not found for summary: " + name)
	}

	return agendaId, err
}

func getMailsByAgendaId(db *sql.DB, agendaId string) ([]string, error) {
	rows, err := db.Query(`
		SELECT mail
		FROM alerts
		WHERE agenda_id = ?
	`, agendaId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mails []string
	for rows.Next() {
		var mail string
		if err := rows.Scan(&mail); err == nil {
			mails = append(mails, mail)
		}
	}

	if len(mails) == 0 {
		return nil, errors.New("no alerts configured for agenda")
	}

	return mails, nil
}
