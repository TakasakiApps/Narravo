package entity

type Content struct {
	Text string `json:"text" binding:"required"`
}
