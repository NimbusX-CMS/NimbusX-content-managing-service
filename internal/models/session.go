package models

type Session struct {
	ID          int    `json:"id" gorm:"primary_key;auto_increment"`
	UserID      int    `json:"user_id"`
	User        User   `json:"user" gorm:"foreignkey:UserID"`
	CookieValue string `json:"value"`
	ValidUntil  int64  `json:"valid_until"`
}

type CookieValueOnly struct {
	CookieValue string `json:"session"`
}
