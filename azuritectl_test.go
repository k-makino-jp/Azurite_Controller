package main

import (
	"testing"

	"github.com/k-makino-jp/azurite-controller/azuritectl"
	"github.com/stretchr/testify/assert"
)

const (
	queueName = "azurite-queue"
	port      = 10001
)

func Test_AzureQueue(t *testing.T) {
	testName := "EnqueueMessage"

	t.Run(testName, func(t *testing.T) {
		azureQueue := azuritectl.NewAzureQueue(queueName)
		if err := azureQueue.Create(t, port); err != nil {
			return
		}

		messageText := "sample message"
		if err := azureQueue.Enqueue(t, messageText); err != nil {
			return
		}

		dequeue, err := azureQueue.Dequeue(t)
		if err != nil {
			return
		}
		assert.Equal(t, messageText, dequeue.Message(0).Text)

		if err := azureQueue.Clear(t); err != nil {
			return
		}

		if err := azureQueue.Delete(t); err != nil {
			return
		}
	})
}
