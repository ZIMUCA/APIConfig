package main

import (
	"middleware/example/internal/controllers/agenda"
	"middleware/example/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {

	r := chi.NewRouter()

	r.Route("/agendas", func(r chi.Router) { // route /agendas
		r.Get("/", agenda.GetAgendas)         // GET /agendas
		r.Post("/", agenda.PostAgenda)        // POST /agendas
		r.Route("/{id}", func(r chi.Router) { // route /agendas/{id}
			r.Use(agenda.Context)              // Use Context method to get agenda ID
			r.Get("/", agenda.GetAgenda)       // GET /agendas/{id}
			r.Delete("/", agenda.DeleteAgenda) // DELETE /agendas/{id}
			r.Put("/", agenda.UpdateAgenda)    //PUT /agendas/{id}
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
			group_id TEXT,
			calendar_id TEXT,
			created_at DATETIME,
			updated_at DATETIME
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
