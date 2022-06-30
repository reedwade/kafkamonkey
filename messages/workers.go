package messages

import (
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func NewWorkerPool(log *logrus.Entry, workerCount int, producer sarama.SyncProducer) (*sync.WaitGroup, chan MessagesAndBatchID) {
	worker := func(worker int, messageBatches <-chan MessagesAndBatchID, wg *sync.WaitGroup) {
		log := log.WithField("worker", worker)
		log.Info("starting")
		for messageBatch := range messageBatches {
			time.Sleep(time.Second)
			t1 := time.Now()
			producer.SendMessages(messageBatch.Messages)
			log.
				WithField("topic", messageBatch.Messages[0].Topic).
				WithField("time_taken", time.Since(t1)).
				WithField("batch", messageBatch.ID).
				Info("batch sent")
		}
		log.Info("done")
		wg.Done()
	}

	workChannel := make(chan MessagesAndBatchID)

	wg := &sync.WaitGroup{}
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go worker(i, workChannel, wg)
	}

	return wg, workChannel
}
