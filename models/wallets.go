package models

import (
	"be-weeklytask/utils"
	"context"
)

type FinanceRecords struct {
	Balance float64 `json:"balance"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
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

func GetFinanceRecords(userId int) (FinanceRecords, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return FinanceRecords{}, err
	}
	defer conn.Close()

	var records FinanceRecords

	err = conn.QueryRow(
		context.Background(),
		`SELECT balance FROM wallets WHERE user_id = $1`,
		userId,
	).Scan(&records.Balance)
	if err != nil && err.Error() != "no rows in result set" {
		return FinanceRecords{}, err
	}

	err = conn.QueryRow(
		context.Background(),
		`SELECT SUM(transfer_amount) FROM transfers WHERE receiver_user_id = $1`,
		userId,
	).Scan(&records.Income)
	if err != nil {
		return FinanceRecords{}, err
	}

	err = conn.QueryRow(
		context.Background(),
		`SELECT SUM(transfer_amount) FROM transfers WHERE sender_user_id = $1`,
		userId,
	).Scan(&records.Expense)
	if err != nil {
		return FinanceRecords{}, err
	}

	return records, nil
}
