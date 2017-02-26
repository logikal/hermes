// Sample pubsub-quickstart creates a Google Cloud Pub/Sub topic.
package publisher

import (
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/yelinaung/go-haikunator"
	"golang.org/x/net/context"
	"github.com/logikal/hermes/queue"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("HERMES_PROJECT_ID")

	// Sets the name for the new topic.
	publishTopicName := os.Getenv("HERMES_PUBLISHER_TOPIC")
	// consumeTopicName := os.Getenv("HERMES_CONSUMER_TOPIC")

	p, err := createTopic(ctx, projectID, publishTopicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	// c, err := createTopic(ctx, projectID, consumeTopicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	haikunator := haikunator.New(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
		msg := haikunator.Haikunate()
		msgIDs, err := p.Publish(ctx, &pubsub.Message{
			Data: []byte(msg),
		})
		if err != nil {
			log.Fatalf("Failed to publish msg %v", err)
		}
		for _, id := range msgIDs {
			fmt.Printf("Published a message; msg ID: %v\n", id)
		}
	}
}
