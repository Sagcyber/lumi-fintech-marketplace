package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelProductLoading(t *testing.T) {
	loader := &ProductLoader{maxWorkers: 5}

	productIDs := []int{1, 2, 3, 4, 5}
	products, err := loader.LoadProductsParallel(context.Background(), productIDs)

	assert.NoError(t, err)
	assert.Equal(t, len(products), 5)
}
