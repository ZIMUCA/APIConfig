package repository

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"time"
)

func GetAllEvents() ([]models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query(`SELECT 
		id, 
		dtStamp, 
		dtStart, 
		dtEnd, 
		summary, 
		location, 
		description, 
		uid, 
		created, 
		lastModified, 
		sequence 
		FROM Event`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Event{}

	for rows.Next() {
		var e models.Event
		var dtStampStr, dtStartStr, dtEndStr, createdStr, lastModifiedStr string

		err = rows.Scan(
			&e.Id,
			&dtStampStr,
			&dtStartStr,
			&dtEndStr,
			&e.Summary,
			&e.Location,
			&e.Description,
			&e.Uid,
			&createdStr,
			&lastModifiedStr,
			&e.Sequence,
		)
		if err != nil {
			return nil, err
		}

		// Parsing des dates
		e.DtStamp, _ = time.Parse("20060102T150405Z", dtStampStr)
		e.DtStart, _ = time.Parse("20060102T150405Z", dtStartStr)
		e.DtEnd, _ = time.Parse("20060102T150405Z", dtEndStr)
		e.Created, _ = time.Parse("20060102T150405Z", createdStr)
		e.LastModified, _ = time.Parse("20060102T150405Z", lastModifiedStr)

		events = append(events, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
