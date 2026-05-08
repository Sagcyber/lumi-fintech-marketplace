package workers

import (
	"context"
	"database/sql"
)

type ProductLoader struct {
	db         *sql.DB
	maxWorkers int
}

func (pl *ProductLoader) LoadProductsParallel(ctx context.Context,
	productIDs []int) ([]Product, error) {
	return nil, nil
}
