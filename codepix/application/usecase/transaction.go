package usecase

import (
	"log"

	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository         model.PixKeyRepositoryInterface
}

func (usecase *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string, id string) (*model.Transaction, error) {
	account, err := usecase.PixRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := usecase.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description, id)
	if err != nil {
		return nil, err
	}

	err = usecase.TransactionRepository.Register(transaction)
	if transaction.ID == "" {
		return nil, err
	}

	return transaction, nil
}

func (usecase *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := usecase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (usecase *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := usecase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (usecase *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := usecase.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
