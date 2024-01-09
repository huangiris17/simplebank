package db

import (
	"context"
	"testing"
	"time"

	"github.com/huangiris17/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account1 Accounts) Entries {
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account1 := CreateRandomAccount(t)
	CreateRandomEntry(t, account1)
}

func TestGetEntry(t *testing.T) {
	account1 := CreateRandomAccount(t)
	entry1 := CreateRandomEntry(t, account1)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestListEntry(t *testing.T) {
	account1 := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, account1)
	}
	arg := ListEntriesParams{
		AccountID: account1.ID,
		Limit:     5,
		Offset:    5,
	}

	entrys, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entrys, 5)

	for _, entry := range entrys {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, account1.ID)
	}
}
