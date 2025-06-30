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
type HistoryTransaction struct {
	ID             int     `json:"id"`
	Fullname       string  `json:"fullname"`
	Phone          string  `json:"phone"`
	TransferAmount float64 `json:"transfer_amount"`
	Status         string  `json:"status"`
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

func GetTransactionHistory(userId int) ([]HistoryTransaction, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(
		context.Background(),
		`
    SELECT 
    	t.id,
      CASE 
				WHEN t.sender_user_id = $1 THEN u_receiver.fullname
      	ELSE u_sender.fullname
      	END AS fullname,
      CASE 
				WHEN t.sender_user_id = $1 THEN u_receiver.phone
      	ELSE u_sender.phone
      	END AS phone,
      t.transfer_amount,
      CASE 
        WHEN t.sender_user_id = $1 THEN 'Send'
        ELSE 'Receive'
        END AS status
    FROM transfers t
    JOIN users u_sender ON u_sender.id = t.sender_user_id
    JOIN users u_receiver ON u_receiver.id = t.receiver_user_id
    WHERE t.sender_user_id = $1 OR t.receiver_user_id = $1
    ORDER BY t.transfer_date DESC
    `,
		userId,
	)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []HistoryTransaction
	for rows.Next() {
		var transaction HistoryTransaction
		err = rows.Scan(&transaction.ID, &transaction.Fullname, &transaction.Phone, &transaction.TransferAmount, &transaction.Status); 
		
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

