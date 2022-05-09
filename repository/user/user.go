package user

import (
	"context"
	"database/sql"

	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/entity"
	"github.com/t3be8/altacommerce/utils"
)

func New(db *sql.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

type UserRepo struct {
	Db *sql.DB
}

// Check users islogin with payload
func (ur *UserRepo) IsLogin(ctx context.Context, email, password string) (entity.User, bool, error) {
	var u entity.User
	var pwd string

	state := "SELECT * FROM users WHERE email = ?"
	err := ur.Db.QueryRowContext(ctx, state, email).Scan(
		&u.ID, &u.Email, &pwd,
	)

	if err == sql.ErrNoRows {
		log.Warn("email not found")
		return u, false, err
	}

	if err != nil {
		log.Warn("Query error")
		return u, false, err
	}

	match, err := utils.CheckPasswordHash(password, pwd)
	if !match {
		log.Warn("Hash and password doesnt match")
		return u, false, err
	}

	log.Info()
	return u, true, nil
}

func (ur *UserRepo) Register(newUser entity.User) (entity.User, error) {
	var u entity.User
	query := "INSERT INTO users(name, email, phone, password) VALUES(?,?,?,?)"

	rows, err := ur.Db.Query(query, newUser.Name, newUser.Email, newUser.Phone, newUser.Password)
	if err != nil {
		log.Warn("Query error")
		return u, err
	}
	defer rows.Close()

	return u, nil
}
