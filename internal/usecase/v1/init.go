package v1

import (
	rShop "github.com/fajarabdillahfn/shoping-gql/internal/repository"
	uShop "github.com/fajarabdillahfn/shoping-gql/internal/usecase"
)

type useCase struct {
	ShopRepo rShop.Repository
}

func NewShopUseCase(shopRepo rShop.Repository) uShop.UseCase {
	return &useCase{
		ShopRepo: shopRepo,
	}
}
