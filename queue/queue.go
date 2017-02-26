// Sample pubsub-quickstart creates a Google Cloud Pub/Sub topic.
package queue

import (
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

func createTopic(ctx context.Context, projectID string, topicName string) (*pubsub.Topic, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	topic := client.Topic(topicName)
	ok, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if topic exists; topic: %v", err)
	}
	if !ok {
		// Creates the new topic.
		topic, err := client.CreateTopic(ctx, topicName)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}
		fmt.Printf("Topic %v created.\n", topic)
	}
	if ok {
		fmt.Printf("Topic %v already exists.\n", topic)
	}
	return topic, nil
}

// func main() {
// 	ctx := context.Background()

// 	// Sets your Google Cloud Platform project ID.
// 	projectID := os.Getenv("HERMES_PROJECT_ID")

// 	// Sets the name for the new topic.
// 	publishTopicName := os.Getenv("HERMES_PUBLISHER_TOPIC")
// 	// consumeTopicName := os.Getenv("HERMES_CONSUMER_TOPIC")

// 	p, err := createTopic(ctx, projectID, publishTopicName)
// 	if err != nil {
// 		log.Fatalf("Failed to create topic: %v", err)
// 	}

// 	// c, err := createTopic(ctx, projectID, consumeTopicName)
// 	if err != nil {
// 		log.Fatalf("Failed to create topic: %v", err)
// 	}

// 	haikunator := haikunator.New(time.Now().UTC().UnixNano())
// 	msg := haikunator.Haikunate()
// 	msgIDs, err := p.Publish(ctx, &pubsub.Message{
// 		Data: []byte(msg),
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to publish msg %v", err)
// 	}
// 	for _, id := range msgIDs {
// 		fmt.Printf("Published a message; msg ID: %v\n", id)
// 	}

// 	client, err := pubsub.NewClient(ctx, projectID)
// 	// subscription, err := client.CreateSubscription(ctx, "hermes-subscriber", p, 20*time.Second, nil)
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to create subscription %v", err)
// 	// }
// 	// fmt.Printf("Created subscription: %v\n", subscription)

// 	sub := client.Subscription("hermes-subscriber")
// 	it, err := sub.Pull(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create subscription %v", err)
// 	}
// 	defer it.Stop()

// 	// Consume 2 messages.
// 	for i := 0; i < 2; i++ {
// 		msg, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Failed to consume msg %v", err)
// 		}
// 		fmt.Printf("Got message: %q\n", string(msg.Data))
// 		msg.Done(true)
// 	}
// }
