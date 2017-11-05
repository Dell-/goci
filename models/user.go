package models

import (
	"time"
	"strings"
	"crypto/sha256"
	"fmt"
	"crypto/subtle"
	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	Id          int64     `xorm:"pk autoincr"`
	FullName    string    `xorm:"varchar(125) NOT NULL"`
	Username    string    `xorm:"varchar(125) NOT NULL UNIQUE"`
	Email       string    `xorm:"UNIQUE NOT NULL"`
	Password    string    `xorm:"varchar(200) NOT NULL"`
	Salt        string    `xorm:"VARCHAR(10)"`
	IsActive    bool      `xorm:"tinyint(1) NUT NULL"`
	Created     time.Time `xorm:"-"`
	CreatedUnix int64
	Updated     time.Time `xorm:"-"`
	UpdatedUnix int64
}

// GetUserByUsername returns user by given username.
func GetUserByUsername(username string) (*User) {
	if len(username) == 0 {
		return nil
	}
	user := &User{Username: strings.ToLower(username)}
	has, err := engine.Get(user)
	if err != nil {
		return nil
	}

	if !has {
		return nil
	}

	return user
}

// GetUserByUsername returns user by given id.
func getUserByID(e Engine, id int64) (*User) {
	user := new(User)
	has, err := e.Id(id).Get(user)
	if err != nil {
		return nil
	}

	if !has {
		return nil
	}

	return user
}

// GetUserByEmail returns the user object by given e-mail if exists.
func GetUserByEmail(email string) (*User) {
	if len(email) == 0 {
		return nil
	}

	email = strings.ToLower(email)
	// First try to find the user by  email
	user := &User{Email: email}
	has, err := engine.Get(user)
	if err != nil {
		return nil
	}

	if has {
		return user
	}

	return nil
}

// TODO : Change according this doc https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/09.5.html
// EncodePassword encodes password to safe format.
func (user *User) EncodePasswd() {
	newPassword := pbkdf2.Key([]byte(user.Password), []byte(user.Salt), 10000, 50, sha256.New)
	user.Password = fmt.Sprintf("%x", newPassword)
}

// ValidatePassword checks if given password matches the one belongs to the user.
func (user *User) ValidatePassword(password string) bool {
	newUser := &User{Password: password, Salt: user.Salt}
	newUser.EncodePasswd()
	return subtle.ConstantTimeCompare([]byte(user.Password), []byte(newUser.Password)) == 1
}
