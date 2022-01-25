package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("%v %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64   `json:"from_account_id"`
	ToAccountID   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEmtry     Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        float64(arg.Amount),
		})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -float64(arg.Amount),
		})
		if err != nil {
			return err
		}
		result.ToEmtry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    float64(arg.Amount),
		})
		if err != nil {
			return err
		}

		result.FromAccount, err = q.AddAccountBalane(ctx, AddAccountBalaneParams{
			ID:      arg.FromAccountID,
			Ammount: -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToAccount, err = q.AddAccountBalane(ctx, AddAccountBalaneParams{
			ID:      arg.ToAccountID,
			Ammount: arg.Amount,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

func addMoney(ctx context.Context,q *Queries,accountID1 int64,amount1 float64,accountID2 int64,amount2 float64) (account1 Account, account2 Account,err error) {
	account1,err= q.AddAccountBalane(ctx,AddAccountBalaneParams{
		Ammount: amount1,
		ID: accountID1,
	})
	
}