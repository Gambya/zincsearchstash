package broker

import (
	"context"
	"encoding/json"
	"zincsearchstash/internal/setup"

	"github.com/makasim/amqpextra"
	"github.com/makasim/amqpextra/consumer"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

// Declaring record handler to consume message from log
type RecordHandler = func(
	message interface{},
) error

type RMQClient struct {
	Dialer *amqpextra.Dialer
	cfg    *setup.Config
}

func NewRMQClient(cfg *setup.Config, url string) (res *RMQClient) {
	d, err := amqpextra.NewDialer(amqpextra.WithURL(url))
	if err != nil {
		log.Fatal().Err(err).Msg("error while connecting to rabbitmq")
	}

	res = &RMQClient{
		Dialer: d,
		cfg:    cfg,
	}
	return
}

func (p *RMQClient) RecordListening(ctx context.Context, handler RecordHandler) (err error) {
	log.Trace().Msg("starting listening a new record")
	h := consumer.HandlerFunc(func(ctx context.Context, msg amqp.Delivery) interface{} {
		defer msg.Ack(false)
		var data interface{}
		err = json.Unmarshal(msg.Body, &data)
		if err != nil {
			return nil
		}
		err = handler(data)
		// If the handler throws any error, we will not make ack of the message
		// and the message will came again. The handler must process this message
		// right to get it off the queue. If the problem is a poison message, then the
		// handler must handle it and get it of the queue
		if err != nil {
			// TODO: Log message here!
		}

		return nil
	})

	err = func() (err error) {
		con, err := p.Dialer.Connection(ctx)
		if err != nil {
			return
		}
		ch, err := con.Channel()
		if err != nil {
			return
		}
		defer ch.Close()

		// create exchange
		err = ch.ExchangeDeclare(p.cfg.Exchange,
			"direct",
			true,  // durable
			false, // auto-delete
			false, // internal
			false, // nowait
			nil)
		if err != nil {
			return
		}

		queue, err := ch.QueueDeclare(
			p.cfg.QueueName,
			true, false, false, false, nil)
		if err != nil {
			log.Error().Err(err).Msg("error on queue declare")
			return
		}

		err = ch.QueueBind(queue.Name, p.cfg.RoutingKey, p.cfg.Exchange, false, nil)
		if err != nil {
			log.Error().Err(err).Msg("error on queue bind")
			return
		}

		return
	}()
	if err != nil {
		log.Error().Err(err).Msg("error on connect to rabbitmq")
		return
	}

	options := []consumer.Option{}
	options = append(options, consumer.WithHandler(h))
	options = append(options, consumer.WithQueue(p.cfg.QueueName))
	c, err := p.Dialer.Consumer(options...)
	if err != nil {
		return
	}
	defer c.Close()
	<-ctx.Done()
	return
}
