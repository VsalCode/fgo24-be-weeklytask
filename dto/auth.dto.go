package dto;

type RegisterRequest struct {
    Fullname string `json:"fullname" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Phone    string `json:"phone" binding:"required"`
    Password string `json:"password" binding:"required"`
    Pin      string `json:"pin" binding:"required"`
}

type LoginRequest struct {
	Email    string `db:"email" json:"email" binding:"required,email"`
	Password string `db:"password" json:"password" binding:"required"`
	Pin      string `db:"pin" json:"pin" binding:"required"`
}

type UserLoginData struct {
    UserId   int    `db:"id" json:"userId"`
    Email    string `db:"email" json:"email"`
    Password string `db:"password" json:"password"`
    Pin      string `db:"pin" json:"pin"`
}