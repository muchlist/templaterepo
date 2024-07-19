package user

import "time"

/*
DTO adalah data class yang digunakan untuk berkomunikasi dengan service lain
DTO harus menyediakan fungsi atau method untuk merubah ke model konkrit (dalam hal ini User)
*/

type UserDTO struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Password string `json:"password"`
	CreateAt int64  `json:"create_at"` // int64 -> time.Time
}

func (ud UserDTO) ToUser() UserEntity {
	return UserEntity{
		UID:      ud.UID,
		Name:     ud.Name,
		Password: ud.Password,
		CreateAt: time.Unix(ud.CreateAt, 0),
	}
}

type UserResp struct {
	UID         string `json:"uid"`
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}

func FromUserToResponse(user UserEntity) UserResp {
	return UserResp{
		UID:         user.UID,
		Name:        user.Name,
		AccessToken: user.AccessToken,
	}
}
