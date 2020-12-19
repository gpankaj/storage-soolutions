package services

import (
	"github.com/gpankaj/storage-soolutions/application/domain"
	"github.com/gpankaj/storage-soolutions/application/utils"
)

func GetUser(userId int64) (*domain.User, *utils.StorageApplicationError) {
	return domain.GetUser(userId)
}
