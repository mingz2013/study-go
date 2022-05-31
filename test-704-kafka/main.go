package test_704_kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {

	conn, _ := kafka.Dial("tcp", "localhost:9092")
	conn.CreateTopics(kafka.TopicConfig{})

	conn.DeleteTopics("")

	conn.WriteMessages(kafka.Message{})

	reader := kafka.NewReader(kafka.ReaderConfig{})

	m, _ := reader.ReadMessage(context.TODO())

	fmt.Println(m)

	writer := kafka.NewWriter(kafka.WriterConfig{})

	writer.WriteMessages(context.TODO(), kafka.Message{})

	client := &kafka.Client{}
	fmt.Println(client)
	//client.

}
