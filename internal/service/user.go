package service

import (
	"errors"

	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/repository"
)

type UserService interface {
	Login(email string, password string) (*entity.User, error)
	FindAllUser() ([]entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Login(email string, password string) (*entity.User, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return user, errors.New("email/password yang anda masukkan salah")
	}

	if user.Password != password {
		return user, errors.New("email/password yang anda masukkan salah")
	}

	return user, nil
}

func (s *userService) FindAllUser() ([]entity.User, error) {
	return s.userRepository.FindAllUser()
}