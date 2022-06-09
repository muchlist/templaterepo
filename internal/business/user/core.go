package user

import (
	"context"
	"monorepo/internal/data/duser"
)

type UserCore struct {
	storer   UserStorer
	notifier NotifSender
}

func NewUserCore(store UserStorer, notifier NotifSender) *UserCore { // return tipe konkrit daripada interface
	return &UserCore{storer: store, notifier: notifier}
}

func (u *UserCore) Login(ctx context.Context, username, password string) (duser.UserResp, error) {
	user, _ := u.storer.Get(ctx, username)
	// komparasi hashing password
	// valid

	// send notif
	u.notifier.DummySendNotification(username + " success login")

	return duser.UserResp{
		UID:         user.UID,
		Name:        user.Name,
		AccessToken: "dummy",
	}, nil
}
