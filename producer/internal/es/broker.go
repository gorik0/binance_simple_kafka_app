package es

import (
	"bina/internal/core"
	"bina/internal/es/kafka"
	ser "bina/internal/service/webapi"
	"context"
)

type AccountMessageBroker interface {
PublishAccountCreate(ctx context.Context,acc core.Account)error
PublishAccountUpdate(ctx context.Context,acc core.Account)error
PublishAccountDelete(ctx context.Context,acc core.Account)error
PublishAccountLocked(ctx context.Context,acc core.Account)error
PublishAccountUnlocked(ctx context.Context,acc core.Account)error
PublishAccountBalanceUpdated(ctx context.Context,acc core.Account)error
}


type TransferMessageBroker interface {
PublishTransferCreate(ctx context.Context,acc core.Transfer)error
}

type MessageBroker struct {
	Account AccountMessageBroker
	Transfer TransferMessageBroker

}


func NewKafkaMessageBroker(writer *kafka.Writer) *MessageBroker{
return &MessageBroker{
	Account:  ser.NewAccountWriter(writer,writer.Topic),
	Transfer: ser.NewTransferWriter(writer,writer.Topic),
}
}



