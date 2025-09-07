package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {

	GetById(id string) (*models.User, error)
	Create(username string, email string, password string) ( error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteByID(id int64 )error

}

type UserRepositoryImpl struct {
	db *sql.DB

}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}


func (u *UserRepositoryImpl) Create(username string, email string, password string) ( error) {
	// Implementation for creating a user in the database
	fmt.Println("Creating user in UserRepository")
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, username, email, password)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return err
	}
	rowsAffected , rowsErr := result.RowsAffected()

	if rowsErr != nil {
		fmt.Println("Error fetching rows affected:", rowsErr)
		return rowsErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were inserted")
		return nil
	}
	if rowsAffected > 0 {
		fmt.Println("User created successfully")
	}

	return nil
}
func (u *UserRepositoryImpl) GetById(id string) (*models.User, error) {
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
			return nil,err
		} else{
			fmt.Println("Error fetching user by ID:", err)
			return nil, err
		}
	}
	fmt.Println(user, "User fetched successfully")


	// Example: return a dummy user for demonstration
	return  user, nil
}
func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = ?"

	row := u.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Email, &user.Password) // hashed password

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}

	return user, nil
}


func (u *UserRepositoryImpl) LoginUser(email string , password  string) (error){
	//fetch user by email and password from database

	fmt.Println("Logging in user in UserRepository")
	query := "SELECT id, username, email , password , created_at , updated_at FROM users WHERE email = ? AND password = ?"

	row := u.db.QueryRow(query, email, password)

	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email and password")
			return err
		}
		fmt.Println("Error logging in user:", err)
		return err
	}
	
	fmt.Println("User logged in successfully")
	return nil
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	fmt.Println("Fetching all users in UserRepository")
	
	query := "SELECT id, username, email, password, created_at, updated_at FROM users"
	
	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching all users:", err)
		return nil, err
	}
	defer rows.Close()
	
	var users []*models.User
	
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}
	
	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	
	fmt.Println("All users fetched successfully")
	return users, nil
}

func (u *UserRepositoryImpl) DeleteByID(id int64) error {
	fmt.Println("Deleting user by ID in UserRepository")
	
	query := "DELETE FROM users WHERE id = ?"
	
	result, err := u.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error fetching rows affected:", err)
		return err
	}
	
	if rowsAffected == 0 {
		fmt.Println("No user found with the given ID")
		return fmt.Errorf("no user found with ID %d", id)
	}
	
	fmt.Println("User deleted successfully")
	return nil
}