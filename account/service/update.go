package service

import (
	"github.com/kyhsa93/gin-rest-example/account/dto"
)

// Update update account by accountID
func (service *Service) Update(accountID string, data *dto.Account) {
	oldData := service.ReadAccountByID(accountID)
	if oldData == nil {
		return
	}
	service.repository.Save(data, accountID)
}
