package user

import (
	"context"
	"monorepo/internal/data/duser"
)

type UserStorer interface {
	Get(ctx context.Context, uid string) (duser.UserDTO, error)
	CreateOne(ctx context.Context, user *duser.User) error
}

type NotifSender interface {
	DummySendNotification(message string) error
}
