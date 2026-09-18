package database

import (
	"errors"
	"url_short/util"
)

type User struct {
	ID       int    `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var userList []User

func (u *User) CheckUser() bool {
	for _, user := range userList {
		if u.Email == user.Email {
			return true
		}
	}
	return false
}

func (u *User) CreateUser() error {
	if len(userList) == 0 {
		u.ID = 1
	} else {
		u.ID = userList[len(userList)-1].ID + 1
	}

	hashPass := util.Hash(u.Password)
	if hashPass == nil {
		return errors.New("Couldn't generate hash pass")
	}
	u.Password = *hashPass
	userList = append(userList, *u)
	return nil
}

func (u *User) VerifyUser() (*User, error) {
	pass := u.Password
	hashPass := util.Hash(pass)

	if hashPass == nil {
		return nil, errors.New("Generating Hash Failed at VerifyUser")
	}

	for _, usr := range userList {
		if u.Email == usr.Email {
			if *hashPass == usr.Password {
				return &usr, nil
			}
		}
	}
	return nil, nil
}

func GetUserList() []User {
	return userList
}

func (u *User) GetUserById() bool {
	for _, user := range userList {
		if u.ID == user.ID {
			u = &user
			return true
		}
	}
	return false
}
