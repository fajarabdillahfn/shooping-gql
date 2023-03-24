package gql

import (
	"context"

	model "github.com/fajarabdillahfn/shoping-gql/internal/model"
)

// Checkout is the resolver for the checkout field.
func (r *mutationResolver) Checkout(ctx context.Context, input []*model.ProductInput) (*model.Cart, error) {
	productsBought := map[string]int{}

	for _, data := range input {
		productsBought[data.Sku]++
	}

	cart, err := r.ShopUC.Checkout(ctx, productsBought)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
