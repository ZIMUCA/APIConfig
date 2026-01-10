package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

var MAIL_API_TOKEN = "AABuyqRBSemauCjvJZBAEvgGvbwbVVwoSGLhiSxq"

type MailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	HTML    string `json:"html"`
}

func SendAlertMail(mail string, alert models.Alert) error {

	// Génération du contenu HTML + subject depuis le template
	html, matter, err := helpers.GetStringFromEmbeddedTemplate(
		"templates/event_changed.html",
		alert,
	)
	if err != nil {
		return err
	}

	reqBody := MailRequest{
		To:      mail,
		Subject: matter.Subject,
		HTML:    html,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"https://mail.edu.forestier.re/api/send",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv(MAIL_API_TOKEN))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Vérification du statut HTTP
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("mail API returned status %d", resp.StatusCode)
	}

	return nil
}
