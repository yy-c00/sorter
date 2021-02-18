package database

import (
	"github.com/yy-c00/sorter/model"
)

//NewUserFinder returns an implementation of model.UserFinder
func NewUserFinder() model.UserFinder {
	return access{}
}

type access struct{}

func (a access) FindUser(user *model.User) error {
	row := Connection().QueryRow("SELECT name, lastname FROM user WHERE user = ? OR iduser = ?", user.User, user.Id)

	err := row.Scan(&user.Name, &user.LastName)
	if err != nil {
		return err
	}

	err = Connection().QueryRow("SELECT idroot FROM root WHERE iduser = ?", user.Id).Scan(&user.Id)
	if err == nil {
		user.Root = true
		return nil
	}

	err = Connection().QueryRow("SELECT idcommon FROM common WHERE iduser = ?", user.Id).Scan(&user.Id)
	if err == nil {
		return nil
	}

	return nil
}