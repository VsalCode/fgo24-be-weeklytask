package models

import (
	"be-weeklytask/utils"
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRetrieved struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
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

func GetUpdateUser(userId int, user User) (User, error) {
	conn, err := utils.DBConnect()

	if err != nil {
		return User{}, err
	}

	defer conn.Close()

	oldUser, err := FindUserById(userId)

	if user.Fullname == "" {
		user.Fullname = oldUser.Fullname
	}
	if user.Email == "" {
		user.Email = oldUser.Email
	}
	if user.Phone == "" {
		user.Phone = oldUser.Phone
	}
	if user.Password == "" {
		user.Password = oldUser.Password
	}
	if user.Pin == "" {
		user.Pin = oldUser.Pin
	}

	_, err = conn.Exec(
		context.Background(),
		`
        UPDATE users
        SET fullname = $1, email = $2, phone = $3, password = $4, pin = $5
        WHERE id = $6
        `,
		user.Fullname, user.Email, user.Phone, user.Password, user.Pin, userId,
	)

	if err != nil {
		return User{}, err
	}

	user.UserId = userId

	return user, nil
}

func FindUserByName(search string) ([]UserRetrieved, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var rows pgx.Rows
	if search != "" {
		rows, err = conn.Query(
			context.Background(),
			`SELECT id, fullname, phone FROM users WHERE fullname ILIKE '%' || $1 || '%'`,
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

	var users []UserRetrieved
	for rows.Next() {
		var user UserRetrieved
		if err := rows.Scan(&user.Id, &user.Fullname, &user.Phone); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

