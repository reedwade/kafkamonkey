package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/reedwade/kafkamonkey/config"
	"github.com/reedwade/kafkamonkey/messages"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

const (
	topicName = "test-topic"

	workerCountOpt              = "workers"
	workerCountDefault          = 5
	batchCountOpt               = "batches"
	batchCountDefault           = 10
	messageCountPerBatchOpt     = "messages-per-batch"
	messageCountPerBatchDefault = 10
	messageValueLengthOpt       = "message-length"
	messageValueLengthDefault   = 10
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

		producer, err := config.GetProducer(certFile, keyFile, caFile, broker)
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

		worker := func(worker int, messageBatches <-chan []*sarama.ProducerMessage, wg *sync.WaitGroup) {
			log := log.WithField("worker", worker)
			log.Info("starting")
			for messageBatch := range messageBatches {
				time.Sleep(time.Second)
				t1 := time.Now()
				producer.SendMessages(messageBatch)
				log.
					WithField("topic", messageBatch[0].Topic).
					WithField("time_taken", time.Since(t1)).
					Info("batch sent")
			}
			log.Info("done")
			wg.Done()
		}

		workChannel := make(chan []*sarama.ProducerMessage)

		wg := &sync.WaitGroup{}
		wg.Add(workerCount)

		for i := 1; i <= workerCount; i++ {
			go worker(i, workChannel, wg)
		}

		for i := 0; i < batchCount; i++ {
			batchTopicName := fmt.Sprintf("%v-%v", topicName, i)
			messages, _ := messages.MakeMessages(batchTopicName, messageCountPerBatch, messageValueLength)
			workChannel <- messages
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
	produceCmd.Flags().Int(batchCountOpt, batchCountDefault, "how many batches - each batch gets a new topic")
	produceCmd.Flags().Int(messageCountPerBatchOpt, messageCountPerBatchDefault, "")
	produceCmd.Flags().Int(messageValueLengthOpt, messageValueLengthDefault, "")
}
