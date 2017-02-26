// Sample pubsub-quickstart creates a Google Cloud Pub/Sub topic.
package main

import (
	"fmt"
	"log"
	"os"
	// "time"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

var consume = &Command{
	Run:       runConsume,
	UsageLine: "consume",
	Short:     "consume queues",
	Long:      "'consume' continually consumes from the queue",
}

func runConsume(cmd *Command, args []string) (err error) {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("HERMES_PROJECT_ID")

	// Sets the name for the new topic.
	// publishTopicName := os.Getenv("HERMES_PUBLISHER_TOPIC")
	// consumeTopicName := os.Getenv("HERMES_CONSUMER_TOPIC")

	client, err := pubsub.NewClient(ctx, projectID)

	// p, err := createTopic(ctx, projectID, publishTopicName)
	// if err != nil {
	// 	log.Fatalf("Failed to create topic: %v", err)
	// }

	// subscription, err := client.CreateSubscription(ctx, "hermes-subscriber", p, 20*time.Second, nil)
	// if err != nil {
	// 	log.Fatalf("Failed to create subscription %v", err)
	// }
	// fmt.Printf("Created subscription: %v\n", subscription)

	sub := client.Subscription("hermes-subscriber")
	it, err := sub.Pull(ctx)
	if err != nil {
		log.Fatalf("Failed to create subscription %v", err)
	}
	defer it.Stop()

	// Consume 2 messages.
	for {
		msg, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to consume msg %v", err)
		}
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Done(true)
	}
	return err
}
