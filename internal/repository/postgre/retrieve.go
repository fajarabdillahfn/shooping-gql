package repository

import (
	"context"

	"github.com/fajarabdillahfn/shoping-gql/internal/model"
)

type key string

const skuKey key = "sku"

func (r *repository) GetBySku(ctx context.Context) (*model.Product, error) {
	var product model.Product

	sku := ctx.Value(skuKey)

	res := r.conn.WithContext(ctx).First(&product, "sku = ?", sku)
	if res.Error != nil {
		return nil, res.Error
	}

	return &product, nil
}
