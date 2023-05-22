package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"gitlab.com/upn-belajar-go/configs"
	"golang.org/x/crypto/bcrypt"
)

// UserService is the service interface for User entities.
type UserService interface {
	Login(input InputLogin) (ResponseLogin, bool, error)
}

// ServiceImpl is the service implementation for User entities.
type UserServiceImpl struct {
	Config         *configs.Config
	UserRepository UserRepository
}

// ProvideUserServiceImpl is the provider for this service.
func ProvideUserServiceImpl(userRepository UserRepository, config *configs.Config) *UserServiceImpl {
	return &UserServiceImpl{
		Config:         config,
		UserRepository: userRepository,
	}
}

// Login is the service to process user signin
func (u *UserServiceImpl) Login(input InputLogin) (response ResponseLogin, exist bool, err error) {
	exist, err = u.UserRepository.ExistUserLoginByUsername(input.Username)
	if !exist {
		err = errors.New("Username tidak ditemukan")
		return
	}

	user, err := u.UserRepository.ResolveUserByUsername(input.Username)
	if err != nil {
		return
	}

	// Pengecekan Password bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		err = errors.New("Password Salah")
		return
	}

	role, err := u.UserRepository.ResolveRoleByID(user.IDRole)
	if err != nil {
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, NewUserLoginClaims(user, u.Config.Token.JWT.ExpiredInHour))
	token, err := claims.SignedString([]byte(u.Config.Token.JWT.AccessToken))
	if err != nil {
		fmt.Println(err)
		return
	}

	response = input.Response(user, role, string(token))
	return
}
