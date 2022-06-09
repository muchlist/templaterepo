package user

import (
	"context"
	"monorepo/internal/data/duser"
)

type UserRepo struct{}

func NewRepoUser() *UserRepo { // return tipe konkrit daripada interface
	return &UserRepo{}
}

func (u *UserRepo) Get(ctx context.Context, uid string) (duser.UserDTO, error) { // hindari return pointer
	return duser.UserDTO{
		UID:      uid,
		Name:     "dummy",
		Password: "dummy",
		CreateAt: 0,
	}, nil
}

func (u *UserRepo) CreateOne(ctx context.Context, user *duser.User) error {
	// create on database
	// return id
	user.UID = "new uid"
	return nil
}
