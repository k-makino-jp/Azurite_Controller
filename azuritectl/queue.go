package azuritectl

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"testing"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
)

type queue struct {
	name string
	ctx  context.Context
	url  azqueue.QueueURL
}

// NewAzureQueue is constructor for queue.
func NewAzureQueue(name string) *queue {
	return &queue{
		name: name,
	}
}

// Create creates azurite queue storage.
func (q *queue) Create(t *testing.T, port int) error {
	credential, err := azqueue.NewSharedKeyCredential(azuriteDefaultAccountName, azuriteDefaultAccountKey)
	if err != nil {
		t.Errorf("Failed to azqueue.NewSharedKeyCredential = %v", err)
		return err
	}

	azqueuePipeline := azqueue.NewPipeline(credential, azqueue.PipelineOptions{})

	u, err := url.Parse(fmt.Sprintf("https://127.0.0.1:%d/%s", port, azuriteDefaultAccountName))
	if err != nil {
		t.Errorf("Failed to url.Parse = %v", err)
		return err
	}

	serviceURL := azqueue.NewServiceURL(*u, azqueuePipeline)
	q.ctx = context.TODO()
	q.url = serviceURL.NewQueueURL(q.name)
	queueCreateResponse, err := q.url.Create(q.ctx, azqueue.Metadata{})
	if err != nil {
		t.Errorf("Failed to q.url.Create = %v", err)
		return err
	}

	if isSuccessfulStatusCode(queueCreateResponse.StatusCode()) {
		return nil
	}

	respBytes, err := ioutil.ReadAll(queueCreateResponse.Response().Body)
	defer queueCreateResponse.Response().Body.Close()
	if err != nil {
		t.Errorf("Failed to ioutil.ReadAll = %v", err)
		return err
	}

	errMessage := fmt.Sprintf("q.url.Create returns invalid response. status code = %d, body = %s",
		queueCreateResponse.StatusCode(),
		string(respBytes),
	)
	return errors.New(errMessage)
}

// Enqueue enqueues message.
func (q queue) Enqueue(t *testing.T, messageText string) error {
	messagesURL := q.url.NewMessagesURL()
	visibilityTimeout := 0 * time.Second
	timeToLive := 1 * time.Minute
	enqueueMessageResponse, err := messagesURL.Enqueue(q.ctx, messageText, visibilityTimeout, timeToLive)
	if err != nil {
		t.Errorf("Failed to messagesURL.Enqueue = %v", err)
		return err
	}

	if isSuccessfulStatusCode(enqueueMessageResponse.StatusCode()) {
		return nil
	}

	respBytes, err := ioutil.ReadAll(enqueueMessageResponse.Response().Body)
	defer enqueueMessageResponse.Response().Body.Close()
	if err != nil {
		t.Errorf("Failed to ioutil.ReadAll = %v", err)
		return err
	}

	errMessage := fmt.Sprintf("messagesURL.Enqueue returns invalid response. status code = %d, body = %s",
		enqueueMessageResponse.StatusCode(),
		string(respBytes),
	)
	return errors.New(errMessage)
}

// Dequeue dequeues message from queue.
func (q queue) Dequeue(t *testing.T) (*azqueue.DequeuedMessagesResponse, error) {
	messagesURL := q.url.NewMessagesURL()
	dequeuedMessagesResponse, err := messagesURL.Dequeue(q.ctx, azqueue.QueueMaxMessagesDequeue, 10*time.Second)
	if err != nil {
		t.Errorf("Failed to messagesURL.Dequeue = %v", err)
	}

	if isSuccessfulStatusCode(dequeuedMessagesResponse.StatusCode()) {
		return dequeuedMessagesResponse, nil
	}

	respBytes, err := ioutil.ReadAll(dequeuedMessagesResponse.Response().Body)
	defer dequeuedMessagesResponse.Response().Body.Close()
	if err != nil {
		t.Errorf("Failed to ioutil.ReadAll = %v", err)
		return &azqueue.DequeuedMessagesResponse{}, err
	}

	errMessage := fmt.Sprintf("messagesURL.Dequeue returns invalid response. status code = %d, body = %s",
		dequeuedMessagesResponse.StatusCode(),
		string(respBytes),
	)
	return &azqueue.DequeuedMessagesResponse{}, errors.New(errMessage)
}

// Clear deletes all messages from a queue.
func (q queue) Clear(t *testing.T) error {
	messagesURL := q.url.NewMessagesURL()
	messagesClearResponse, err := messagesURL.Clear(q.ctx)
	if err != nil {
		t.Errorf("Failed to messagesURL.Clear = %v", err)
	}

	if isSuccessfulStatusCode(messagesClearResponse.StatusCode()) {
		return nil
	}

	respBytes, err := ioutil.ReadAll(messagesClearResponse.Response().Body)
	defer messagesClearResponse.Response().Body.Close()
	if err != nil {
		t.Errorf("Failed to ioutil.ReadAll = %v", err)
		return err
	}

	errMessage := fmt.Sprintf("messagesURL.Dequeue returns invalid response. status code = %d, body = %s",
		messagesClearResponse.StatusCode(),
		string(respBytes),
	)
	return errors.New(errMessage)
}

// Deleted deletes queue.
func (q queue) Delete(t *testing.T) error {
	queueDeleteResponse, err := q.url.Delete(q.ctx)
	if err != nil {
		t.Errorf("Failed to q.url.Delete = %v", err)
	}

	if isSuccessfulStatusCode(queueDeleteResponse.StatusCode()) {
		return nil
	}

	respBytes, err := ioutil.ReadAll(queueDeleteResponse.Response().Body)
	defer queueDeleteResponse.Response().Body.Close()
	if err != nil {
		t.Errorf("Failed to ioutil.ReadAll = %v", err)
		return err
	}

	errMessage := fmt.Sprintf("messagesURL.Dequeue returns invalid response. status code = %d, body = %s",
		queueDeleteResponse.StatusCode(),
		string(respBytes),
	)
	return errors.New(errMessage)
}

// CreateSasQueryParameters creates sas query parameters.
func (q queue) CreateSasQueryParameters(t *testing.T) (azqueue.SASQueryParameters, error) {
	credential, err := azqueue.NewSharedKeyCredential(azuriteDefaultAccountName, azuriteDefaultAccountKey)
	if err != nil {
		t.Errorf("Failed to azqueue.NewSharedKeyCredential = %v", err)
		return azqueue.SASQueryParameters{}, err
	}

	sasQueryParams := azqueue.QueueSASSignatureValues{
		Protocol:   azqueue.SASProtocolHTTPS,
		ExpiryTime: time.Now().Add(1 * time.Hour),
		QueueName:  q.name,
		Permissions: azqueue.QueueSASPermissions{
			Add:     true,
			Read:    true,
			Process: true,
		}.String(),
	}.NewSASQueryParameters(credential)
	return sasQueryParams, nil
}
