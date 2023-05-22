package auth

import (
	"database/sql"

	"github.com/gofrs/uuid"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/logger"
)

var (
	userQuery = struct {
		Select,
		SelectLogin,
		ExistLogin,
		SelectRole string
	}{
		Select:      ` SELECT id, name, username, email, password, id_role, created_at, updated_at, is_deleted FROM public.users `,
		SelectLogin: ` SELECT id, name, username, email, password, id_role, created_at, updated_at, is_deleted FROM public.users `,
		ExistLogin:  `SELECT COUNT(u.id) > 0 FROM users u `,
		SelectRole:  `select id, name from role `,
	}
)

// UserRepositoryPostgreSQL digunakan untuk Repository User
type UserRepositoryPostgreSQL struct {
	DB     *infras.PostgresqlConn
	Config *configs.Config
}

// ProvideUserRepositoryPostgreSQL is the provider for this repository.
func ProvideUserRepositoryPostgreSQL(db *infras.PostgresqlConn, conf *configs.Config) *UserRepositoryPostgreSQL {
	return &UserRepositoryPostgreSQL{
		DB:     db,
		Config: conf,
	}
}

type UserRepository interface {
	ResolveUserByUsername(username string) (User, error)
	ExistUserLoginByUsername(username string) (exist bool, err error)
	ResolveRoleByID(id string) (Role, error)
}

// ExistByUsername is function to check that username exist or not
func (u *UserRepositoryPostgreSQL) ExistUserLoginByUsername(username string) (exist bool, err error) {
	err = u.DB.Read.Get(&exist, userQuery.ExistLogin+" WHERE (u.username = $1) AND coalesce(u.is_deleted, false) = false", username)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return exist, err
}

// ResolveUserByUsername is function resolving user data by username and role id
func (u *UserRepositoryPostgreSQL) ResolveUserByUsername(username string) (User, error) {
	var user User
	err := u.DB.Read.Get(&user, userQuery.SelectLogin+" where username=$1", username)
	if err != nil {
		logger.ErrorWithStack(err)
		return User{}, err
	}

	return user, nil
}

// ResolveUserByID is function resolving user data by email
func (u *UserRepositoryPostgreSQL) ResolveUserByID(id uuid.UUID) (User, error) {
	var user User
	err := u.DB.Read.Get(&user, userQuery.Select+" WHERE id = $1 AND coalesce(u.is_deleted, false) = false", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, err
		}
		logger.ErrorWithStack(err)
		return User{}, err
	}

	return user, nil
}

func (r *UserRepositoryPostgreSQL) ResolveRoleByID(id string) (Role, error) {
	var role Role
	err := r.DB.Read.Get(&role, userQuery.SelectRole+" where id=$1", id)
	if err != nil {
		logger.ErrorWithStack(err)
		return Role{}, err
	}
	return role, nil
}
