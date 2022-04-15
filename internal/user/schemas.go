package user

type UpdateFirstNameSchema struct {
	NewFirstName string `json:"new_first_name" binding:"required" form:"NewFirstName"`
}
type UpdateLastNameSchema struct {
	NewLastName string `json:"new_last_name" binding:"required" form:"NewLastName"`
}
type UpdateUsernameSchema struct {
	NewUsername string `json:"new_username" binding:"required" form:"NewUsername"`
}
type UpdatePasswordSchema struct {
	NewPassword string `json:"new_password" binding:"required" form:"newPassword"`
}
type GetInfoAboutUserSchema struct {
	UserId int `json:"user_id" form:"userId"`
}
