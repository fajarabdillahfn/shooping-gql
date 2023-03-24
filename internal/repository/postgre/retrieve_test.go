package repository

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fajarabdillahfn/shoping-gql/internal/model"
	"github.com/go-test/deep"

	"github.com/stretchr/testify/require"
)

func (s *Suite) Test_repository_GetBySku() {
	sku := "234234"
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE sku = $1 ORDER BY "products"."sku" LIMIT 1`)).
		WithArgs(sku).
		WillReturnRows(sqlmock.NewRows([]string{"sku", "name", "price", "quantity"}).
			AddRow("234234", "Raspberry Pi B", 30, 2))

	ctxSku := context.WithValue(context.Background(), "sku", sku)
	res, err := s.repository.GetBySku(ctxSku)

	expectedData := model.Product{
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30,
		Quantity: 2,
	}

	require.NoError(s.T(), err, err)
	require.Nil(s.T(), deep.Equal(&expectedData, res))
}

func (s *Suite) Test_repository_GetBySku_NoData() {
	sku := "abcbac"
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE sku = $1 ORDER BY "products"."sku" LIMIT 1`)).
		WithArgs(sku).
		WillReturnRows(sqlmock.NewRows([]string{"sku", "name", "price", "quantity"}))

	ctxSku := context.WithValue(context.Background(), "sku", sku)
	res, err := s.repository.GetBySku(ctxSku)

	require.Error(s.T(), err)
	require.Nil(s.T(), res)
}
