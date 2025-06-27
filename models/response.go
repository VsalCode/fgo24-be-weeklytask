package models;

type Response struct {
	Success  bool	 `json:"success"`
	Message string `json:"message"`
	Error  string `json:"error,omitempty"`
	Result  any    `json:"data,omitempty"`
}
