package cmd

import (
	"fmt"
	"sort"

	"github.com/reedwade/kafkamonkey/config"
	"github.com/spf13/cobra"
)

var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "list the topics and in sync replica situation by partition",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.GetClient(certFile, keyFile, caFile, broker)
		if err != nil {
			log.Error(err)
			return
		}

		topics, err := client.Topics()
		if err != nil {
			log.Error(err)
			return
		}
		fmt.Printf("%v topics\n", len(topics))
		sort.Strings(topics)
		for _, topic := range topics {
			fmt.Printf("  %v - isr ", topic)
			partitions, err := client.Partitions(topic)
			if err != nil {
				log.Error(err)
				return
			}
			sortInt32Slice(partitions)
			for _, partition := range partitions {
				isr, err := client.InSyncReplicas(topic, partition)
				if err != nil {
					log.Error(err)
					return
				}
				sortInt32Slice(isr)
				fmt.Printf("%v:%v", partition, isr)
			}
			fmt.Println()
		}
	},
}

func sortInt32Slice(slice []int32) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func init() {
	rootCmd.AddCommand(topicsCmd)
}
