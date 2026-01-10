package main

import (
	"net/http"

	consumers "middleware/example/internal/consumers"
	"middleware/example/internal/controllers/agenda"
	"middleware/example/internal/helpers"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {

	// Connexion NATS
	if err := helpers.ConnectNats(); err != nil {
		logrus.Fatalf("error while connecting to NATS: %v", err)
	}
	defer helpers.CloseNats()

	// Lancement du consumer Alerter
	go func() {
		consumer, err := consumers.AlerterConsumer()
		if err != nil {
			logrus.Fatalf("error creating alerter consumer: %v", err)
		}

		if err := consumers.ConsumeAlerter(*consumer); err != nil {
			logrus.Fatalf("error consuming alerts: %v", err)
		}
	}()

	// API REST (Config)
	r := chi.NewRouter()

	r.Route("/agendas", func(r chi.Router) {
		r.Get("/", agenda.GetAgendas)
		r.Post("/", agenda.PostAgenda)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(agenda.Context)
			r.Get("/", agenda.GetAgenda)
			r.Delete("/", agenda.DeleteAgenda)
			r.Put("/", agenda.UpdateAgenda)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}

	schemes := []string{
		`CREATE TABLE IF NOT EXISTS agenda (
			id TEXT PRIMARY KEY NOT NULL UNIQUE,
			agenda_id INTEGER,
			name TEXT
		);`,
		
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table! Error was: " + err.Error())
		}
	}

	logrus.Info("Database schema initialized successfully")
	helpers.CloseDB(db)
}
