package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var Client *firestore.Client

// InitFirestore initializes the Firestore client.
// It retrieves the Google Cloud project ID from the environment variable,
// creates a Firestore client, and assigns it to the global Client variable.
func InitFirestore() {
	// Get the Google Cloud project ID from the environment variable
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatalf("GOOGLE_CLOUD_PROJECT environment variable is missing")
	}

	// Create a new context
	ctx := context.Background()

	// Create a new Firestore client
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	// Assign the created client to the global Client variable
	Client = client
}
