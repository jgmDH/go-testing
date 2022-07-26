package transactions

import (
	"errors"
	"testing"

	"github.com/bootcamp-go/internal/domains"
	"github.com/bootcamp-go/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	tsExpected := &domains.Transaction{
		Id:       13, // lasId incrementara el id cuando recibamos enviemos datos a la función store
		Codigo:   "asfk323",
		Moneda:   "dolar",
		Emisor:   "Miguel",
		Receptor: "Luciana",
		Monto:    899,
	}

	ts := []*domains.Transaction{
		{
			Id:       12,
			Codigo:   "asfk323",
			Moneda:   "dolar",
			Emisor:   "Miguel",
			Receptor: "Luciana",
			Monto:    899,
		},
	}

	mock := store.MockStore{Data: ts}
	repo := NewRepository(&mock)
	serv := NewService(repo)

	transaction, err := serv.Store(tsExpected.Codigo, tsExpected.Moneda, tsExpected.Emisor, tsExpected.Receptor, tsExpected.Monto)

	assert.Nil(t, err)
	assert.Equal(t, tsExpected, transaction)
}

func TestStoreError(t *testing.T) {
	var (
		codigo   string  = "abc233"
		moneda   string  = "dolar"
		emisor   string  = "Lali"
		receptor string  = "Christian"
		monto    float64 = 87.2
	)

	errorExpected := errors.New("error for Store")
	mock := store.MockStore{ErrorWrite: errorExpected}
	repo := NewRepository(&mock)
	serv := NewService(repo)
	transaction, err := serv.Store(codigo, moneda, emisor, receptor, monto)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, errorExpected, err)
}
