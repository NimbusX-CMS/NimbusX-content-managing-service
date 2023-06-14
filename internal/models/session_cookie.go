package models

type SessionCookie struct {
	ID         int    `json:"id" gorm:"primary_key;auto_increment"`
	UserID     int    `json:"user_id"`
	User       User   `json:"user" gorm:"foreignkey:UserID"`
	Value      string `json:"value"`
	ValidUntil int64  `json:"valid_until"`
}
