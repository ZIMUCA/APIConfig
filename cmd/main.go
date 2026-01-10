package main

import (
	consumers "middleware/example/internal/consumers"
	"middleware/example/internal/controllers"
	"middleware/example/internal/controllers/agenda"
	"middleware/example/internal/controllers/alerts"
	"net/http"

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
			logrus.Errorf("error creating alerter consumer: %v", err)
			return
		}

		if err := consumers.ConsumeAlerter(*consumer); err != nil {
			logrus.Errorf("error consuming alerts: %v", err)
		}
	}()

	// API REST (Config)
	r := chi.NewRouter()

	r.Route("/agendas", func(r chi.Router) { // route /agendas
		r.Get("/", agenda.GetAgendas)         // GET /agendas
		r.Post("/", agenda.PostAgenda)        // POST /agendas
		r.Route("/{id}", func(r chi.Router) { // route /agendas/{id}
			r.Use(controllers.Context)         // Use Context method to get agenda ID
			r.Get("/", agenda.GetAgenda)       // GET /agendas/{id}
			r.Delete("/", agenda.DeleteAgenda) // DELETE /agendas/{id}
		})
	})

	r.Route("/alerts", func(r chi.Router) { // route /alerts
		r.Get("/", alerts.GetAlerts)          // GET /alerts
		r.Post("/", alerts.PostAlert)         // POST /alerts
		r.Route("/{id}", func(r chi.Router) { // route /alerts/{id}
			r.Use(controllers.Context)     // Use Context method to get alert ID
			r.Put("/", alerts.UpdateAlert) //PUT /alerts/{id}
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8081")
	logrus.Fatalln(http.ListenAndServe(":8081", r))

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
		`CREATE TABLE IF NOT EXISTS alerts (
			id TEXT PRIMARY KEY NOT NULL UNIQUE,
			agenda_id INTEGER,
			mail TEXT
		);`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalf("Could not generate table! Error was: %s", err.Error())
		}
	}

	logrus.Info("Database schema initialized successfully")
	helpers.CloseDB(db)
}
