package models

import (
	"github.com/gofrs/uuid"
)

type Alert struct {
	Id                *uuid.UUID `json:"id"`
	Destinataire      string     `json:"destinataire"`
	AgendaAssocie     Agenda     `json:"agenda_associe"`
	QuandEnvoyeAlerte string     `json:"quand_envoye_alerte"`
}
