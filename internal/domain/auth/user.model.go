package auth

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	IDRole    string     `json:"idRole" db:"id_role"`
	Username  string     `json:"username" db:"username"`
	Nama      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
}

// UserDTO digunakan untuk model join ke Role
type UserDTO struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	IDRole    string     `json:"idRole" db:"id_role"`
	Username  string     `json:"username" db:"username"`
	Nama      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
}

type StatusLogin string

const (
	SuccessLogin StatusLogin = "success"
	FailedLogin  StatusLogin = "failed"
)

// InputLogin is struct as login json body
type InputLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Response is represent respond login
func (r *InputLogin) Response(user User, role Role, accessToken string) ResponseLogin {
	return ResponseLogin{
		Token: ResponseLoginToken{
			AccessToken: accessToken,
		},
		User: ResponseLoginUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Nama:     user.Nama,
			Role:     role,
		},
	}
}

// ResponseLogin is result processing from login process
type ResponseLogin struct {
	Token ResponseLoginToken `json:"token"`
	User  ResponseLoginUser  `json:"user"`
}

// ResponseLoginUser deliver result of user entity
type ResponseLoginUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Nama     string    `json:"name" db:"name"`
	Role     Role      `json:"role"`
}

// ResponseLoginToken deliver result of user token
type ResponseLoginToken struct {
	AccessToken string
}

// NewUserLoginClaims digunakan untuk mengeset nilai dari JWT
func NewUserLoginClaims(user User, expiredIn int) jwt.MapClaims {
	claims := jwt.MapClaims{}
	claims["userId"] = user.ID
	claims["roleId"] = user.IDRole
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Duration(expiredIn) * time.Hour).Unix()

	return claims
}

func NewRoleLoginClaims(roleId string, expiredIn int) jwt.MapClaims {
	claims := jwt.MapClaims{}
	claims["roleId"] = roleId
	return claims
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

type UserHasRole struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	IdRole    string     `json:"idRole" db:"id_role"`
	Name      *string    `json:"name" db:"name"`
	IdUser    string     `json:"idUser" db:"id_user"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
}

type UserHasRoleDTO struct {
	IdRole string  `json:"idRole" db:"id_role"`
	Name   *string `json:"name" db:"name"`
	Utama  bool    `json:"utama" db:"utama"`
}

type UserHasRoleRequest struct {
	Id     *string `db:"id" json:"id"`
	IdRole string  `db:"id_role" json:"idRole"`
}

type Role struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
