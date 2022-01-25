package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(conn)
	accout1, _ := store.Queries.CreateAccount(context.Background(),
		CreateAccountParams{Owner: "nothing", Balance: 100, Currency: "USD"})
	accout2, _ := store.Queries.CreateAccount(context.Background(),
		CreateAccountParams{Owner: "sub", Balance: 100, Currency: "USD"})
	fmt.Println("before : >>", accout1.Balance, accout2.Balance)
	errs := make(chan error)
	results := make(chan TransferTxResult)
	for i := 0; i < 5; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: accout1.ID,
				ToAccountID:   accout2.ID,
				Amount:        float64(10),
			})
			errs <- err
			results <- result
		}()
	}
	exist := make(map[int]bool)
	for i := 0; i < 5; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, accout1.ID, transfer.FromAccountID)
		require.Equal(t, accout2.ID, transfer.ToAccountID)
		require.Equal(t, float64(10), transfer.Amount)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)

		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		fmt.Println("after", fromAccount.Balance, toAccount.Balance)
		diff1 := accout1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - accout2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, int(diff1)%10 == 0)
		k := int(diff1 / 10)
		require.True(t, k >= 1 && k <= 5)
		require.NotContains(t, exist, k)
		exist[k] = true

	}
}
