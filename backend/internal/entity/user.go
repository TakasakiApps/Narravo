package entity

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement" json:"uid"`
	Name     string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

type UserToken struct {
	Token string `json:"token"`
}
