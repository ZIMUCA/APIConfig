package consumers

import (
	"context"
	"time"

	"middleware/example/internal/helpers"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
)

const (
	streamName   = "ALERTS"
	consumerName = "alerter_consumer"
)

func AlerterConsumer() (*jetstream.Consumer, error) {

	js, err := jetstream.New(helpers.NatsConn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Stream
	stream, err := js.Stream(ctx, streamName)
	if err != nil {
		stream, err = js.CreateStream(ctx, jetstream.StreamConfig{
			Name:     streamName,
			Subjects: []string{"ALERTS.*"},
		})
		if err != nil {
			return nil, err
		}
		logrus.Infof("Created stream %s", streamName)
	} else {
		logrus.Infof("Got existing stream %s", streamName)
	}

	// Consumer
	consumer, err := stream.Consumer(ctx, consumerName)
	if err != nil {
		consumer, err = stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
			Durable:       consumerName,
			Name:          consumerName,
			Description:   "Config durable consumer",
			AckPolicy:     jetstream.AckExplicitPolicy, // maitrise quand un message passe
			DeliverPolicy: jetstream.DeliverAllPolicy,  //relit les messages si nouveau consumer
		})
		if err != nil {
			return nil, err
		}
		logrus.Infof("Created consumer %s", consumerName)
	} else {
		logrus.Infof("Got existing consumer %s", consumerName)
	}

	return &consumer, nil
}
