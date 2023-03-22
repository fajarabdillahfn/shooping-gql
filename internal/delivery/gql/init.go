package gql

import uShop "github.com/fajarabdillahfn/shoping-gql/internal/usecase"

type Resolver struct{
	ShopUC uShop.UseCase
}

func NewShopGQL(shopUC uShop.UseCase) *Resolver {
	return &Resolver{
		ShopUC: shopUC,
	}
}

