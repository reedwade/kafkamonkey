package config

import (
	"errors"

	"github.com/Shopify/sarama"
)

func GetClient(useTLS bool, certFile, keyFile, caFile, broker string) (sarama.Client, error) {
	if broker == "" {
		return nil, errors.New("no broker set")
	}

	config, err := kafkaConfig(useTLS, certFile, keyFile, caFile)
	if err != nil {
		return nil, err
	}

	return sarama.NewClient([]string{broker}, config)
}
