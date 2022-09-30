package req

type CreateUserReq struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	NikeName    string `json:"nikeName"`
	PhotoURL    string `json:"photoURL"`
}
