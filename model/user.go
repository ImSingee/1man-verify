package model

import (
	"github.com/ImSingee/1man-verify/config"
	"github.com/lithammer/shortuuid/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username     string `gorm:"type:varchar(64);uniqueIndex;not null" validate:"required"` // 用户名
	Email        string `gorm:"type:varchar(64);uniqueIndex;not null" validate:"required"`
	Password     string `gorm:"not null"`                                   // 密码（加密后）
	RealPassword string `gorm:"-:all" validate:"required_without=Password"` // 密码（明文）

	UUID string `gorm:"type:varchar(64);uniqueIndex;not null"`

	Role InternalRole `gorm:"not null;index;default:user" validate:"oneof=user admin superadmin"` // 内部角色
}

func (u *User) GenUUID() string {
	return shortuuid.NewWithNamespace("user")
}

// CheckPassword compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns nil on success, or an error on failure.
func (u *User) CheckPassword(with string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(with))
}

func (u *User) updatePassword() error {
	if u.RealPassword == "" {
		return nil
	}

	p, err := bcrypt.GenerateFromPassword([]byte(u.RealPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(p)
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = u.GenUUID()

	if err := config.Validator.Struct(u); err != nil {
		return err
	}

	return u.updatePassword()
}
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := config.Validator.Struct(u); err != nil {
		return err
	}

	return u.updatePassword()
}
