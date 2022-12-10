package models

type RegisterDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthenticateDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ForgetPasswordDTO struct {
	Name string `json:"name"`
	New  string `json:"new"`
}
