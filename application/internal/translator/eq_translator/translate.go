package eq_translator

import (
	"fmt"
	"github.com/brimb0r-org/eq/application/internal/eq_repo"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

type ITranslator interface {
	SendSuccessCallbackPublished() error
	SendSuccessCallbackConsumed() error
	Translate() *kafka.Message
}

type EqTranslator struct {
	Eq   *eq_repo.Eq
	Repo eq_repo.IEqRepo
}

func (t *EqTranslator) SendSuccessCallbackPublished() error {
	return t.Repo.UpdateEqPublished(t.Eq)
}

func (t *EqTranslator) SendSuccessCallbackConsumed() error {
	return t.Repo.UpdateEqPublished(t.Eq)
}

func (t *EqTranslator) Translate() *kafka.Message {
	value := fmt.Sprintf("Producer , message #%v", t.Eq)
	topic := os.Args[1]
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}
	return msg
}
