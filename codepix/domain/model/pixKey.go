package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
	Base							 `valid:"required"`
	King 			string	 `json:"kind" valid:"notnull"`
	Key 			string 	 `json:"key" valid:"notnull"`
	AccountID string	 `json:"account_id" valid:"notnull"`
	Account 	*Account `valid:"-"`
	Status 		string	 `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_,err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}

	if pixKey.King != "email" && pixKey.King != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid status")
	}

	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		King: kind,
		Key: key,
		Account: account,
		Status: "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
