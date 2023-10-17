package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transaction struct {
	Base              				 `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyIdTo        string   `gorm:"column:pix_key_id_to;type:uuid;" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"-"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

type Transactions struct {
	Transaction []Transaction
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return err
	}

	if t.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("invalid status for the transaction")
	}

	if t.PixKeyTo.AccountID == t.AccountFromID {
		return errors.New("the source and destination account cannot be the same")
	}
	
	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	transaction.AccountFromID = accountFrom.ID
	transaction.PixKeyIdTo = pixKeyTo.ID

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (transaction *Transaction) Complete() error {
	transaction.Status = TransactionCompleted
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Cancel(description string) error {
	transaction.Status = TransactionError
	transaction.CancelDescription = description
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err
}
