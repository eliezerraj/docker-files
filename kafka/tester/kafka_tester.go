package main

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "time"

    "github.com/segmentio/kafka-go"
    "github.com/segmentio/kafka-go/sasl/plain"
)

type Payload struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    topic := "topic.local.test.01"
    mechanism := plain.Mechanism{Username: "admin", Password: "admin"}
    brokers := []string{"127.0.0.1:9092"}

    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    dialer := &kafka.Dialer{
        Timeout:       3 * time.Second,
        DualStack:     true,
        SASLMechanism: mechanism,
    }
    conn, err := dialer.DialContext(ctx, "tcp", brokers[0])
    if err != nil {
        log.Fatalf("connect failed: %v", err)
    } else {
        log.Println("connected to kafka broker")
    }
    defer conn.Close()

    err = conn.CreateTopics(kafka.TopicConfig{
        Topic:             topic,
        NumPartitions:     1,
        ReplicationFactor: 1,
    })
    var kafkaErr kafka.Error
    if err != nil && (!errors.As(err, &kafkaErr) || kafkaErr != kafka.TopicAlreadyExists) {
        log.Fatalf("topic creation failed: %v", err)
    } else {
        log.Printf("topic %s created or already exists", topic)
    }

    w := &kafka.Writer{
        Addr:     kafka.TCP(brokers...),
        Topic:    topic,
        Balancer: &kafka.LeastBytes{},
        Transport: &kafka.Transport{
            SASL: mechanism,
        },
    }
    defer w.Close()

    key := "teste-01:abc-123"
	payload := Payload{ ID: 1,
						Name: "my msg 1 from producer-01",
						}
	payload_bytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("failed to marshal payload : %s", err)
	} else {
        log.Println("payload marshaled successfully")
    }
                            
    if err := w.WriteMessages(ctx, kafka.Message{
        Key:   []byte(key),
        Value: payload_bytes,
    }); err != nil {
        log.Fatalf("write failed: %v", err)
    } else {
        log.Println("message written to kafka:" , " key: ", key, " value: ", string(payload_bytes))
    }

    fmt.Println("Kafka write succeeded")

    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: brokers,
        Topic:   topic,
        GroupID: "go-test-group",
        Dialer: &kafka.Dialer{
            Timeout:       10 * time.Second,
            DualStack:     true,
            SASLMechanism: mechanism,
        },
        MaxBytes: 10e6,
    })
    defer r.Close()

    msg, err := r.ReadMessage(ctx)
    if err != nil {
        log.Fatalf("read failed: %v", err)
    } else {
        log.Println("message read from kafka")
    }

    fmt.Printf("message received: %s\n", msg.Value)
}
