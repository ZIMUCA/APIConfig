package helpers

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var NatsConn *nats.Conn

func ConnectNats() error {
	if NatsConn != nil && NatsConn.IsConnected() {
		return nil
	}

	// URL du serveur NATS
	url := nats.DefaultURL // "nats://localhost:4222", à adapter si nécessaire

	opts := []nats.Option{
		nats.Name("APIConfig Alerter"),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2 * time.Second),
		nats.Timeout(10 * time.Second),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			logrus.Warnf("NATS disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logrus.Infof("Reconnected to NATS (%s)", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			logrus.Warn("NATS connection closed")
		}),
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return err
	}

	NatsConn = nc
	logrus.Infof("Connected to NATS at %s", url)
	return nil
}

func CloseNats() {
	if NatsConn != nil && !NatsConn.IsClosed() {
		NatsConn.Close()
		logrus.Info("NATS connection closed")
	}
}
