package service

import (
	"admin/app/admin/model"
	"admin/code"
	"admin/pkg/utils"
	"admin/proto/admin_account"
	"context"
	"gorm.io/gorm"
)

type AccountService struct {
	dao *model.AdminAccount
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{
		dao: model.NewAdminAccount(),
	}
}

func (a *AccountService) Login(ctx context.Context, params *admin_account.LoginReq) (*admin_account.LoginDataResp, error) {
	data, err := a.dao.FindAdminUserByUsername(ctx, params.Username)
	if err != nil {
		return nil, err
	}
	if !(utils.PasswordUtil{}).Matches(params.Password, data.Password) {
		return nil, code.NewCodeError(code.AdminAccountPasswordInvalid)
	}
	// TODO
	return nil, err
}
