package infrastructure

import (
	"context"

	"github.com/TakaWakiyama/forcusing-backend/cmd/user/common"
	model "github.com/TakaWakiyama/forcusing-backend/cmd/user/domain/repository"
)

type userRepository struct{}

func NewUserRepository() model.UserRepository {
	return userRepository{}
}

// トランザクションどこで貼るか問題
// 参考: https://qiita.com/miya-masa/items/316256924a1f0d7374bb

func (userRepository) GetUserByPK(ctx context.Context, pk common.PrimaryKey) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (userRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (userRepository) DeleteUser(ctx context.Context, user common.PrimaryKey) (common.PrimaryKey, error) {
	panic("not implemented") // TODO: Implement
}
