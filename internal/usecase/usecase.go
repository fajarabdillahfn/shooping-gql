package usecase

import (
	"context"

	"github.com/fajarabdillahfn/shoping-gql/internal/model"
)

type UseCase interface {
	Checkout(ctx context.Context, productsBought map[string]int) (*model.Cart, error)
}
