package services

import (
	"bytes"
	"encoding/json"
	"middleware/example/internal/helpers"
	"net/http"
	"os"
)

type MailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	HTML    string `json:"html"`
}

func SendAlertMail(alert Alert, event interface{}) error {

	html, matter, err := helpers.GetStringFromEmbeddedTemplate(
		"templates/event_changed.html",
		event,
	)
	if err != nil {
		return err
	}

	reqBody := MailRequest{
		To:      alert.Email,
		Subject: matter.Subject,
		HTML:    html,
	}

	data, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(
		"POST",
		"https://mail.edu.forestier.re/api/send",
		bytes.NewBuffer(data),
	)

	req.Header.Set("Authorization", "Bearer "+os.Getenv("MAIL_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
