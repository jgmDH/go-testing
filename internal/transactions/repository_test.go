package transactions

import (
	"errors"
	"testing"

	"github.com/bootcamp-go/internal/domains"
	"github.com/bootcamp-go/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ts := []*domains.Transaction{
		{
			Id:       1,
			Codigo:   "abc23",
			Moneda:   "peso",
			Emisor:   "Juan",
			Receptor: "María",
			Monto:    100,
		},
		{
			Id:       2,
			Codigo:   "abc24",
			Moneda:   "dolar",
			Emisor:   "Sol",
			Receptor: "María",
			Monto:    100,
		},
	}

	mock := &store.MockStore{Data: ts}
	repo := NewRepository(mock)
	transactions, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, ts, transactions)
}

func TestGetAllError(t *testing.T) {
	expectedError := errors.New("error for GetAll")

	mock := store.MockStore{ErrorRead: expectedError}
	repo := NewRepository(&mock)
	ts, err := repo.GetAll()

	assert.Nil(t, ts) // Transactions deber ser nil
	assert.Equal(t, expectedError, err)
}
