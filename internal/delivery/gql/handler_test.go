package gql

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fajarabdillahfn/shoping-gql/internal/model"
	"github.com/fajarabdillahfn/shoping-gql/internal/usecase"
)

func Test_checkout(t *testing.T) {
	productSku := "120P90"
	uMock := &usecase.UseCaseMock{
		CheckoutFunc: func(ctx context.Context, productsBought map[string]int) (*model.Cart, error) {
			var cart *model.Cart

			switch productsBought[productSku] {
			case 1:
				cart = &model.Cart{
					Products: []*model.CheckoutProduct{
						{
							Sku:        "120P90",
							Name:       "Google Home",
							Price:      49.99,
							Quantity:   1,
							TotalPrice: 49.99,
						},
					},
					TotalPrice: 49.99,
				}
			case 11:
				return nil, errors.New("")
			case 3:
				cart = &model.Cart{
					Products: []*model.CheckoutProduct{
						{
							Sku:        "120P90",
							Name:       "Google Home",
							Price:      49.99,
							Quantity:   3,
							TotalPrice: 99.98,
						},
					},
					TotalPrice: 99.98,
				}
			}
			return cart, nil
		},
	}
	type fields struct {
		ShopUC usecase.UseCase
	}
	type args struct {
		ctx   context.Context
		input []*model.ProductInput
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *model.Cart
		wantErr  bool
	}{
		{
			name:   "normal",
			fields: fields{ShopUC: uMock},
			args: args{
				ctx: context.Background(),
				input: []*model.ProductInput{
					{Sku: productSku},
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "120P90",
						Name:       "Google Home",
						Price:      49.99,
						Quantity:   1,
						TotalPrice: 49.99,
					},
				},
				TotalPrice: 49.99,
			},
			wantErr: false,
		},
		{
			name:   "error",
			fields: fields{ShopUC: uMock},
			args: args{
				ctx: context.Background(),
				input: []*model.ProductInput{
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
					{Sku: productSku},
				},
			},
			wantData: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mutationResolver{
				&Resolver{
					ShopUC: tt.fields.ShopUC,
				},
			}
			gotData, err := m.Checkout(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Checkout() error = %v, wantErr %v, data %v", err, tt.wantErr, gotData)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.Checkout() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
