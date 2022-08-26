package user

type RequestUser struct {
	Account  string `json:"account" form:"account" query:"account"`
	Password string `json:"password" form:"password" query:"password"`
}
