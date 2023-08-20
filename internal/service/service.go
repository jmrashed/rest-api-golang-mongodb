package service

import (
	"errors"

	"github.com/jmrashed/rest-api-golang-mongodb/internal/model"
	"github.com/jmrashed/rest-api-golang-mongodb/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) ListUsers() ([]model.User, error) {
	return s.userRepository.ListUsers()
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) UpdateUser(id string, user *model.User) (*model.User, error) {
	existingUser, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}
	updatedUser := &model.User{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
	}
	return s.userRepository.UpdateUser(updatedUser)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepository.DeleteUser(id)
}
