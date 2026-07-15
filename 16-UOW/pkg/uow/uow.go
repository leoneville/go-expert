package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(*sql.Tx) any

type UowInterface interface {
	Register(name string, fn RepositoryFactory)
	GetRepository(ctx context.Context, name string) (any, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
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

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	if u.Tx != nil {
		return errors.New("transaction already started")
	}
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx

	err = fn(u)
	if err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}

	err := u.Tx.Rollback()
	if err != nil {
		return err
	}

	u.Tx = nil
	return nil
}

func (u *Uow) CommitOrRollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to commit")
	}

	err := u.Tx.Commit()
	if err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}

	u.Tx = nil
	return nil
}

func (u *Uow) GetRepository(ctx context.Context, name string) (any, error) {
	if u.Tx == nil {
		tx, err := u.DB.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}

	repo := u.Repositories[name](u.Tx)
	return repo, nil
}
