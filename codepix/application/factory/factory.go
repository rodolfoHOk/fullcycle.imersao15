package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/application/usecase"
	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{
		Db: database,
	}
	transactionRepository := repository.TransactionRepositoryDb{
		Db: database,
	}
	transactionUseCase := usecase.TransactionUseCase{
		PixRepository: &pixRepository,
		TransactionRepository: &transactionRepository,
	}
	return transactionUseCase
}
