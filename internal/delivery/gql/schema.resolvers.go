package gql

import (
	"context"

	model "github.com/fajarabdillahfn/shoping-gql/internal/model"
)

// Checkout is the resolver for the checkout field.
func (r *mutationResolver) Checkout(ctx context.Context, input []*model.ProductInput) (*model.Cart, error) {
	return &model.Cart{
		Products: []*model.Product{
			{Sku: "Hello", Name: "World", Price: 12, Quantity: 23},
		},
	}, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
