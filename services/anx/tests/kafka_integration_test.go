package tests

import (
	"context"
	"testing"
)

func TestKafkaEventProcessing(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

}
