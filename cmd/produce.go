package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/reedwade/kafkamonkey/config"
	"github.com/reedwade/kafkamonkey/messages"

	"github.com/spf13/cobra"
)

const (
	workerCountOpt              = "workers"
	workerCountDefault          = 5
	batchCountOpt               = "batches"
	batchCountDefault           = 10
	messageCountPerBatchOpt     = "messages-per-batch"
	messageCountPerBatchDefault = 10
	messageValueLengthOpt       = "message-length"
	messageValueLengthDefault   = 10
	topicOpt                    = "topic"
	topicDefault                = "monkey"
)

var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "generate a lot of messages and send them to your kafka",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		workerCount, _ := cmd.Flags().GetInt(workerCountOpt)
		batchCount, _ := cmd.Flags().GetInt(batchCountOpt)
		messageCountPerBatch, _ := cmd.Flags().GetInt(messageCountPerBatchOpt)
		messageValueLength, _ := cmd.Flags().GetInt(messageValueLengthOpt)
		topic, _ := cmd.Flags().GetString(topicOpt)

		producer, err := config.GetProducer(!skipTLS, certFile, keyFile, caFile, broker)
		if err != nil {
			log.Error(err)
			return
		}

		totalT1 := time.Now()

		log.
			WithField(workerCountOpt, workerCount).
			WithField(batchCountOpt, batchCount).
			WithField(messageCountPerBatchOpt, messageCountPerBatch).
			WithField(messageValueLengthOpt, messageValueLength).
			Info("starting")

		wg, workChannel := messages.NewWorkerPool(log, workerCount, producer)

		now := time.Now().UTC().Format("2006-01-02T150405")

		for i := 0; i < batchCount; i++ {
			batchTopic := strings.ReplaceAll(topic, "BATCHID", fmt.Sprintf("%05d", i))
			batchTopic = strings.ReplaceAll(batchTopic, "NOW", now)
			// batchTopicName := fmt.Sprintf("%v-%v", topicName, i)
			m, _ := messages.MakeMessages(batchTopic, messageCountPerBatch, messageValueLength)
			workChannel <- messages.MessagesAndBatchID{
				Messages: m,
				ID:       i,
			}
		}
		close(workChannel)

		wg.Wait()

		producer.Close()
		log.
			WithField("total_time_taken", time.Since(totalT1)).
			Info("totals")
	},
}

func init() {
	rootCmd.AddCommand(produceCmd)

	produceCmd.Flags().Int(workerCountOpt, workerCountDefault, "how many concurrent workers")
	produceCmd.Flags().Int(batchCountOpt, batchCountDefault, "how many batches")
	produceCmd.Flags().Int(messageCountPerBatchOpt, messageCountPerBatchDefault, "")
	produceCmd.Flags().String(topicOpt, topicDefault, "'BATCHID' is replaced with the batch number, 'NOW' is replaced with time and date (ex: 2022-06-29T231846)")
	produceCmd.Flags().Int(messageValueLengthOpt, messageValueLengthDefault, "")
}
