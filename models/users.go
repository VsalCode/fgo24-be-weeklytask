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

func GetUpdateUser(userId int, user dto.UpdatedUser) (dto.UpdatedUser, error) {
    conn, err := utils.DBConnect()
    if err != nil {
        return dto.UpdatedUser{}, err
    }
    defer conn.Close()

    var oldUser dto.UpdatedUser
    err = conn.QueryRow(
        context.Background(),
        `SELECT id, fullname, email, phone, password, pin FROM users WHERE id = $1`,
        userId,
    ).Scan(&oldUser.Id, &oldUser.Fullname, &oldUser.Email, &oldUser.Phone, &oldUser.Password, &oldUser.Pin)
    if err != nil {
        return dto.UpdatedUser{}, err
    }

    if user.Fullname == nil  {
        user.Fullname = oldUser.Fullname
    }
    if user.Email == nil  {
        user.Email = oldUser.Email
    }
    if user.Phone == nil  {
        user.Phone = oldUser.Phone
    }
    if user.Password == nil  {
        user.Password = oldUser.Password
    }
    if user.Pin == nil  {
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
        return dto.UpdatedUser{}, err
    }

    user.Id = userId
    return user, nil
}

func FindUserByName(search string) ([]dto.UserRetrieved, error) {
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

	var users []dto.UserRetrieved
	for rows.Next() {
		var user dto.UserRetrieved
		if err := rows.Scan(&user.Id, &user.Fullname, &user.Phone); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

