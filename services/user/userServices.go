package servicesUser

import (
	"errors"
	"fmt"

	r "github.com/dancannon/gorethink"
	"github.com/nfrush/G2R2-Blog-Build/database/rethink"
	"github.com/nfrush/G2R2-Blog-Build/models/user"
	"golang.org/x/crypto/bcrypt"
)

var session = rethink.GetSession()

//CreateUser creates a new user
func CreateUser(u *modelUser.User) error {
	exists := FindOneUser(u.Username)
	fmt.Println(exists)
	if exists == nil {
		generatehash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(generatehash)
		if err := r.Table("users").Insert(u).Exec(session); err != nil {
			return err
		}
		return nil
	}
	err := errors.New("User already exists!")
	fmt.Println(err)
	return err
}

//FindAllUser finds all users
func FindAllUser() []*modelUser.User {
	res, err := r.Table("users").Run(session)
	fmt.Println(err)

	var users []*modelUser.User
	res.All(&users)
	res.Close()
	return users
}

//FindOneUser finds user by username
func FindOneUser(username string) *modelUser.User {
	res, err := r.Table("users").Filter(map[string]interface{}{"Name": &username}).Run(session)
	fmt.Println(err)

	var user *modelUser.User
	res.One(&user)
	res.Close()
	return user
}

//UpdateUser updates the information for a single user
func UpdateUser(user *modelUser.User) error {
	if err := r.Table("users").Filter(map[string]interface{}{"Name": &user.Username}).Update(&user).Exec(session); err != nil {
		return err
	}
	return nil
}

//DeleteUser deletes the selected user
func DeleteUser(username string) error {
	if err := r.Table("users").Filter(map[string]interface{}{"Name": &username}).Delete().Exec(session); err != nil {
		return err
	}
	return errors.New("The user has been deleted successful!")
}
