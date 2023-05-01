package message_broker

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

const (
	CrawlingTopic = "topic-A"

	Broker1 = "localhost:9092"
)

type Kafka struct {
}

func New() Kafka {
	return Kafka{}
}

//

func (q Kafka) SubmitMessages(ctx context.Context, TopicName string, key []byte, DTO ...interface{}) error {
	var messages []kafka.Message
	for _, d := range DTO {
		var bu bytes.Buffer        // Stand-in for a network connection
		enc := gob.NewEncoder(&bu) // Will write to network.
		// Encode (send) the value.
		err := enc.Encode(d)
		if err != nil {
			log.Fatal("encode error:", err)
		}
		messages = append(messages, kafka.Message{
			Key:   key,
			Value: bu.Bytes(),
		})

	}
	kafkaWriter := &kafka.Writer{
		Addr:     kafka.TCP([]string{Broker1}...),
		Topic:    CrawlingTopic,
		Balancer: &kafka.LeastBytes{},
	}
	err := kafkaWriter.WriteMessages(ctx,
		messages...,
	)
	if err != nil {
		return err
	}

	if err = kafkaWriter.Close(); err != nil {
		return err
	}
	log.Printf("message submitted")
	return nil
}

func (q Kafka) Listen(ctx context.Context, TopicName string, v interface{}, wg *sync.WaitGroup) chan interface{} {
	ch := make(chan interface{}, 5)
	go func(c chan interface{}, wg *sync.WaitGroup) {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{Broker1},
			Topic:     TopicName,
			Partition: 0,
			MinBytes:  10e3, // 10KB
			GroupID:   "abbas",
			MaxBytes:  10e6, // 10MB
		})
		for {
			m, err := r.ReadMessage(ctx)

			fmt.Printf("message offset ==========> %v", m.Offset)

			if err != nil {
				log.Fatal(err, "eee")
			}
			fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
			dec := gob.NewDecoder(bytes.NewReader(m.Value))
			err = dec.Decode(v)
			if err != nil {
				log.Fatal(err)
			}
			c <- v
		}
	}(ch, wg)
	return ch
}

func init() {
	gob.Register(uint(0))
}
