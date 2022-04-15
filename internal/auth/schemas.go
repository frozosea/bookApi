package auth

type UserAuth struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type UserRegister struct {
	Username  string `json:"username" form:"username"`
	FirstName string `json:"firstName" form:"firstName"`
	Lastname  string `json:"lastname" form:"lastName"`
	Password  string `json:"password" form:"password"`
}
