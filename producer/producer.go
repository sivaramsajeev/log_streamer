package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sivaramsajeev/log_streamer/configs"
)

type ISender interface {
	Send() error
	Stream() error
}

type KafkaSender struct {
	Producer *kafka.Producer
	Config   configs.IConfigs
	Message  *Message
}

func NewKafkaSender() *KafkaSender {
	config := configs.GetConfig()
	producer, _ := kafka.NewProducer(config.Read())
	return &KafkaSender{
		Producer: producer,
		Config:   config,
		Message:  NewMessage(),
	}
}

func (ks *KafkaSender) Send() error {
	p := ks.Producer
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	p.Produce(ks.Message.Message, nil)
	p.Flush(15 * 1000)
	p.Close()
	return nil
}

type Message struct {
	Message *kafka.Message
	Config  *configs.MessageConfig
}

func NewMessage() *Message {
	config := configs.NewMessageConfig()
	topic := configs.MustReadEnv(configs.ConfigTopicName)
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            config.Key,
		Value:          config.Value,
	}

	return &Message{
		Message: msg,
		Config:  config,
	}
}
