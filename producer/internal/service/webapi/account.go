package ser

import (
	"bina/internal/core"
	"bina/internal/es/kafka"
	"bytes"
	"context"
	"encoding/json"
)

type AccountWriter struct {
	writer *kafka.Writer
	topic string
	
}
type eventAcc struct {
	Type string
	Value core.Account
}

func (a AccountWriter) PublishAccountCreate(ctx context.Context, acc core.Account) error {
	return a.publish(ctx,"account.event.created",acc)
}

func (a AccountWriter) PublishAccountUpdate(ctx context.Context, acc core.Account) error {
	return a.publish(ctx,"account.event.update",acc)
}

func (a AccountWriter) PublishAccountDelete(ctx context.Context, acc core.Account) error {
	return a.publish(ctx,"account.event.delete",acc)
}

func (a AccountWriter) PublishAccountLocked(ctx context.Context, acc core.Account) error {
	return a.publish(ctx,"account.event.locked",acc)
}

func (a AccountWriter) PublishAccountUnlocked(ctx context.Context, acc core.Account) error {
	return a.publish(ctx,"account.event.unlocked",acc)
}

func (a AccountWriter) PublishAccountBalanceUpdated(ctx context.Context, acc core.Account) error {
	return a.publish(ctx,"account.event.balanceupdated",acc)
}

func NewAccountWriter(writer *kafka.Writer, topic string) *AccountWriter {
	return &AccountWriter{
		writer: writer,
		topic:  topic,
	}
}

func (a AccountWriter) publish(c context.Context,msgType string, acc core.Account) error {
	var bu bytes.Buffer
	event:= eventAcc{
		Type:  msgType,
		Value: acc,
	}
	err := json.NewEncoder(&bu).Encode(event)
	if err != nil {
		return err
	}
	return 	a.writer.WriteMessages(c, a.topic,bu.Bytes())


}

