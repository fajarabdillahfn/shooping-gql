package repository

import (
	"context"

	"github.com/fajarabdillahfn/shoping-gql/internal/model"
)

func (r *repository) UpdateQuantity(ctx context.Context, newQuantity int) error {
	sku := ctx.Value("sku")

	res := r.conn.WithContext(ctx).
	Model(&model.Product{}).
	Where("sku = ?", sku).
	Update("quantity", newQuantity)

	if res.Error != nil {
		return res.Error
	}

	return nil
}