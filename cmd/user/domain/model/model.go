package model

import (
	"context"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/common"
)

type User struct {
	PK           string
	Email        string
	Name         string
	AvatarSource string
}

type UserRepository interface {
	GetUserByPK(context.Context, common.PrimaryKey) (*User, error)
	CreateUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, common.PrimaryKey) (common.PrimaryKey, error)
}
