package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"blog-services/common/broker"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type consumer struct {
	service PostsService
}

func NewConsumer(service PostsService) *consumer {
	return &consumer{service}
}

type LikeData struct {
	PostId string `JSON:"postId"`
	UserId string `JSON:"userId"`
}

func (c *consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare("", true, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(q.Name, "", broker.PostLikedEvent, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", d.Body)

			// Extract the headers
			ctx := broker.ExtractAMQPHeader(context.Background(), d.Headers)

			tr := otel.Tracer("amqp")
			_, messageSpan := tr.Start(ctx, fmt.Sprintf("AMQP - consume - %s", q.Name))

			likeData := &LikeData{}
			if err := json.Unmarshal(d.Body, likeData); err != nil {
				d.Nack(false, false)
				log.Printf("failed to unmarshal like data: %v", err)
				continue
			}

			_, err := c.service.IncrementLikeCount(context.Background(), likeData.PostId)
			if err != nil {
				log.Printf("failed to update post: %v", err)

				if err := broker.HandleRetry(ch, &d); err != nil {
					log.Printf("Error handling retry: %v", err)
				}

				continue
			}

			messageSpan.AddEvent("post.updated")
			messageSpan.End()

			log.Println("Post has been updated from AMQP")
			d.Ack(false)
		}
	}()
}
