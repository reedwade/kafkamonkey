package config

import (
	"errors"

	"github.com/Shopify/sarama"
)

func GetProducer(certFile, keyFile, caFile, broker string) (sarama.SyncProducer, error) {
	if broker == "" {
		return nil, errors.New("no broker set")
	}

	config, err := kafkaConfig(certFile, keyFile, caFile)
	if err != nil {
		return nil, err
	}

	brokers := []string{broker}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}
