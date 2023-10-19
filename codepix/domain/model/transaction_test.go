package model_test

import (
	"testing"

	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "12345-6"
	ownerName := "Rudolf HiOk"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "13579-0"
	ownerName = "Mariana Ba"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "mariba@testemail.com"
	pixKey, _ := model.NewPixKey(kind, accountDestination, key)

	require.NotEqual(t, account.ID, accountDestination.ID)

	amount := 10.50
	statusTransaction := "pending"
	transaction, err := model.NewTransaction(account, amount, pixKey, "Description test", "")
	
	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, "Description test")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, _ := model.NewPixKey(kind, account, key)

	_, err = model.NewTransaction(account, amount, pixKeySameAccount, "Description test", "")
	require.NotNil(t, err)

	_, err = model.NewTransaction(account, 0, pixKey, "Description test", "")
	require.NotNil(t, err)
}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "12345-6"
	ownerName := "Rudolf HiOk"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "13579-0"
	ownerName = "Mariana Ba"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)
	
	kind := "email"
	key := "mariba@testemail.com"
	pixKey, _ := model.NewPixKey(kind, accountDestination, key)

	amount := 10.50
	transaction, _ := model.NewTransaction(account, amount, pixKey, "Description test", "")

	transaction.Complete()
	require.Equal(t, transaction.Status, model.TransactionCompleted)

	transaction.Cancel("Error")
	require.Equal(t, transaction.Status, model.TransactionError)
	require.Equal(t, transaction.CancelDescription, "Error")
}
