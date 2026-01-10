package models

import (
	"github.com/gofrs/uuid"
)

type Agenda struct {
	Id       *uuid.UUID `json:"id"`
	AgendaId string     `json:"agenda_id"`
	Name     string     `json:"name"`
}
