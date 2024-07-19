package user

import (
	"context"
	modelUser "templaterepo/models/user"
)

// UserStorer adalah interface yang mendefinisikan operasi yang dapat dilakukan terhadap database user.
// Interface ini Merupakan milik dari layer service dan dimaksudkan ditulis pada bagian layer service
// Meskipun kita tau persis implementasinya ada di business/user/repo.go, tetap layer service (core) hanya bergantung pada interface ini.
// Implementasi konkret dari antarmuka ini akan ditentukan oleh pengaturan dependensi di folder /app.
type UserStorer interface {
	Get(ctx context.Context, uid string) (modelUser.UserDTO, error)
	CreateOne(ctx context.Context, user *modelUser.UserEntity) error
}

// NotifSender adalah interface yang mendefinisikan operasi untuk mengirim notifikasi.
// Interface ini Merupakan milik dari layer service dan dimaksudkan ditulis pada bagian layer service
// Objek yang digunakan untuk mengirim notifikasi akan ditentukan oleh pengaturan dependensi di folder /app.
type NotifSender interface {
	SendNotification(message string) error
}
