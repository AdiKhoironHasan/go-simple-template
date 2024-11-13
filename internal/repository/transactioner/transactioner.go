package transactioner

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type TxBeginner interface {
	Begin(ctx context.Context, opts ...*sql.TxOptions) (Transactioner, error)
}

type Transactioner interface {
	Commit() error
	Rollback() error

	// methods return repositories with transactioner
}

type txBeginner struct {
	db *gorm.DB
}

func NewTxBeginner(db *gorm.DB) TxBeginner {
	return &txBeginner{db: db}
}

func (t *txBeginner) Begin(ctx context.Context, opts ...*sql.TxOptions) (Transactioner, error) {
	tx := t.db.WithContext(ctx).Begin(opts...)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &transactioner{db: tx}, nil
}

type transactioner struct {
	db *gorm.DB
}

func (t *transactioner) Commit() error {
	tx := t.db.Commit()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (t *transactioner) Rollback() error {
	tx := t.db.Rollback()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
