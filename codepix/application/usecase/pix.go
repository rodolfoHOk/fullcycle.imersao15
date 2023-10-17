package usecase

import (
	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (usecase *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := usecase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	pixKey, err = usecase.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (usecase *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := usecase.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
