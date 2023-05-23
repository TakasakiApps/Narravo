package enitity

type VerifyData struct {
	Data  string `json:"data" binding:"required"`
	Token string `json:"token" binding:"required"`
}
