package messages

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func MakeMessages(topicName string, messageCount, messageValueLength int) ([]*sarama.ProducerMessage, error) {

	// Create a byte array this length full of zeros for the message value.
	value := make([]byte, messageValueLength)

	messages := make([]*sarama.ProducerMessage, messageCount)
	for i := 0; i < messageCount; i++ {
		messages[i] = &sarama.ProducerMessage{
			Topic: topicName,
			Key:   sarama.StringEncoder(fmt.Sprintf("key-%05d", i)),
			Value: sarama.ByteEncoder(value),
		}
	}
	return messages, nil
}
