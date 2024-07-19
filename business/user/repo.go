package user

import (
	"context"
	modelUser "templaterepo/models/user"
)

type UserRepo struct{}

// perhatikan retrun dari NewRepoUser, alih alih meretrun langsung interface,
// kita justru mereturn tipe konkrit UserRepo
func NewRepoUser() *UserRepo {
	return &UserRepo{}
}

// Get mengembalikan data dan error, disini kita menghindari return pointer
func (u *UserRepo) Get(ctx context.Context, uid string) (modelUser.UserDTO, error) {
	return modelUser.UserDTO{
		UID:      uid,
		Name:     "dummy",
		Password: "dummy",
		CreateAt: 0,
	}, nil
}

func (u *UserRepo) CreateOne(ctx context.Context, user *modelUser.UserEntity) error {
	// create on database
	// return id
	user.UID = "new uid"
	return nil
}
