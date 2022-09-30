package models

type LoginUser struct {
	Username string `form:"username"`
	Pwd      string `form:"password"`
}
