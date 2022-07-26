package store

import "github.com/bootcamp-go/internal/domains"

type MockStore struct {
	ReadInvoked bool
	Data        []*domains.Transaction
	ErrorRead   error
	ErrorWrite  error
}

func (fs *MockStore) Read(data interface{}) error {
	if fs.ErrorRead != nil {
		return fs.ErrorRead
	}
	fs.ReadInvoked = true
	a := data.(*[]*domains.Transaction)
	*a = fs.Data
	return nil
}

func (fs *MockStore) Write(data interface{}) error {
	if fs.ErrorWrite != nil {
		return fs.ErrorWrite
	}
	a := data.(*[]*domains.Transaction)
	*a = fs.Data
	return nil
}
