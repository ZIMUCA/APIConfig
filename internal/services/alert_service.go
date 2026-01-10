package services

type Alert struct {
	Email string
}

func GetAlertsForEvent(eventUID string) ([]Alert, error) {
	// TP : logique simplifiée
	// En réel : DB + lien agenda/event

	return []Alert{
		{Email: "Maxime.VIMPERE@etu.uca.fr"},
	}, nil
}
