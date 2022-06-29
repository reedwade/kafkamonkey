package config

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/Shopify/sarama"
)

// func Ping(certFile, keyFile, caFile, broker string) error {

// 	client, err := GetClient(certFile, keyFile, caFile, broker)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("BROKERS\n")
// 	brokers := client.Brokers()
// 	// sort.Sort(brokers)
// 	for _, v := range brokers {
// 		fmt.Printf("broker id%v %v\n", v.ID(), v.Addr())
// 	}

// 	topics, err := client.Topics()
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("TOPICS (%v)\n", len(topics))
// 	sort.Strings(topics)
// 	for _, topic := range topics {
// 		fmt.Printf("%v - \n", topic)
// 		// client.InSyncReplicas()
// 	}

// 	return nil
// }

func kafkaConfig(certFile, keyFile, caFile string) (*sarama.Config, error) {
	keypair, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{keypair},
		RootCAs:      caCertPool,
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Net.TLS.Enable = true
	config.Net.TLS.Config = tlsConfig
	config.Version = sarama.V0_10_2_0

	return config, nil
}
