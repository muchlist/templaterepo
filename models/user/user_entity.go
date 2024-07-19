package user

import "time"

type UserEntity struct {
	UID         string
	Name        string
	Password    string
	AccessToken string
	CreateAt    time.Time
}
