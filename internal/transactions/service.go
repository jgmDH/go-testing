package transactions

import "github.com/bootcamp-go/internal/domains"

type Service interface {
	GetAll() ([]*domains.Transaction, error)
	Store(codigo, moneda, emisor, receptor string, monto float64) (*domains.Transaction, error)
}

type service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return &service{rep: r}
}

func (s *service) GetAll() ([]*domains.Transaction, error) {
	return s.rep.GetAll()
}

func (s *service) Store(codigo, moneda, emisor, receptor string, monto float64) (*domains.Transaction, error) {
	lastId, err := s.rep.LastId()
	if err != nil {
		return nil, err
	}

	lastId++
	return s.rep.Store(lastId, codigo, moneda, emisor, receptor, monto)
}
