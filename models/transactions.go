package models

import (
	"be-weeklytask/utils"
	"context"
	"errors"
)

type TopupRequest struct {
	Amount float64 `json:"amount" binding:"required"`
	Method string  `json:"method" binding:"required"`
}

type TransferRequest struct {
	ReceiverUserID int     `json:"receiver_user_id" binding:"required"`
	TransferAmount float64 `json:"transfer_amount" binding:"required"`
	Notes          string  `json:"notes"`
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

func HandleTransfer(userId int, req TransferRequest, senderBalance float64) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	if senderBalance < req.TransferAmount {
		return errors.New("saldo tidak cukup")
	}

	_, err = conn.Exec(
		context.Background(),
		`UPDATE wallets SET balance = balance - $1, updated_at = CURRENT_TIMESTAMP WHERE user_id = $2`,
		req.TransferAmount, userId,
	)
	if err != nil {
		return err
	}

	_, err = conn.Exec(
		context.Background(),
		`UPDATE wallets SET balance = balance + $1, updated_at = CURRENT_TIMESTAMP WHERE user_id = $2`,
		req.TransferAmount, req.ReceiverUserID,
	)
	if err != nil {
		return err
	}

	_, err = conn.Exec(
		context.Background(),
		`
		INSERT INTO transfers (sender_user_id, receiver_user_id, transfer_amount, notes, success)
		VALUES ($1, $2, $3, $4, TRUE)
		`,
		userId, req.ReceiverUserID, req.TransferAmount, req.Notes,
	)

	if err != nil {
		return err
	}

	return nil
}
