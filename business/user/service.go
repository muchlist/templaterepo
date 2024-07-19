package user

import (
	"context"
	modelUser "templaterepo/models/user"
)

type UserService struct {
	storer   UserStorer
	notifier NotifSender
}

// perhatikan retrun dari NewUserService, alih alih meretrun langsung interface,
// kita justru mereturn tipe konkrit UserService
func NewUserService(store UserStorer, notifier NotifSender) *UserService {
	return &UserService{storer: store, notifier: notifier}
}

func (u *UserService) Login(ctx context.Context, username, password string) (modelUser.UserResp, error) {
	user, _ := u.storer.Get(ctx, username)
	// TODO : komparasi hashing password

	// TODO : validasi

	// TODO : memanggil database, rest api, ataupun penyimpanan internal dan memadukannya
	// semuanya terjadi di layer service

	// send notif
	u.notifier.SendNotification(username + " success login")

	return modelUser.UserResp{
		UID:         user.UID,
		Name:        user.Name,
		AccessToken: "dummy",
	}, nil
}
