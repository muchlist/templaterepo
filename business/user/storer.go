package user

import (
	"context"
	"monorepo/models/duser"
)

type UserStorer interface {
	Get(ctx context.Context, uid string) (duser.UserDTO, error)
	CreateOne(ctx context.Context, user *duser.User) error
}

// User dapat memakai notifikasi yang notabene adalah service lain
// keduanya terhubungkan oleh interface
type NotifSender interface {
	DummySendNotification(message string) error
}
