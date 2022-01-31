package db

import (
	"context"

	"github.com/azimuth3d/woof-service/pkg/schema"
)

type DatabaseRepository interface {
	Close()
	InsertWoof(ctx context.Context, woof schema.Woof) error
	ListWoof(ctx context.Context, skip uint64, take uint64) ([]schema.Woof, error)
}

// func InsertWoof(ctx context.Context, woof schema.Woof) error {
// 	fmt.Println("insert woof ")
// 	return nil
// }
