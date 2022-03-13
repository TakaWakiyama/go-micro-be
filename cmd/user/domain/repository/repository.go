package repository

import (
	"context"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/common"
	"github.com/TakaWakiyama/forcusing-backend/cmd/user/domain/model"
)

type UserRepository interface {
	GetUserByPK(context.Context, common.PrimaryKey) (*model.User, error)
	CreateUser(context.Context, *model.User) (*model.User, error)
	UpdateUser(context.Context, *model.User) (*model.User, error)
	DeleteUser(context.Context, common.PrimaryKey) (common.PrimaryKey, error)
}
