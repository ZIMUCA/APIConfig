package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Agenda struct {
	Id         *uuid.UUID `json:"id"`
	GroupID    string     `json:"group_id"`
	CalendarID string     `json:"calendar_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
