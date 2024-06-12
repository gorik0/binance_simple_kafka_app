package kafka

import  kafkago "github.com/segmentio/kafka-go"
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
	return &Writer{
		Writer: &writer,
		Topic:  cfg.Topic,
	}
}
