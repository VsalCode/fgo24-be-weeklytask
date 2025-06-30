package dto;

type UserRetrieved struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
}

type UpdatedUser struct {
	Id   int    `json:"Id,omitempty"`
	Fullname *string `json:"fullname,omitempty"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Password *string `json:"password,omitempty"`
	Pin      *int `json:"pin,omitempty"`
}