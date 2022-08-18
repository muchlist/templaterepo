package user

import (
	"context"
	"monorepo/models/duser"
)

type UserService struct {
	storer   UserStorer
	notifier NotifSender
}

func NewUserService(store UserStorer, notifier NotifSender) *UserService { // return tipe konkrit daripada interface
	return &UserService{storer: store, notifier: notifier}
}

func (u *UserService) Login(ctx context.Context, username, password string) (duser.UserResp, error) {
	user, _ := u.storer.Get(ctx, username)
	// komparasi hashing password

	// validasi

	// memanggil database, rest api, ataupun penyimpanan internal dan memadukannya
	// semuanya terjadi di layer service

	// send notif
	u.notifier.DummySendNotification(username + " success login")

	return duser.UserResp{
		UID:         user.UID,
		Name:        user.Name,
		AccessToken: "dummy",
	}, nil
}
