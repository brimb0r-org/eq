package produce

import (
	"fmt"
	"github.com/brimb0r-org/eq/application/internal/translator/eq_translator"
	"github.com/brimb0r-org/eq/application/pkg/worker_pool"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type IProduce interface {
	Produce(incoming <-chan eq_translator.ITranslator) error
}

type Produce struct {
	Producer *kafka.Producer
}

func NewProducer() IProduce {
	// https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:27777"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	buildProducerService := &Produce{
		Producer: p,
	}

	return buildProducerService
}

func (p *Produce) Produce(incoming <-chan eq_translator.ITranslator) error {
	var err error
	totalMsgcnt := len(incoming)
	fmt.Printf("incoming %v", totalMsgcnt)

	eventWorker := worker_pool.Worker(func(i interface{}) {})
	for i := 0; i < totalMsgcnt; i++ {
		eventWorker <- p.events()
	}

	for t := range incoming {
		err = p.produceOne(t, err)
		if err != nil {
			return err
		}
	}

	// Flush and close the producer and the events channel
	for p.Producer.Flush(10000) > 0 {
		fmt.Print("Still waiting to flush outstanding messages\n")
	}
	close(eventWorker)
	return err
}

func (p *Produce) produceOne(t eq_translator.ITranslator, err error) error {
	err = p.Producer.Produce(t.Translate(), nil)

	if err != nil {
		if err.(kafka.Error).Code() == kafka.ErrQueueFull {
			// Producer queue is full, wait 1s for messages
			// to be delivered then try again.
			time.Sleep(time.Second)
			return err.(kafka.Error)
		}
		fmt.Printf("Failed to produce message: %v\n", err)
		return err
	}
	err = t.SendSuccessCallbackPublished()
	if err != nil {
		return err
	}
	return err
}

func (p *Produce) events() error {
	go func() {
		for e := range p.Producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				// The message delivery report, indicating success or
				// permanent failure after retries have been exhausted.
				// Application level retries won't help since the client
				// is already configured to do that.
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
			case kafka.Error:
				// Generic client instance-level errors, such as
				// broker connection failures, authentication issues, etc.
				//
				// These errors should generally be considered informational
				// as the underlying client will automatically try to
				// recover from any errors encountered, the application
				// does not need to take action on them.
				fmt.Printf("Error: %v\n", ev)
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()
	return nil
}
