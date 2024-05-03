package usecases

import (
	"fmt"
	"tder/internal/domain"
	"tder/internal/entities"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
	Hasher    domain.PasswordHasher
	Jwt       domain.JWTTokenService
}

// Me implements entities.UserUsecase.
func (u *usersUse) Me(req *entities.UserFindByIdDto) (*entities.UserMeResponseDto, error) {
	user, err := u.UsersRepo.FindById(req)
	if err != nil {
		return nil, err
	}

	newUser := &entities.UserMeResponseDto{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return newUser, nil
}

// Login implements entities.UserUsecase.
func (u *usersUse) Login(req *entities.UserLoginDto) (*entities.UserLoginResponseDto, error) {
	user, err := u.UsersRepo.Login(req)
	if err != nil {
		return nil, err
	}
	hash := u.Hasher.CheckPasswordHash(req.Password, user.Password)
	fmt.Println(hash)
	if hash {
		token, err := u.Jwt.GenerateToken(user.ID)
		fmt.Println(token)
		if err != nil {
			return nil, err
		}
		response := &entities.UserLoginResponseDto{
			Token: token,
		}
		return response, nil
	}
	return nil, fmt.Errorf("login failed")
}

func (u *usersUse) List() ([]*entities.UserListDto, error) {
	user, err := u.UsersRepo.List()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersUse) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {

	existingUser, err := u.UsersRepo.ExistingUser(&req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing username: %v", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	hashed, err := u.Hasher.HashPassword(req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = hashed

	user, err := u.UsersRepo.Register(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUsersUsecase(usersRepo entities.UsersRepository, hasher domain.PasswordHasher, jwt domain.JWTTokenService) entities.UserUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
		Hasher:    hasher,
		Jwt:       jwt,
	}
}
