package repository

import (
	"context"

	"github.com/fajarabdillahfn/shoping-gql/internal/model"
)

type Repository interface {
	GetBySku(ctx context.Context) (*model.Product, error)
}
