package produce

import (
	"fmt"
	"github.com/brimb0r-org/eq/application/internal/translator/eq_translator"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type IProduce interface {
	Produce(incoming <-chan eq_translator.ITranslator) error
}

type Produce struct {
	translator *eq_translator.EqTranslator
}

func (p *Produce) Produce(incoming <-chan eq_translator.ITranslator) error {
	var err error
	for t := range incoming {
		t.Translate()
		if err != nil {
			return err
		}
		go func(t eq_translator.ITranslator) error {
			err = t.SendSuccessCallback()
			if err != nil {
				return err
			}
			return err
		}(t)
	}
	return err
}

func NewProducer() *kafka.Producer {
	// https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": ""})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)
	return p
}
