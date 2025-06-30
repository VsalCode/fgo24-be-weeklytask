package models

import (
	"be-weeklytask/dto"
	"be-weeklytask/utils"
	"context"
)

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
