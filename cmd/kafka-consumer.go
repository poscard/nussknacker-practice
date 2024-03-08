package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	// 创建一个新的读取器，连接到指定的 Kafka 代理和主题
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "my-new-topic",
		GroupID:  "my-new-topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()

	for {
		// 从 Kafka 读取消息
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
