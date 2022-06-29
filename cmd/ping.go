package cmd

import (
	"fmt"
	"sort"

	"github.com/reedwade/kafkamonkey/config"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "test kafka broker connectivity",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.GetClient(certFile, keyFile, caFile, broker)
		if err != nil {
			log.Error(err)
			return
		}

		brokers := client.Brokers()
		fmt.Printf("%v brokers\n", len(brokers))
		sort.Slice(brokers, func(i, j int) bool {
			return brokers[i].ID() < brokers[j].ID()
		})
		for _, v := range brokers {
			fmt.Printf("  id%v %v\n", v.ID(), v.Addr())
		}

		topics, err := client.Topics()
		if err != nil {
			log.Error(err)
			return
		}
		fmt.Printf("%v topics\n", len(topics))
		sort.Strings(topics)
		if len(topics) > 5 {
			topics = topics[:5]
			topics = append(topics, "...")
		}
		if len(topics) > 0 {
			fmt.Printf("  %v\n", topics)
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
