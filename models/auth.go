package models

import (
	"be-weeklytask/dto"
	"be-weeklytask/utils"
	"context"
	"strconv"
)

func HandleRegister(user dto.RegisterRequest) (int, error) {
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

func FindUserByEmail(email string) (dto.UserLoginData, error) {
	var user dto.UserLoginData

	conn, err := utils.DBConnect()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	row := conn.QueryRow(
		context.Background(),
		`SELECT id, email, password, pin FROM users WHERE email = $1`,
		email,
	)

	err = row.Scan(&user.UserId, &user.Email, &user.Password, &user.Pin)
	if err != nil {
		return user, err
	}

	return user, nil
}
