package models

import (
	"be-weeklytask/utils"
	"context"
)

func GetBalance(userId int) (float64, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var balance float64
	err = conn.QueryRow(
		context.Background(),
		`SELECT amount FROM balance WHERE user_id = $1`,
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
