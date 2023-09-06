package user

type UserResult struct {
	User
}

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
