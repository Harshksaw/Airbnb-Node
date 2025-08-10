package db

import (
	"database/sql"
	"fmt"
	models "AuthInGo/models"
)

type UserRepository interface {
	// Define methods for user repository
	CreateUser() error

}

type UserRepositoryImpl struct {
	db *sql.DB

}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) Create() error {
	// Implementation for creating a user in the database
	fmt.Println("fetching  user in UserRepository")
	return nil
}
func (u *UserRepositoryImpl) GetById(id int) ( error) {
	// Implementation for getting a user by ID from the database
	fmt.Println("Fetching user by ID in UserRepository")


	//Step1 :Prepare the SQL query

	query := "SELECT id, username, email , password , created_at , updated_at FROM users WHERE id = ?"

	//Step2 : Execute the query
	row  := u.db.QueryRow(query, 1)

	user := &models.User{}

	err  := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil
		} else{
			fmt.Println("Error fetching user by ID:", err)
			return err
		}
	}
	fmt.Println(user, "User fetched successfully")


	// Example: return a dummy user for demonstration
	return  nil
}