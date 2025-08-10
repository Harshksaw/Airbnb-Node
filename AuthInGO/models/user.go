package models


type User struct {	
	Id int64
	Username string
	Email string
	Password string
	CreatedAt string
	UpdatedAt string
}

// UserRepository interface defines methods for user repository operations