package transactions

import (
	"github.com/bootcamp-go/internal/domains"
	"github.com/bootcamp-go/pkg/store"
)

type Repository interface {
	LastId() (int64, error)
	GetAll() ([]*domains.Transaction, error)
	Store(id int64, codigo, moneda, emisor, receptor string, monto float64) (*domains.Transaction, error)
}

// No lo usamos más var transactions []*domains.Transaction

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]*domains.Transaction, error) {
	var ts []*domains.Transaction
	if err := r.db.Read(&ts); err != nil {
		return nil, err
	}

	return ts, nil
}

func (r *repository) Store(id int64, codigo, moneda, emisor, receptor string, monto float64) (*domains.Transaction, error) {
	var ts []*domains.Transaction
	transaction := &domains.Transaction{
		Id:       id,
		Codigo:   codigo,
		Moneda:   moneda,
		Emisor:   emisor,
		Receptor: receptor,
		Monto:    monto,
	}

	err := r.db.Read(&ts) // Return err if json file is empty

	if err != nil {
		ts = append(ts, transaction)
		if err := r.db.Write(&ts); err != nil {
			return nil, err
		}

		return transaction, nil
	}

	ts = append(ts, transaction)
	if err := r.db.Write(&ts); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) LastId() (int64, error) {
	var ts []*domains.Transaction

	if err := r.db.Read(&ts); err != nil {
		return 0, nil
	}

	if len(ts) == 0 {
		return 0, nil
	}

	return ts[len(ts)-1].Id, nil
}
