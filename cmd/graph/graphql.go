package graph

import (
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql"
	
	cDB "github.com/fajarabdillahfn/shoping-gql/cmd/helper/db/postgre"

	gqlShop "github.com/fajarabdillahfn/shoping-gql/internal/delivery/gql"
	rShop "github.com/fajarabdillahfn/shoping-gql/internal/repository"
	uShop "github.com/fajarabdillahfn/shoping-gql/internal/usecase"

	pgShopRepo "github.com/fajarabdillahfn/shoping-gql/internal/repository/postgre"
	uShopV1 "github.com/fajarabdillahfn/shoping-gql/internal/usecase/v1"
)

var (
	pgShop       *gorm.DB
	shopRepo     rShop.Repository
	shopUseCase  uShop.UseCase
	ShopResolver *gqlShop.Resolver
)

func Initialize() graphql.ExecutableSchema {
	pgShop = cDB.OpenDB()

	shopRepo = pgShopRepo.NewShopRepo(pgShop)
	shopUseCase = uShopV1.NewShopUseCase(shopRepo)
	ShopResolver = gqlShop.NewShopGQL(shopUseCase)

	return gqlShop.NewExecutableSchema(gqlShop.Config{Resolvers: ShopResolver})
}
