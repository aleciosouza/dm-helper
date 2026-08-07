package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Name         string

	Sheets []Sheet `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{ID: u.ID, Email: u.Email, Name: u.Name}
}
