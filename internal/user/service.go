package user

import (
	"errors"
	"time"

	"fyrnoapi/model"
	"fyrnoapi/pkg/database"
)

type UserService struct {
	DB *database.Database
}

func NewUserService(db *database.Database) *UserService {
	return &UserService{
		DB: db,
	}
}

// CreateUser creates a new user with the given name and phone number
func (s *UserService) CreateUser(name string, phone string) (*model.User, error) {
	user := &model.User{
		Name:      name,
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.DB.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByPhone fetches a user by phone number
func (s *UserService) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	if err := s.DB.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID fetches a user by ID
func (s *UserService) GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	if err := s.DB.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// IsPhoneRegistered checks if the phone number already exists
func (s *UserService) IsPhoneRegistered(phone string) (bool, error) {
	var count int64
	if err := s.DB.DB.Model(&model.User{}).Where("phone = ?", phone).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// UpdateUserName updates a user’s name
func (s *UserService) UpdateUserName(userID uint, newName string) error {
	result := s.DB.DB.Model(&model.User{}).Where("id = ?", userID).Update("name", newName)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no user found with given ID")
	}
	return nil
}
