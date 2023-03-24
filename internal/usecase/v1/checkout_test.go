package v1

import (
	"context"
	"reflect"
	"testing"

	"github.com/fajarabdillahfn/shoping-gql/internal/model"
	rShop "github.com/fajarabdillahfn/shoping-gql/internal/repository"
)

func Test_checkout_GoogleHome(t *testing.T) {
	rMock := &rShop.RepositoryMock{
		GetBySkuFunc: func(ctx context.Context) (*model.Product, error) {
			return &model.Product{
				Sku:      "120P90",
				Name:     "Google Home",
				Price:    49.99,
				Quantity: 10,
			}, nil
		},
		UpdateQuantityFunc: func(ctx context.Context, newQuantity uint) error {
			return nil
		},
	}
	type fields struct {
		ShopRepo rShop.Repository
	}
	type args struct {
		ctx            context.Context
		productsBought map[string]int
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
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"120P90": 1,
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
			name:   "insufficient product",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"120P90": 11,
				},
			},
			wantData: nil,
			wantErr:  true,
		},
		{
			name:   "promotion 2",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"120P90": 3,
				},
			},
			wantData: &model.Cart{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				ShopRepo: tt.fields.ShopRepo,
			}
			gotData, err := u.Checkout(tt.args.ctx, tt.args.productsBought)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.Checkout() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_checkout_AlexaSpeaker(t *testing.T) {
	rMock := &rShop.RepositoryMock{
		GetBySkuFunc: func(ctx context.Context) (*model.Product, error) {
			return &model.Product{
				Sku:      "A304SD",
				Name:     "Alexa Speaker",
				Price:    109.5,
				Quantity: 10,
			}, nil
		},
		UpdateQuantityFunc: func(ctx context.Context, newQuantity uint) error {
			return nil
		},
	}
	type fields struct {
		ShopRepo rShop.Repository
	}
	type args struct {
		ctx            context.Context
		productsBought map[string]int
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
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"A304SD": 1,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "A304SD",
						Name:       "Alexa Speaker",
						Price:      109.5,
						Quantity:   1,
						TotalPrice: 109.5,
					},
				},
				TotalPrice: 109.5,
			},
			wantErr: false,
		},
		{
			name:   "insufficient product",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"A304SD": 11,
				},
			},
			wantData: nil,
			wantErr:  true,
		},
		{
			name:   "promotion 2",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"A304SD": 4,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "A304SD",
						Name:       "Alexa Speaker",
						Price:      109.5,
						Quantity:   4,
						TotalPrice: 394.2,
					},
				},
				TotalPrice: 394.2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				ShopRepo: tt.fields.ShopRepo,
			}
			gotData, err := u.Checkout(tt.args.ctx, tt.args.productsBought)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.Checkout() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_checkout_RaspberryPiB(t *testing.T) {
	rMock := &rShop.RepositoryMock{
		GetBySkuFunc: func(ctx context.Context) (*model.Product, error) {
			return &model.Product{
				Sku:      "234234",
				Name:     "Raspberry Pi B",
				Price:    30.0,
				Quantity: 2,
			}, nil
		},
		UpdateQuantityFunc: func(ctx context.Context, newQuantity uint) error {
			return nil
		},
	}
	type fields struct {
		ShopRepo rShop.Repository
	}
	type args struct {
		ctx            context.Context
		productsBought map[string]int
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
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"234234": 1,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "234234",
						Name:       "Raspberry Pi B",
						Price:      30.0,
						Quantity:   1,
						TotalPrice: 30,
					},
				},
				TotalPrice: 30,
			},
			wantErr: false,
		},
		{
			name:   "insufficient product",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"234234": 3,
				},
			},
			wantData: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				ShopRepo: tt.fields.ShopRepo,
			}
			gotData, err := u.Checkout(tt.args.ctx, tt.args.productsBought)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.Checkout() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_checkout_MacBookPro(t *testing.T) {

	rMock := &rShop.RepositoryMock{
		GetBySkuFunc: func(ctx context.Context) (*model.Product, error) {
			var product *model.Product
			sku := ctx.Value("sku")

			switch sku {
			case "43N23P":
				product = &model.Product{
					Sku:      "43N23P",
					Name:     "Macbook Pro",
					Price:    5399.99,
					Quantity: 5,
				}
			case "234234":
				product = &model.Product{
					Sku:      "234234",
					Name:     "Raspberry Pi B",
					Price:    30.0,
					Quantity: 2,
				}
			}
			return product, nil
		},
		UpdateQuantityFunc: func(ctx context.Context, newQuantity uint) error {
			return nil
		},
	}
	type fields struct {
		ShopRepo rShop.Repository
	}
	type args struct {
		ctx            context.Context
		productsBought map[string]int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *model.Cart
		wantErr  bool
	}{
		{
			name:   "promotion 1 - buy macbook only",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"43N23P": 1,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "43N23P",
						Name:       "Macbook Pro",
						Price:      5399.99,
						Quantity:   1,
						TotalPrice: 5399.99,
					},
					{
						Sku:        "234234",
						Name:       "Raspberry Pi B",
						Price:      30.0,
						Quantity:   1,
						TotalPrice: 0,
					},
				},
				TotalPrice: 5399.99,
			},
			wantErr: false,
		},
		{
			name:   "insufficient product",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"234234": 6,
				},
			},
			wantData: nil,
			wantErr:  true,
		},
		{
			name:   "promotion 1 - buy macbook and raspi",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"43N23P": 1,
					"234234": 1,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "43N23P",
						Name:       "Macbook Pro",
						Price:      5399.99,
						Quantity:   1,
						TotalPrice: 5399.99,
					},
					{
						Sku:        "234234",
						Name:       "Raspberry Pi B",
						Price:      30.0,
						Quantity:   1,
						TotalPrice: 0,
					},
				},
				TotalPrice: 5399.99,
			},
			wantErr: false,
		},
		{
			name:   "promotion 1 - quantity raspi < quantity macbook",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"43N23P": 3,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "43N23P",
						Name:       "Macbook Pro",
						Price:      5399.99,
						Quantity:   3,
						TotalPrice: 16199.97,
					},
					{
						Sku:        "234234",
						Name:       "Raspberry Pi B",
						Price:      30.0,
						Quantity:   2,
						TotalPrice: 0,
					},
				},
				TotalPrice: 16199.97,
			},
			wantErr: false,
		},
		{
			name:   "promotion 1 - quantity raspi > quantity macbook",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"43N23P": 1,
					"234234": 2,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "43N23P",
						Name:       "Macbook Pro",
						Price:      5399.99,
						Quantity:   1,
						TotalPrice: 5399.99,
					},
					{
						Sku:        "234234",
						Name:       "Raspberry Pi B",
						Price:      30.0,
						Quantity:   2,
						TotalPrice: 30.0,
					},
				},
				TotalPrice: 5429.99,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				ShopRepo: tt.fields.ShopRepo,
			}
			gotData, err := u.Checkout(tt.args.ctx, tt.args.productsBought)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.Checkout() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_checkout_MultipleProduct(t *testing.T) {

	rMock := &rShop.RepositoryMock{
		GetBySkuFunc: func(ctx context.Context) (*model.Product, error) {
			var product *model.Product
			sku := ctx.Value("sku")

			switch sku {
			case "120P90":
				product = &model.Product{
					Sku:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Quantity: 10,
				}
			case "A304SD":
				product = &model.Product{
					Sku:      "A304SD",
					Name:     "Alexa Speaker",
					Price:    109.5,
					Quantity: 10,
				}
			case "43N23P":
				product = &model.Product{
					Sku:      "43N23P",
					Name:     "Macbook Pro",
					Price:    5399.99,
					Quantity: 5,
				}
			case "234234":
				product = &model.Product{
					Sku:      "234234",
					Name:     "Raspberry Pi B",
					Price:    30.0,
					Quantity: 2,
				}
			}
			return product, nil
		},
		UpdateQuantityFunc: func(ctx context.Context, newQuantity uint) error {
			return nil
		},
	}
	type fields struct {
		ShopRepo rShop.Repository
	}
	type args struct {
		ctx            context.Context
		productsBought map[string]int
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
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"120P90": 1,
					"A304SD": 1,
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
					{
						Sku:        "A304SD",
						Name:       "Alexa Speaker",
						Price:      109.5,
						Quantity:   1,
						TotalPrice: 109.5,
					},
				},
				TotalPrice: 159.49,
			},
			wantErr: false,
		},
		{
			name:   "promotion 1",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"120P90": 1,
					"43N23P": 1,
					"234234": 1,
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
					{
						Sku:        "43N23P",
						Name:       "Macbook Pro",
						Price:      5399.99,
						Quantity:   1,
						TotalPrice: 5399.99,
					},
					{
						Sku:        "234234",
						Name:       "Raspberry Pi B",
						Price:      30.0,
						Quantity:   1,
						TotalPrice: 0,
					},
				},
				TotalPrice: 5449.98,
			},
			wantErr: false,
		},
		{
			name:   "promotion 2",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"120P90": 3,
					"234234": 1,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "120P90",
						Name:       "Google Home",
						Price:      49.99,
						Quantity:   3,
						TotalPrice: 99.98,
					},
					{
						Sku:        "234234",
						Name:       "Raspberry Pi B",
						Price:      30.0,
						Quantity:   1,
						TotalPrice: 30.0,
					},
				},
				TotalPrice: 129.98000000000002,
			},
			wantErr: false,
		},
		{
			name:   "promotion 3",
			fields: fields{ShopRepo: rMock},
			args: args{
				ctx: context.Background(),
				productsBought: map[string]int{
					"A304SD": 4,
					"234234": 1,
				},
			},
			wantData: &model.Cart{
				Products: []*model.CheckoutProduct{
					{
						Sku:        "A304SD",
						Name:       "Alexa Speaker",
						Price:      109.5,
						Quantity:   4,
						TotalPrice: 394.2,
					},
					{
						Sku:        "234234",
						Name:       "Raspberry Pi B",
						Price:      30.0,
						Quantity:   1,
						TotalPrice: 30.0,
					},
				},
				TotalPrice: 424.2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				ShopRepo: tt.fields.ShopRepo,
			}
			gotData, err := u.Checkout(tt.args.ctx, tt.args.productsBought)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.Checkout() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
