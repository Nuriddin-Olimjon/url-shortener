package entity

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"-"`
}

type CreateUserParams struct {
	Username string `json:"username" binding:"required,min=5,alphanum"`
	FullName string `json:"full_name" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
}
