package duser

import "time"

type User struct {
	UID         string
	Name        string
	Password    string
	AccessToken string
	CreateAt    time.Time
}
