package consumer

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sivaramsajeev/log_streamer/configs"
)

type KafkaReceiver struct {
	Consumer *kafka.Consumer
	Config   configs.IConfigs
}

func NewKafkaReceiver() *KafkaReceiver {
	config := configs.GetConfig()
	cfg := config.Read()
	cfg.Set("group.id=log_streamer")
	cfg.Set("auto.offset.reset=earliest")
	consumer, _ := kafka.NewConsumer(cfg)
	return &KafkaReceiver{
		Consumer: consumer,
		Config:   config,
	}
}

func (ks *KafkaReceiver) Receive() error {

	topic := configs.MustReadEnv(configs.ConfigTopicName)

	consumer := ks.Consumer

	consumer.SubscribeTopics([]string{topic}, nil)

Outer:
	for {
		e := consumer.Poll(1000)
		switch ev := e.(type) {
		case *kafka.Message:
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", ev)
			break Outer
		}
	}

	consumer.Close()
	return nil
}
