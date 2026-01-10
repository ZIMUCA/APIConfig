package consumers

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
)

func ConsumeAlerter(consumer jetstream.Consumer) error {

	cc, err := consumer.Consume(func(msg jetstream.Msg) {

		// Désérialisation du message
		var alert models.Alert
		if err := json.Unmarshal(msg.Data(), &alert); err != nil {
			logrus.Errorf("invalid alert payload: %v", err)
			_ = msg.Nak()
			return
		}

		logrus.Infof(
			"Alert received for event %s (fields=%v)",
			alert.UID,
			alert.ChangedFields,
		)

		mails, err := services.GetMailsForAlert(alert)
		if err != nil {
			logrus.Errorf("cannot resolve alert recipients: %v", err)
			_ = msg.Nak()
			return
		}

		for _, mail := range mails {
			if err := services.SendAlertMail(mail, alert); err != nil {
				logrus.Errorf("mail error: %v", err)
			}
		}

		_ = msg.Ack()
	})

	if err != nil {
		return err
	}

	<-cc.Closed()
	cc.Stop()
	return nil
}
