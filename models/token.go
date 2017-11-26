package models

import (
	"time"
	"github.com/go-xorm/xorm"
	gouuid "github.com/satori/go.uuid"
	"github.com/Dell-/goci/pkg/tool"
)

const EXPIRED_HOUR int64 = 3600

// AccessToken represents a personal access token.
type AccessToken struct {
	ID          int64     `xorm:"pk autoincr"`
	UID         int64     `xorm:"INDEX"`
	Token       string    `xorm:"UNIQUE VARCHAR(40)"`
	Created     time.Time `xorm:"-"`
	CreatedUnix int64
	Expired     time.Time `xorm:"-"`
	ExpiredUnix int64
}

func (t *AccessToken) BeforeInsert() {
	t.CreatedUnix = time.Now().Unix()
	t.ExpiredUnix = time.Now().Unix() + EXPIRED_HOUR
}

func (t *AccessToken) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_unix":
		t.Created = time.Unix(t.CreatedUnix, 0).Local()
	case "expired_unix":
		t.Expired = time.Unix(t.ExpiredUnix, 0).Local()
	}
}

// NewAccessToken creates new access token.
func NewAccessToken(accessToken *AccessToken) error {
	accessToken.Token = tool.SHA1(gouuid.NewV4().String())
	_, err := engine.Insert(accessToken)
	return err
}

// GetAccessTokenBySHA returns access token by given sha1.
func GetAccessTokenBySHA(sha string) (*AccessToken) {
	if sha == "" {
		return nil
	}
	token := &AccessToken{Token: sha}
	has, err := engine.Get(token)
	if err != nil {
		return nil
	} else if !has {
		return nil
	}
	return token
}

// DeleteAccessTokenOfUserByID deletes access token by given ID.
func DeleteAccessTokenOfUserByID(userID int64) error {
	_, err := engine.Delete(&AccessToken{
		UID: userID,
	})
	return err
}

// DeleteAccessTokenOfUser deletes access token by given token.
func DeleteAccessTokenOfUser(token string) error {
	_, err := engine.Delete(&AccessToken{
		Token: token,
	})
	return err
}
