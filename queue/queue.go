package azure

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
)

const (
	azuriteDefaultAccountName = "devstoreaccount1"
	azuriteDefaultAccountKey  = "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw=="
)

type queue struct {
	accountName string
	accountKey  string
	ctx         context.Context
	url         azqueue.QueueURL
}

func NewAzureQueue() *queue {
	return &queue{
		accountName: azuriteDefaultAccountName,
		accountKey:  azuriteDefaultAccountKey,
	}
}

func (q *queue) Create() error {
	credential, err := azqueue.NewSharedKeyCredential(q.accountName, q.accountKey)
	if err != nil {
		return err
	}
	p := azqueue.NewPipeline(credential, azqueue.PipelineOptions{})
	u, _ := url.Parse(fmt.Sprintf("https://127.0.0.1:10001/%s", q.accountName))
	serviceURL := azqueue.NewServiceURL(*u, p)
	q.ctx = context.TODO()
	q.url = serviceURL.NewQueueURL("queue-name")
	_, err = q.url.Create(q.ctx, azqueue.Metadata{})
	return err
}

func (q queue) Enqueue(message string) error {
	messagesURL := q.url.NewMessagesURL()
	_, err := messagesURL.Enqueue(q.ctx, message, time.Second*0, time.Minute)
	return err
}

func (q queue) Dequeue() (*azqueue.DequeuedMessagesResponse, error) {
	messagesURL := q.url.NewMessagesURL()
	return messagesURL.Dequeue(q.ctx, azqueue.QueueMaxMessagesDequeue, 10*time.Second)
}

func (q queue) Clear() error {
	messagesURL := q.url.NewMessagesURL()
	_, err := messagesURL.Clear(q.ctx)
	return err
}

func (q queue) Delete() error {
	_, err := q.url.Delete(q.ctx)
	return err
}
