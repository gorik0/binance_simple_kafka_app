package kafka

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
)
type Writer struct {

	Writer *kafkago.Writer
	Topic string

}

type Config struct {
	Addres string
	Topic string
}


func NewKafkaWriter(cfg *Config) *Writer{
	writer:= kafkago.Writer{
		Addr:                   kafkago.TCP(cfg.Addres),
		Topic:                  cfg.Topic,
	}
	println("addr ::: ",cfg.Addres)
	return &Writer{
		Writer: &writer,
		Topic:  cfg.Topic,
	}
}


func (w *Writer) WriteMessages(ctx context.Context, topic string, messages ...[]byte) error {
	for _, message := range messages {
		err := w.Writer.WriteMessages(ctx, kafkago.Message{
			Topic: "",
			Value: message,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
