package repository

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/require"
)

func (s *Suite) Test_repository_UpdateQuantity() {
	sku := "234234"
	var newQty uint = 3
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "products" SET "quantity"=$1 WHERE sku = $2`)).
		WithArgs(newQty, sku).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	ctxSku := context.WithValue(context.Background(), skuKey, sku)
	err := s.repository.UpdateQuantity(ctxSku, newQty)

	require.NoError(s.T(), err, err)
}