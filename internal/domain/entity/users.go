package entity

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"-"`
}

type CreateUserParams struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}
