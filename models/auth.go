package models

import (
	"be-weeklytask/utils"
	"context"
	"strconv"
)

type LoginRequest struct {
	Email    string `db:"email" json:"email" binding:"required,email"`
	Password string `db:"password" json:"password" binding:"required"`
	Pin      string `db:"pin" json:"pin" binding:"required"`
}

func HandleRegister(user User) (int, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	pin, _ := strconv.Atoi(user.Pin)

	var userId int
	err = conn.QueryRow(
		context.Background(),
		`
		INSERT INTO users (fullname, email, phone, password, pin)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`,
		user.Fullname, user.Email, user.Phone, user.Password, pin,
	).Scan(&userId)

	if err != nil {
		return 0, err
	}
	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO wallets (user_id, balance) VALUES ($1, 0)`,
		userId,
	)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func FindUserByEmail(email string) (User, error) {
	var user User

	conn, err := utils.DBConnect()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	row := conn.QueryRow(
		context.Background(),
		`
		SELECT id, fullname, email, phone, password, pin FROM users WHERE email = $1
		`,
		email,
	)

	row.Scan(&user.UserId, &user.Fullname, &user.Email, &user.Phone, &user.Password, &user.Pin)

	return user, nil
}
