package usecase

import (
	"context"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/common"
	model "github.com/TakaWakiyama/forcusing-backend/cmd/user/domain"
)

type UserUsecase interface {
	GetUserByPK(context.Context, common.PrimaryKey) (*model.User, error)
	CreateUser(context.Context, *model.User) (*model.User, error)
	PartialUpdateUser(context.Context, *common.UserInput) (*model.User, error)
	DeleteUser(context.Context, common.PrimaryKey) (*model.User, error)
}
