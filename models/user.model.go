package models

type User struct {
	// ID     int            `json:"id" form:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username" gorm:"primaryKey"`
	Password string `json:"-" form:"password" gorm:"not null"`
	Email    string `json:"email" form:"email" gorm:"not null"`
	FullName string `json:"full_name" form:"full_name" gorm:"not null"`
	// Locker LockerResponse `json:"locker"`
	// Posts  []PostResponse `json:"posts"`
}
type UserResponse struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	FullName string `json:"full_name" form:"full_name"`
}

func (UserResponse) TableName() string {
	return "users"
}
