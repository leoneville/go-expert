package uow

import (
	"context"
	"database/sql"
)

type RepositoryFactory func(*sql.Tx) any

type UowInterface interface {
	Register(name string, fn RepositoryFactory)
	GetRepository(ctx context.Context, name string) (any, error)
	Do(ctx context.Context, fn func(uow UowInterface) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type Uow struct {
	DB           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(db *sql.DB) *Uow {
	return &Uow{
		DB:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fn RepositoryFactory) {
	u.Repositories[name] = fn
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}
