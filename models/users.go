package models

import (
	"be-weeklytask/utils"
	"context"
	"strconv"
)

type User struct {
	// ID        int       `db:"id" json:"id"`
	Fullname  string    `db:"fullname" json:"fullname"`
	Email     string    `db:"email" json:"email" binding:"required,email"`
	Phone     string    `db:"phone" json:"phone"` 
	Password  string    `db:"password" json:"password" binding:"required"`
	Pin       string    `db:"pin" json:"pin" binding:"required"`
	// CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func HandleRegister(user User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	pin, _ := strconv.Atoi(user.Pin)

	_, err = conn.Exec(
		context.Background(),
		`
		INSERT INTO users (fullname, email, phone, password, pin)
		VALUES ($1, $2, $3, $4, $5)
		`,
		user.Fullname, user.Email, user.Phone, user.Password, pin,
	)

	if err != nil {
		return err
	}

	return nil
}
