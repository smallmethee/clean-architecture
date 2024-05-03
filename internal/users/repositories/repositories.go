package repositories

import (
	"fmt"
	"tder/internal/entities"

	"gorm.io/gorm"
)

type usersRepo struct {
	Db *gorm.DB
}

// ExistingUser implements entities.UsersRepository.
func (r *usersRepo) ExistingUser(req *string) (*entities.User, error) {
	var existingUser *entities.User
	if err := r.Db.Where("username = ?", req).First(&existingUser).Error; err != nil {
		return nil, err
	}
	return existingUser, nil
}

// FindById implements entities.UsersRepository.
func (r *usersRepo) FindById(req *entities.UserFindByIdDto) (*entities.User, error) {
	user := &entities.User{}
	if err := r.Db.Model(&entities.User{}).Where("id = ?", req.UserId).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Login implements entities.UsersRepository.
func (r *usersRepo) Login(req *entities.UserLoginDto) (*entities.User, error) {
	user := &entities.User{}
	if err := r.Db.Model(&entities.User{}).Where("username = ?", req.Username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *usersRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {

	user := &entities.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := r.Db.Create(user).Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	newUser := &entities.UsersRegisterRes{
		Id:       uint64(user.ID),
		Username: user.Username,
	}
	return newUser, nil
}

func (r *usersRepo) List() ([]*entities.UserListDto, error) {
	var users []*entities.User
	if err := r.Db.Model(&entities.User{}).Find(&users).Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var userList []*entities.UserListDto
	for _, user := range users {
		userDTO := &entities.UserListDto{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		userList = append(userList, userDTO)
	}

	return userList, nil
}

func NewUsersRepository(db *gorm.DB) entities.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}
