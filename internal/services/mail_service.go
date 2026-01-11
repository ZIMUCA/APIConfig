package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/sirupsen/logrus"
)

var MAIL_API_TOKEN = "AABuyqRBSemauCjvJZBAEvgGvbwbVVwoSGLhiSxq"

type MailRequest struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}

func SendAlertMail(mail string, alert models.Alert) error {
	html, matter, err := helpers.GetStringFromEmbeddedTemplate(
		"config/event_changed.html",
		alert,
	)
	if err != nil {
		return err
	}

	reqBody := MailRequest{
		Recipient: mail,
		Subject:   matter.Subject,
		Content:   html,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"https://mail-api.edu.forestier.re/mail",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", MAIL_API_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("mail API returned %d", resp.StatusCode)
	}

	logrus.Infof("Mail envoyé avec succès à %s", mail)
	return nil
}
