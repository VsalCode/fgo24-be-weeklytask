package models

import (
	"be-weeklytask/utils"
	"context"
)

type TopupRequest struct {
	Amount float64 `json:"amount" binding:"required"`
	Method string  `json:"method" binding:"required"`
}

func GetBalance(userId int) (float64, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var balance float64
	err = conn.QueryRow(
		context.Background(),
		`SELECT balance FROM wallets WHERE user_id = $1`,
		userId,
	).Scan(&balance)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0, nil
		}
		return 0, err
	}

	return balance, nil
}

func GetMethodIDByName(methodName string) (int, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var methodID int
	err = conn.QueryRow(
		context.Background(),
		`SELECT id FROM payment_method WHERE method_name = $1`,
		methodName,
	).Scan(&methodID)
	if err != nil {
		return 0, err
	}
	return methodID, nil
}

func HandleTopup(userId int, amount float64, methodID int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO topup (user_id, topup_amount, method_id, success) VALUES ($1, $2, $3, TRUE)`,
		userId, amount, methodID,
	)
	if err != nil {
		return err
	}

	_, err = conn.Exec(
		context.Background(),
		`UPDATE wallets SET balance = balance + $1, updated_at = CURRENT_TIMESTAMP WHERE user_id = $2`,
		amount, userId,
	)
	if err != nil {
		return err
	}

	return nil
}
