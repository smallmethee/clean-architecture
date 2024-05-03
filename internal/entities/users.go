package entities

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username string `gorm:"uniqueIndex;not null"`
		Password string `gorm:"not null"`
	}

	UsersRegisterReq struct {
		Username string `json:"username" db:"username" validate:"required,min=3,max=25"`
		Password string `json:"password" db:"password" validate:"required,min=6"`
	}

	UsersRegisterRes struct {
		Id       uint64 `json:"id" db:"id"`
		Username string `json:"username" db:"username"`
	}

	UserListDto struct {
		ID        uint      `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserLoginDto struct {
		Username string `json:"username" db:"username" validate:"required,min=3,max=25"`
		Password string `json:"password" db:"password" validate:"required,min=6"`
	}

	UserLoginResponseDto struct {
		Token string `json:"token"`
	}

	UserFindByIdDto struct {
		UserId uint `json:"id"`
	}

	UserMeResponseDto struct {
		ID        uint      `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

type UserUsecase interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
	List() ([]*UserListDto, error)
	Login(req *UserLoginDto) (*UserLoginResponseDto, error)
	Me(req *UserFindByIdDto) (*UserMeResponseDto, error)
}

type UsersRepository interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
	List() ([]*UserListDto, error)
	Login(req *UserLoginDto) (*User, error)
	FindById(req *UserFindByIdDto) (*User, error)
	ExistingUser(req *string) (*User, error)
}
