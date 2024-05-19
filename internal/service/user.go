package service

import (
	"errors"
	"log"
	"time"

	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/repository"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/encrypt"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(email string, password string) (string, error)
	FindAllUser() ([]entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id uuid.UUID) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
	tokenUseCase   token.TokenUseCase
	encryptTool    encrypt.EncryptTool
}

func NewUserService(userRepository repository.UserRepository, tokenUseCase token.TokenUseCase, encryptTool encrypt.EncryptTool) *userService {
	return &userService{
		userRepository: userRepository,
		tokenUseCase:   tokenUseCase,
		encryptTool:    encryptTool,
	}
}

func (s *userService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("email/password yang anda masukkan salah")
	}

	// Debugging: log hash kata sandi dari database
	log.Printf("Hashed password from database: %s", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Debugging: log kesalahan yang dikembalikan oleh bcrypt
		log.Printf("Password comparison error: %v", err)
		return "", errors.New("email/password yang anda masukkan salah")
	}

	// Lanjutkan dengan pembuatan token dan logika lainnya
	expiredTime := time.Now().Local().Add(5 * time.Minute)

	user.Alamat, _ = s.encryptTool.Decrypt(user.Alamat)
	user.NoHp, _ = s.encryptTool.Decrypt(user.NoHp)

	claims := token.JwtCustomClaims{
		ID:     user.ID.String(),
		Email:  user.Email,
		Role:   user.Role,
		Alamat: user.Alamat,
		NoHP:   user.NoHp,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Go-Commerce",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		return "", errors.New("ada kesalahan dari sistem")
	}
	return token, nil
}

func (s *userService) FindAllUser() ([]entity.User, error) {
	users, err := s.userRepository.FindAllUser()
	if err != nil {
		return nil, err
	}

	formattedUser := make([]entity.User, 0)
	for _, v := range users {
		v.Alamat, _ = s.encryptTool.Decrypt(v.Alamat)
		v.NoHp, _ = s.encryptTool.Decrypt(v.NoHp)
		formattedUser = append(formattedUser, v)
	}

	return formattedUser, nil
}

func (s *userService) CreateUser(user *entity.User) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	user.Alamat, _ = s.encryptTool.Encrypt(user.Alamat)
	user.NoHp, _ = s.encryptTool.Encrypt(user.NoHp)

	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	newUser.Alamat, _ = s.encryptTool.Decrypt(newUser.Alamat)
	newUser.NoHp, _ = s.encryptTool.Decrypt(newUser.NoHp)

	return newUser, nil
}

func (s *userService) UpdateUser(user *entity.User) (*entity.User, error) {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if user.Alamat != "" {
		user.Alamat, _ = s.encryptTool.Encrypt(user.Alamat)
	}
	if user.NoHp != "" {
		user.NoHp, _ = s.encryptTool.Encrypt(user.NoHp)
	}
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(id uuid.UUID) (bool, error) {
	user, err := s.userRepository.FindUserByID(id)
	if err != nil {
		return false, err
	}

	return s.userRepository.DeleteUser(user)
}
