package models

import (
	"time"
	"strings"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	log "github.com/go-clog/clog"
	"github.com/go-xorm/xorm"
)

type User struct {
	ID          int64     `xorm:"pk autoincr"`
	FullName    string    `xorm:"varchar(125) NOT NULL"`
	Username    string    `xorm:"varchar(125) NOT NULL UNIQUE"`
	Email       string    `xorm:"UNIQUE NOT NULL"`
	Password    string    `xorm:"varchar(255) NOT NULL"`
	IsActive    bool      `xorm:"tinyint(1) NUT NULL"`
	Created     time.Time `xorm:"-"`
	CreatedUnix int64
	Updated     time.Time `xorm:"-"`
	UpdatedUnix int64
}

func (user *User) BeforeInsert() {
	user.CreatedUnix = time.Now().Unix()
}

func (user *User) BeforeUpdate() {
	user.UpdatedUnix = time.Now().Unix()
}

func (user *User) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_unix":
		user.Created = time.Unix(user.CreatedUnix, 0).Local()
	case "updated_unix":
		user.Updated = time.Unix(user.UpdatedUnix, 0).Local()
	}
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

func (user *User) HashPassword(password string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Warn("Fail to hash password: %v\n", err)
		return
	}

	user.Password = fmt.Sprintf("%s", string(bytes))
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// CreateUser creates record of a new user.
func CreateUser(user *User) (err error) {
	user.Email = strings.ToLower(user.Email)
	isExist, err := IsEmailUsed(user.Email)
	if err != nil {
		return err
	} else if isExist {
		return ErrUserAlreadyExist{user.Email}
	}

	user.HashPassword(user.Password)

	sess := engine.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(user); err != nil {
		return err
	}

	return sess.Commit()
}

// IsUserExist checks if given user email exist.
func IsUserExist(uid int64, email string) (bool, error) {
	if len(email) == 0 {
		return false, nil
	}
	return engine.Where("id != ?", uid).Get(User{Email: strings.ToLower(email)})
}

func isEmailUsed(engine Engine, email string) (bool, error) {
	if len(email) == 0 {
		return true, nil
	}

	// We need to check primary email of users as well.
	return engine.Where("email=?", email).Get(new(User))
}

// IsEmailUsed returns true if the email has been used.
func IsEmailUsed(email string) (bool, error) {
	return isEmailUsed(engine, email)
}
