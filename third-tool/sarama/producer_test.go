package sarama

import (
	"log"
	"testing"

	"github.com/Shopify/sarama"
)

// 采集消息（同步）
func Test_DataCollector(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	log.Println("producer init success")
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("key-1"),
		Value: sarama.StringEncoder("hello gina"),
	})
	if err != nil {
		log.Fatalln("SendMessage failed, err:", err)
	}
	log.Printf("partition = %d, offset = %d", partition, offset)
}
