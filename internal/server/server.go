package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	audit "github.com/BalamutDiana/crud_audit/pkg/domain"
	"github.com/streadway/amqp"
)

type Server struct {
	auditServer *AuditServer
	conn        *amqp.Connection
	channel     *amqp.Channel
	queue       amqp.Queue
}

func New(auditServer *AuditServer, port int) *Server {
	adr := fmt.Sprintf("amqp://guest:guest@localhost:%d/", port)
	conn, err := amqp.Dial(adr)
	if err != nil {
		log.Fatal("failed to connect to rabbitmq")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel")
	}

	q, err := ch.QueueDeclare(
		"audit_logs", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal("failed to declare a queue")
	}

	return &Server{
		conn:        conn,
		channel:     ch,
		queue:       q,
		auditServer: auditServer,
	}
}

func (s *Server) CloseConnection() error {
	if err := s.conn.Close(); err != nil {
		return err
	}
	if err := s.channel.Close(); err != nil {
		return err
	}
	return nil
}

func (s *Server) ListenAndServe() error {
	msgs, err := s.channel.Consume(
		s.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Fatal("failed to register a consumer")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var reqItem audit.LogItem
			if err := json.Unmarshal(d.Body, &reqItem); err != nil {
				log.Fatal("failed to unmarshal request")
			}
			if err := s.auditServer.service.Insert(context.TODO(), reqItem); err != nil {
				log.Fatal("failed to insert item")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
