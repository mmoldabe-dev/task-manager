package rabbit

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/mmoldabe-dev/task-manager/config"
	"github.com/rabbitmq/amqp091-go"
)

func RabbitClient(cfg config.RabbitMQConfig) (*amqp091.Connection, *amqp091.Channel, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	var conn *amqp091.Connection
	var err error
	for i := 1; i <= 5; i++ {
		slog.Info("Connecting to RabbitMQ...", "attempt", i)
		conn, err = amqp091.Dial(url)
		if err == nil {
			ch, chErr := conn.Channel()
			if chErr != nil {
				conn.Close()
				return nil, nil, fmt.Errorf("failed to open channel: %w", chErr)
			}
			slog.Info("Successfully connected to RabbitMQ")
			return conn, ch, nil
		}
		slog.Warn("Failed to connect to RabbitMQ, retrying...", "attempt", i, "err", err)
		time.Sleep(3 * time.Second)
	}

	return nil, nil, fmt.Errorf("failed to connect RabbitMQ after retries: %w", err)
}
