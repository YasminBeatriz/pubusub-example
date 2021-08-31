package pubsub

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

type PubSub struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func New(projectID, topicID string) *PubSub {
	context := context.Background()
	client, err := pubsub.NewClient(context, projectID)

	if err != nil {
		log.Fatalf("Failed to create PubSub client %v", err)
	}

	topic := client.Topic(topicID)

	return &PubSub{
		client: client,
		topic:  topic,
	}
}

func (p *PubSub) Publish(projectID, topicID string, msg *pubsub.Message) error {
	context := context.Background()

	topic := p.topic
	result := topic.Publish(context, msg)
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(context)
	if err != nil {
		return fmt.Errorf("error when publishing message: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v %v\n", id, err)
	return nil
}

func (p *PubSub) Subscribe(projectID, subID string) error {
	ctx := context.Background()

	sub := p.client.Subscription(subID)
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Fprintf(os.Stdout, "Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("receive: %v", err)
	}

	/* err = models.SaveUser(user.Username, user.Name)
	if err != nil {
		panic(err)
	} */
	return nil
}

func (p *PubSub) Message(message []byte) *pubsub.Message {
	return &pubsub.Message{Data: message}
}

func (p *PubSub) GetTopicID() string {
	return p.topic.ID()
}

func (p *PubSub) Close() {
	p.client.Close()
}
