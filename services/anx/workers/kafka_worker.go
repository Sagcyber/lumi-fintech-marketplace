package workers

import (
	"context"
	"database/sql"

	"github.com/segmentio/kafka-go"
)

type KafkaEventWorker struct {
	id         int
	reader     *kafka.Reader
	clickhouse *sql.DB
}

func (w *KafkaEventWorker) ProcessEvents(ctx context.Context) error {
	for {
		msg, err := w.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		w.reader.CommitMessages(ctx, msg)
	}
}
