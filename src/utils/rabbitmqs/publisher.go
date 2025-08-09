package rabbitmqs

import (
	"context"
	"encoding/json"
	"fmt"
	"publisher-topic/src/helpers"
	"reflect"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMessage(ctx context.Context, queueName, routingKey string, message interface{}) (map[string]interface{}, error) {
	ch, err := GetChannel()
	if err != nil {
		return nil, err
	}

	replyQueue, err := ch.QueueDeclare(
		"",
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a reply queue: %w", err)
	}

	msgs, err := ch.Consume(
		replyQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}

	corrId := helpers.RandomString(32)
	body, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %w", err)
	}

	// Kirim request ke consumer
	err = ch.PublishWithContext(ctx,
		"",        // default exchange (direct)
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       replyQueue.Name,
			Body:          body,
		})
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	// Tunggu response
	for {
		select {
		case d := <-msgs:

			if d.CorrelationId == corrId {
				var response map[string]interface{}

				if err := json.Unmarshal(d.Body, &response); err != nil {
					fmt.Println("Error unmarshalling response:", err)
					return nil, fmt.Errorf("failed to parse response: %w", err)
				}

				return response, nil
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout waiting for response")
		}
	}
}

func checkType(v interface{}) string {
	t := reflect.TypeOf(v)
	fmt.Printf("Variable: %v, Type: %v, Kind: %v\n", v, t, t.Kind())

	if t.Kind() == reflect.Array {
		fmt.Println("  This is an array.")
	} else if t.Kind() == reflect.Slice {
		fmt.Println("  This is a slice.")
	} else {
		fmt.Println("  This is neither an array nor a slice.")
	}

	return fmt.Sprintf("Type: %T, Value: %v", v, v)
}
