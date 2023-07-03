package entity

type VerifyData struct {
	Data         string `json:"data" binding:"required"`
	IntegrityKey string `json:"integrityKey" binding:"required"`
	Timestamp    string `json:"timestamp" binding:"required"`
}
