package ser

import (
	"bina/internal/core"
	"bina/internal/es/kafka"
	"bytes"
	"context"
	"encoding/json"
)

type TransferWriter struct {
	writer *kafka.Writer
	topic string

}

type eventTra struct {
	Type string
	Value core.Transfer

}
func (t TransferWriter) PublishTransferCreate(ctx context.Context, transfer core.Transfer) error {
return t.publish(ctx,"transfer.eventTra.created", transfer)
}

func NewTransferWriter(writer *kafka.Writer,topic string) *TransferWriter {
	return &TransferWriter{
		writer: writer,
		topic:  topic,
	}
}

func (t TransferWriter) publish(c context.Context,msgType string, tra core.Transfer) error {
	var bu bytes.Buffer
event:= eventTra{
	Type:  msgType,
	Value: tra,
}
	err := json.NewEncoder(&bu).Encode(event)
	if err != nil {
		return err
	}
return 	t.writer.WriteMessages(c,t.topic,bu.Bytes())


}
