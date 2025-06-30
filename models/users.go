package models

import (
	"be-weeklytask/dto"
	"be-weeklytask/utils"
	"context"

	"github.com/jackc/pgx/v5"
)

type User struct {
	UserId   int    `db:"id" json:"userId"`
	Fullname string `db:"fullname" json:"fullname"`
	Email    string `db:"email" json:"email" binding:"required,email"`
	Phone    string `db:"phone" json:"phone"`
	Password string `db:"password" json:"password" binding:"required"`
	Pin      string `db:"pin" json:"pin" binding:"required"`
}

func FindUserById(userId int) (User, error) {
	conn, err := utils.DBConnect()

	if err != nil {
		return User{}, err
	}

	defer conn.Close()

	var user User

	err = conn.QueryRow(
		context.Background(),
		`SELECT id, fullname, email, phone, password, pin FROM users WHERE id
		= $1`,
		userId,
	).Scan(&user.UserId, &user.Fullname, &user.Email, &user.Phone, &user.Password, &user.Pin)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func FindUserByNameOrPhone(search string) ([]dto.UserRetrieved, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var rows pgx.Rows
	if search != "" {
		rows, err = conn.Query(
			context.Background(),
			`SELECT id, fullname, phone FROM users 
       WHERE fullname ILIKE '%' || $1 || '%' 
       OR phone ILIKE '%' || $1 || '%'`,
			search,
		)
	} else {
		rows, err = conn.Query(
			context.Background(),
			`SELECT id, fullname, phone FROM users`,
		)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.UserRetrieved])
	if err != nil {
		return nil, err
	}
	return users, nil
}
